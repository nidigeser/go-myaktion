package main

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/nidigeser/go-myaktion/src/banktransfer/grpc/banktransfer"
	"github.com/nidigeser/go-myaktion/src/myaktion/client"
)

func monitortransactions() {
	for {
		connectandmonitor()
		time.Sleep(time.Second)
	}
}

func connectandmonitor() {
	conn, err := client.GetBankTransferConnection()
	if err != nil {
		log.WithError(err).Fatal("error connecting to the banktransfer service")
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	banktransferClient := banktransfer.NewBankTransferClient(conn)
	watcher, err := banktransferClient.ProcessTransactions(ctx)
	if err != nil {
		log.WithError(err).Fatal("error watching transactions")
	}
	log.Info("Successfulyy connected to banktransfer service for processing transactions")
	for {
		transaction, err := watcher.Recv()
		if err != nil {
			if _, deadline := ctx.Deadline(); deadline {
				log.Info("deadline reached- reconnect client")
				break
			}
			log.WithError(err).Error("error receiving transaction")
			continue
		}
		entry := log.WithField("transaction", transaction)
		entry.Info("Received transaction. Sending processing response")
		err = watcher.Send(&banktransfer.ProcessingResponse{Id: transaction.Id})
		if err != nil {
			entry.WithError(err).Error("error sending processing response")
			continue
		}
		entry.Info("Processing response sent")
	}
}
