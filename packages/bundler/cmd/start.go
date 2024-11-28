package cmd

import (
	"fmt"
	"os"

	"github.com/seokheejang/minimal-aa/packages/bundler/internal/start"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "minimal-aa",
	Short: "ERC-4337 Bundler",
	Long:  "A modular Go implementation of an ERC-4337 Bundler.",
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts an instance in the specified mode",
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetString("mode") == "private" {
			start.PrivateMode()
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var mode string

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&mode, "mode", "m", "", "Required. See acceptable values above.")
	if err := startCmd.MarkFlagRequired("mode"); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("mode", startCmd.Flags().Lookup("mode")); err != nil {
		panic(err)
	}
}
