/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/LintaoAmons/kubectl-spec-template/content/namespaces"
	"github.com/spf13/cobra"
)

// namespacesCmd represents the namespaces command
var namespacesCmd = &cobra.Command{
	Use:   "namespaces",
	Aliases: []string{"ns"},
	Run: func(cmd *cobra.Command, args []string) {
		Output(namespaces.Simple)
	},
}

func init() {
	rootCmd.AddCommand(namespacesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// namespacesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// namespacesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
