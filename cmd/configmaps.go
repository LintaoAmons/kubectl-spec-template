/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/LintaoAmons/kubectl-spec-template/content/configmaps"
	"github.com/spf13/cobra"
)

var configmapsCmd = &cobra.Command{
	Use:     "configmaps",
	Aliases: []string{"cm"},
	Run: func(cmd *cobra.Command, args []string) {
		Output(configmaps.Simple)
	},
}

func init() {
	rootCmd.AddCommand(configmapsCmd)
}
