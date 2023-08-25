/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/LintaoAmons/kubectl-spec-template/content/persistentvolumeclaims"
	"github.com/spf13/cobra"
)

// persistentvolumeclaimsCmd represents the persistentvolumeclaims command
var persistentvolumeclaimsCmd = &cobra.Command{
	Use:     "persistentvolumeclaims",
	Aliases: []string{"pvc"},
	Run: func(cmd *cobra.Command, args []string) {
		Output(persistentvolumeclaims.Simply)
	},
}

func init() {
	rootCmd.AddCommand(persistentvolumeclaimsCmd)
}
