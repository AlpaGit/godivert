package main

import (
	"log"
	"os"
	"os/signal"
	"testing"

	"github.com/yanlinLiu0424/godivert/windivert"
)

func checkPacketEx(wd *windivert.WinDivertHandle, packetChan <-chan *windivert.Packet) {
	for packet := range packetChan {
		go func(wd *windivert.WinDivertHandle, packet *windivert.Packet) {
			p, err := wd.HelperParsePacket(packet.Raw)
			if err == nil {
				log.Print(p)
			}
			wd.Send(packet)
		}(wd, packet)

	}
}
func TestXxx(t *testing.T) {
	winDivert, err := windivert.NewWinDivertHandle("!loopback && udp.DstPort==55954")
	if err != nil {
		t.Fatal(err)
	}

	defer winDivert.Close()
	err = winDivert.SetParam(windivert.WinDivertParamQueueSize, windivert.WinDivertParamQueueSizeMax)
	if err != nil {
		t.Fatal(err)
	}
	err = winDivert.SetParam(windivert.WinDivertParamQueueLength, windivert.WinDivertParamQueueLengthMax)
	if err != nil {
		t.Fatal(err)
	}
	err = winDivert.SetParam(windivert.WinDivertParamQueueTime, windivert.WinDivertParamQueueTimeMax)
	if err != nil {
		t.Fatal(err)
	}

	packetChan, err := winDivert.Packets()
	if err != nil {
		t.Fatal(err)
	}

	go checkPacketEx(winDivert, packetChan)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func TestGetParam(t *testing.T) {
	winDivert, err := windivert.NewWinDivertHandle("icmp")
	if err != nil {
		t.Fatal(err)
	}
	defer winDivert.Close()
	v, err := winDivert.GetParam(windivert.WinDivertParamQueueSize)
	if err != nil {
		t.Fatal(err)
	}
	log.Print(v)
}
func TestSetParam(t *testing.T) {
	winDivert, err := windivert.NewWinDivertHandle("true")
	if err != nil {
		t.Fatal(err)
	}
	defer winDivert.Close()
	err = winDivert.SetParam(windivert.WinDivertParamQueueSize, windivert.WinDivertParamQueueSizeMax)
	if err != nil {
		t.Fatal(err)
	}

}