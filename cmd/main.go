package main

import (
	"github.com/gabrielsc1998/go-stress-test/internal/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cobra.Command{}
	rootCmd.AddCommand(cmd.StressTestCmd())
	rootCmd.Execute()
}
