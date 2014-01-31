package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jbcrail/vindinium-starter-go/vindinium"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s <key> [arguments]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	hostname := flag.String("server", "http://vindinium.org", "Hostname of game server")
	mode := flag.String("mode", "training", "Game mode")
	turns := flag.Int("turns", 300, "Turns per game")
	bot := flag.String("bot", "none", "Bot")

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}

	c := vindinium.NewClient(*hostname, args[0], *mode, *bot, *turns)
	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}
}
