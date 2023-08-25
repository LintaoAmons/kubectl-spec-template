/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/LintaoAmons/kubectl-spec-template/content/ingresses"
	"github.com/spf13/cobra"
)

// ingressesCmd represents the ingresses command
var ingressesCmd = &cobra.Command{
	Use:     "ingresses",
	Aliases: []string{"ing"},
	Run: func(cmd *cobra.Command, args []string) {
		Output(ingresses.Simple)
	},
}

func init() {
	rootCmd.AddCommand(ingressesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ingressesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ingressesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
