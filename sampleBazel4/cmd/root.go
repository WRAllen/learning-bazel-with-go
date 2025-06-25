package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "A Bazel + Cobra web service",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
