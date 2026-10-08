package api

import (
	"context"
	"flag"
	"fmt"
)

// StartServerCommand is struct for info required to start an http server
type StartServerCommand struct {
	port    string
	address string
	tls     bool
}

// Flags returns the flag sets
func (sc *StartServerCommand) Flags() *flag.FlagSet {
	f := &flag.FlagSet{}

	f.StringVar(&sc.address, "address", "localhost", "-address=localhost")
	f.StringVar(&sc.port, "port", "8080", "-port=8080")
	f.BoolVar(&sc.tls, "tls", false, "-tls=true")

	return f
}

// Port gets the http server port
func (sc *StartServerCommand) Port() string {
	return sc.port
}

// Address gets the server address
func (sc *StartServerCommand) Address() string {
	return sc.address
}

// TLS gets the flag whether the server should run on TLS
func (sc *StartServerCommand) TLS() bool {
	return sc.tls
}

// Name gets the name of the command used in yacli package
func (sc *StartServerCommand) Name() string {
	return "start-server"
}

// HelpString gets the string shown as usage in cli
func (sc *StartServerCommand) HelpString() string {
	return "Start the server using provided address and port"
}

// Run the start server command
func (sc *StartServerCommand) Run(ctx context.Context, args []string) error {
	config, err := LoadConfig()
	if err != nil {
		return fmt.Errorf("load server configuration: %w", err)
	}
	if sc.Address() == "localhost" {
		sc.address = ""
	}

	return StartServer(sc.Address(), sc.Port(), sc.TLS(), config)
}
