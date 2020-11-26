package main

import (
	"flag"
	"os"
	"runtime"

	configuration "../../internal/configuration"
	cmd "../../pkgs/cmd"

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

func parseFlags() configuration.Config {
	address := flag.String("address", "localhost", "Server hostname")
	port := flag.Int("port", 8000, "Server port")
	command := flag.String("command", "echo apa", "command to be executed")
	debug := flag.Bool("debug", false, "Enable debug logging")

	flag.Parse()

	config := configuration.NewConfig(*address, *port, *command, *debug)
	return config
}

func getShell() string {
	os := runtime.GOOS
	switch os {
	case "windows":
		return "powershell"
	case "linux":
		return "sh"
	default:
		panic("Unsupported os")
	}
}

func main() {
	config := parseFlags()

	if config.Debug {
		log.SetLevel(log.DebugLevel)
	}

	log.Infof("Config: { address: %s, port: %d, command: %v }", config.Address, config.Port, config.Command)

	shell := getShell()
	c := cmd.NewCmd(shell, []string{config.Command})
	result, err := c.Execute()
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Result: %s, %d", result.Data, result.ExitCode)

}
