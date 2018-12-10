package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ivch/go15"
	"github.com/nsf/termbox-go"
)

var size int

const defaultSize = 4

func init() {
	flag.IntVar(&size, "size", defaultSize, "integer value of grid size")
	flag.Parse()
}

func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	game := go15.Init(byte(size))

	ch := make(chan os.Signal, 1)
	defer close(ch)

	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	go stopper(ch, game)

	game.Run()
}

func stopper(ch <-chan os.Signal, g go15.Game) {
	<-ch
	g.Stop()
}
