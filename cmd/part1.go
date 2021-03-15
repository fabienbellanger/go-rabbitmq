package cmd

import (
	"log"
	"rabbitmqtest/part1"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(part1Cmd)
	part1Cmd.AddCommand(part1SendCmd)
	part1Cmd.AddCommand(part1ReceiveCmd)
}

var part1Cmd = &cobra.Command{
	Use:   "part1",
	Short: "Tutorial Part 1",
	Long:  "Tutorial Part 1",
	Args:  cobra.MinimumNArgs(1),
}

var part1SendCmd = &cobra.Command{
	Use:   "producer",
	Short: "Tutorial Part 1 - Producer",
	Long:  "Tutorial Part 1 - Producer",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Part 1 - Producer...")

		part1.Send()
	},
}

var part1ReceiveCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Tutorial Part 1 - Consumer",
	Long:  "Tutorial Part 1 - Consumer",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Part 1 - Consumer...")

		part1.Receive()
	},
}
