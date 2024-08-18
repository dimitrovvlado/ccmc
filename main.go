package main

import (
	"fmt"
	"os"

	"github.com/dimitrovvlado/ccmc/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
