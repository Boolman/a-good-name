package main

type config struct {
	address string
	port    int
	command string
	debug   bool
}

func newConfig(address string, port int, command string, debug bool) config {
	return config{
		address: address,
		port:    port,
		command: command,
		debug:   debug,
	}
}
