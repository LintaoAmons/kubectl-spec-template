/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/LintaoAmons/kubectl-spec-template/content/networkpolicies"
	"github.com/spf13/cobra"
)

var networkpoliciesCmd = &cobra.Command{
	Use:     "networkpolicies",
	Aliases: []string{"netpol"},
	Run: func(cmd *cobra.Command, args []string) {
		Output(networkpolicies.Simple)
	},
}

func init() {
	rootCmd.AddCommand(networkpoliciesCmd)
}
