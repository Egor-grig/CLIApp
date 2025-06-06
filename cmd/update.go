/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"strconv"

	"github.com/spf13/cobra"
)

var elem int64

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update task",
	Long:  `Update task`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requiers one argument")
		}

		num, err := strconv.ParseInt(args[0], 10, 64)
		elem = num
		if err != nil {
			return errors.New("argument must be a number")
		}

		if num < 0 {
			return errors.New("number must be greater than 0")
		}

		openJSON("cmd/data.json", &tasks)

		if len(tasks.Tasks) == 0 {
			return errors.New("there are no tasks")
		}

		return nil

	},
	Run: func(cmd *cobra.Command, args []string) {
		if taskText != "" {
			tasks.Tasks[elem].Text = taskText
		}
		if taskState != "" {
			tasks.Tasks[elem].State = taskState
		}

		writeJSON(jsonName, tasks)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&taskText, "text", "t", "", "Update task text")
	updateCmd.Flags().StringVarP(&taskState, "state", "s", "", "Update task state")
}
