package main

import (
	"os"

	"github.com/DmitryZ-outten/trivybeat/cmd"

	_ "github.com/DmitryZ-outten/trivybeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
