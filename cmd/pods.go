/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/LintaoAmons/kubectl-spec-template/content/pods"
	"github.com/spf13/cobra"
)

// podsCmd represents the pods command
var podsCmd = &cobra.Command{
	Use:   "pods",
	Aliases: []string{"po"},
	Run: func(cmd *cobra.Command, args []string) {
		Output(pods.Simple)
	},
}

func init() {
	rootCmd.AddCommand(podsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// podsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// podsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
