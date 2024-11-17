package main

import (
	"fmt"

	"github.com/alpagit/godivert/windivert"
)

func main() {
	winDivert, err := windivert.NewWinDivertHandle("true")
	if err != nil {
		panic(err)
	}
	defer winDivert.Close()

	packet, err := winDivert.Recv()
	if err != nil {
		panic(err)
	}

	fmt.Println(packet)

}
