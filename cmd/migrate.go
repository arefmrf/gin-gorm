package cmd

import (
	"github.com/spf13/cobra"
	"web/pkg/bootstrap"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "table migration",
	Long:  "table migration",
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

func migrate() {
	bootstrap.Migrate()
}
