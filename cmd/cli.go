package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "RabbitMQ test",
	Short: "RabbitMQ test",
	Long:  "RabbitMQ test",
	Args:  cobra.MinimumNArgs(1),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
