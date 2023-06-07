package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/yanlinLiu0424/godivert/windivert"
)

func checkPacket(wd *windivert.WinDivertHandle, packetChan <-chan *windivert.Packet) {
	for packet := range packetChan {
		go func(wd *windivert.WinDivertHandle, packet *windivert.Packet) {
			log.Print(packet)
			packet.Send(wd)
		}(wd, packet)

	}
}

func main() {
	winDivert, err := windivert.NewWinDivertHandle("true")
	if err != nil {
		panic(err)
	}
	defer winDivert.Close()

	packetChan, err := winDivert.Packets()
	if err != nil {
		panic(err)
	}

	go checkPacket(winDivert, packetChan)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
