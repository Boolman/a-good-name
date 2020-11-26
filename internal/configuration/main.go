package configuration

// Config struct
type Config struct {
	Address string
	Port    int
	Command string
	Debug   bool
}

// Generate en return Config struct
func NewConfig(address string, port int, command string, debug bool) Config {
	return Config{
		Address: address,
		Port:    port,
		Command: command,
		Debug:   debug,
	}
}
