package main

type config struct {
	address string
	port    int
	command string
}

func newConfig(address string, port int, command string) config {
	return config{
		address: address,
		port:    port,
		command: command,
	}
}
