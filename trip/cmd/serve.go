package cmd

import (
	"github.com/spf13/cobra"
	"trip/pkg/bootstrap"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve app on dev server",
	Long:  "Serve app on dev server",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	bootstrap.Run()
}
