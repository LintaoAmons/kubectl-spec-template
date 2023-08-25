/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/LintaoAmons/kubectl-spec-template/content/deployments"
	"github.com/spf13/cobra"
)

// deploymentsCmd represents the deployments command
var deploymentsCmd = &cobra.Command{
	Use:     "deployments",
	Aliases: []string{"deploy"},
	Run: func(cmd *cobra.Command, args []string) {
		Output(deployments.Simple)
	},
}

func init() {
	rootCmd.AddCommand(deploymentsCmd)
}
