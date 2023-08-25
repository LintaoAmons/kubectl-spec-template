/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/LintaoAmons/kubectl-spec-template/content/replicasets"
	"github.com/spf13/cobra"
)

// replicasetsCmd represents the replicasets command
var replicasetsCmd = &cobra.Command{
	Use:     "replicasets",
	Aliases: []string{"rs"},
	Run: func(cmd *cobra.Command, args []string) {
		Output(replicasets.Simple)
	},
}

func init() {
	rootCmd.AddCommand(replicasetsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// replicasetsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// replicasetsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
