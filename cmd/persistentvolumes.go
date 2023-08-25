/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/LintaoAmons/kubectl-spec-template/content/persistentvolumes"
	"github.com/spf13/cobra"
)

// persistentvolumesCmd represents the persistentvolumes command
var persistentvolumesCmd = &cobra.Command{
	Aliases: []string{"pv"},
	Use:     "persistentvolumes",
	Run: func(cmd *cobra.Command, args []string) {
		Output(persistentvolumes.Simple)
	},
}

func init() {
	rootCmd.AddCommand(persistentvolumesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// persistentvolumesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// persistentvolumesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
