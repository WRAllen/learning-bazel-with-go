package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wrallen/sampleBazel7/config"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "A Bazel + Cobra web service",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if _, err := config.Load(); err != nil {
			panic(fmt.Sprintf("加载配置失败: %v", err))
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
