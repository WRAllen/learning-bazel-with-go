package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wrallen/sampleBazel7/pkg"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
