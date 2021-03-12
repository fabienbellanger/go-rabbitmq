package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "RabbitMQ test",
	Short: "RabbitMQ test",
	Long:  `RabbitMQ test`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		println("In cmd...")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
