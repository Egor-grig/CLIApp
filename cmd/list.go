/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show tasks",
	Long:  `Show tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		openJSON("cmd/data.json", &tasks)

		for i := 0; i < len(tasks.Tasks); i++ {
			fmt.Println(tasks.Tasks[i].Text, " ", tasks.Tasks[i].State)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
