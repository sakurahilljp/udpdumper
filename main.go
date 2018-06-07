package main

import (
	"fmt"
	"log"
	"net"
	"os"

	docopt "github.com/docopt/docopt-go"
)

// AppName --
const AppName = "udpdumper"

// Version --
var Version = "0.0.0" // placeholder
// BuildDate --
var BuildDate = "" // placeholder

var usage = `uudpdumperdp

Usage:
  udpdumper [options]
  udpdumper --help
  udpdumper --version

Options:
  --port PORT              Port to be listen  [default: 51515].
  --break                  Insert break or not.                     
  --help                   Show this screen.
  --version                Show version.
  --debug                  Enable debug
`

type options struct {
	Port  int
	Break bool

	Debug bool
}

func fatalIfError(err error, message string) {
	if err != nil {
		log.Fatal(message)
	}
}

func optparse() options {
	version := fmt.Sprintf("%v %v [%v]", AppName, Version, BuildDate)
	args, err := docopt.ParseArgs(usage, os.Args[1:], version)
	fatalIfError(err, "failed to parse arguments")

	var opts options
	if err := args.Bind(&opts); err != nil {
		fatalIfError(err, "failed to Bind")
	}

	return opts
}

func checkerror(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func main() {
	opts := optparse()

	addr := fmt.Sprintf("localhost:%v", opts.Port)
	log.Println("Listening ", addr)

	udpaddr, err := net.ResolveUDPAddr("udp", addr)
	checkerror(err)

	conn, err := net.ListenUDP("udp", udpaddr)
	checkerror(err)

	defer conn.Close()

	recv := make([]byte, 1518)

	for {
		n, err := conn.Read(recv)
		checkerror(err)
		if opts.Break {
			fmt.Println(string(recv[0:n]))
		} else {
			fmt.Print(string(recv[0:n]))
		}
	}
}
