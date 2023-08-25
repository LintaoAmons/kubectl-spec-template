/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/LintaoAmons/kubectl-spec-template/content/cronjobs"
	"github.com/spf13/cobra"
)

// cronjobsCmd represents the cronjobs command
var cronjobsCmd = &cobra.Command{
	Use:     "cronjobs",
	Aliases: []string{"cj"},
	Run: func(cmd *cobra.Command, args []string) {
		Output(cronjobs.Simple)
	},
}

func init() {
	rootCmd.AddCommand(cronjobsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cronjobsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cronjobsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
