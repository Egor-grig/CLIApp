/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete tasks",
	Long:  `Delete tasks from list`,
	Args:  checkArgsInt,
	Run: func(cmd *cobra.Command, args []string) {
		tasks.Tasks = append(tasks.Tasks[:elem], tasks.Tasks[elem+1:]...)
		writeJSON(jsonName, tasks)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
