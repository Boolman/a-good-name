package main

import (
	"flag"
	"os"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true)
}

func parseFlags() config {
	address := flag.String("address", "localhost", "Server hostname")
	port := flag.Int("port", 8000, "Server port")
	command := flag.String("command", "echo apa", "command to be executed")

	flag.Parse()

	config := newConfig(*address, *port, *command)
	return config
}

func getShell() string {
	os := runtime.GOOS
	switch os {
	case "windows":
		return "powershell"
	case "linux":
		return "bash"
	default:
		panic("Unsupported os")
	}
}

func main() {
	config := parseFlags()
	log.Infof("Config: { address: %s, port: %d, command: %v }", config.address, config.port, config.command)

	shell := getShell()
	c := NewCmd(shell, []string{config.command})
	result, err := c.Execute()
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Result: %s, %d", result.Data, result.ExitCode)

}
