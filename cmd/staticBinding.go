/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/LintaoAmons/kubectl-spec-template/content/persistentvolumes"
	"github.com/spf13/cobra"
)

// staticBindingCmd represents the staticBinding command
var staticBindingCmd = &cobra.Command{
	Use: "staticBinding",
	Run: func(cmd *cobra.Command, args []string) {
		Output(persistentvolumes.StaticBinding)
	},
}

func init() {
	persistentvolumesCmd.AddCommand(staticBindingCmd)
}
