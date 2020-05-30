package main

import (
	"github.com/cxnky/AnarchyGrabber-Monitor/threads"
	"github.com/cxnky/AnarchyGrabber-Monitor/util"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	util.Info("Starting AnarchyGrabber Monitor")
	threads.StartFileMonitoring()

	// prevent the application from closing, until we receive ctrl+c
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	util.Info("Closing AnarchyGrabber Monitor")

}
