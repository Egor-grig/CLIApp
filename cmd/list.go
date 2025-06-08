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

		tdIsSet := cmd.Flags().Lookup("todos").Changed
		pIsSet := cmd.Flags().Lookup("progress").Changed
		dIsSet := cmd.Flags().Lookup("done").Changed

		for i := 0; i < len(tasks.Tasks); i++ {

			if tdIsSet && tasks.Tasks[i].State == "To Do" {
				fmt.Println(tasks.Tasks[i].Text, " ", tasks.Tasks[i].State)
			}

			if pIsSet && tasks.Tasks[i].State == "In Progress" {
				fmt.Println(tasks.Tasks[i].Text, " ", tasks.Tasks[i].State)
			}

			if dIsSet && tasks.Tasks[i].State == "Done" {
				fmt.Println(tasks.Tasks[i].Text, " ", tasks.Tasks[i].State)
			}

			if !tdIsSet && !pIsSet && !dIsSet {
				fmt.Println(tasks.Tasks[i].Text, " ", tasks.Tasks[i].State)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("todos", "t", true, "List to do tasks")
	listCmd.Flags().BoolP("progress", "p", true, "List in progress tasks")
	listCmd.Flags().BoolP("done", "d", true, "List done tasks")

}
