/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/LintaoAmons/kubectl-spec-template/content/serviceaccounts"
	"github.com/spf13/cobra"
)

var serviceaccountsCmd = &cobra.Command{
	Use:     "serviceaccounts",
	Aliases: []string{"sa"},
	Run: func(cmd *cobra.Command, args []string) {
		Output(serviceaccounts.Simple)
	},
}

func init() {
	rootCmd.AddCommand(serviceaccountsCmd)
}
