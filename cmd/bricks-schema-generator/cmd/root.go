package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	// rootCmd.AddCommand(
	// 	deployCmd,
	// 	importCmd,
	// 	initCmd,
	// )
}

var rootCmd = &cobra.Command{
	Use:   "bricks-generator",
	Short: "Bricks Schema Generator generates schemas for Bricks IAC tool",
	Long:  "Bricks Schema Generator generates Bricks Infrastructure as Code schemas from Terraform providers.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	rootCmd.Execute()
}
