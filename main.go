package main

import (
	"os"

	"github.com/maehue/mqttbeat/cmd"

	_ "github.com/maehue/mqttbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
