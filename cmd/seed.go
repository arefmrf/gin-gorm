package cmd

import (
	"github.com/spf13/cobra"
	"web/pkg/bootstrap"
)

func init() {
	rootCmd.AddCommand(seedCmd)
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "seed tables",
	Long:  "seed tables",
	Run: func(cmd *cobra.Command, args []string) {
		seed()
	},
}

func seed() {
	bootstrap.Seed()
}
