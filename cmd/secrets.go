/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/LintaoAmons/kubectl-spec-template/content/secrets"
	"github.com/spf13/cobra"
)

var secretsCmd = &cobra.Command{
	Use:     "secrets",
	Aliases: []string{"cm"},
	Run: func(cmd *cobra.Command, args []string) {
		Output(secrets.Simple)
	},
}

func init() {
	rootCmd.AddCommand(secretsCmd)
}
