/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Text  string `json:"text"`
	State string `json:"state"`
}

var jsonName string = "cmd/data.json"

var tasks Tasks

var taskText string
var taskState string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  "Add a new task with name and state",
	RunE: func(cmd *cobra.Command, args []string) error {

		openJSON(jsonName, &tasks)

		if taskState != "To Do" && taskState != "In Progress" && taskState != "Done" {
			return errors.New("must be 'To Do', 'In Progress' or 'Done'")
		}

		tasks.Tasks = append(tasks.Tasks, Task{taskText, taskState})

		writeJSON(jsonName, tasks)

		fmt.Println("Task add")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&taskText, "text", "t", "", "Save task text")
	addCmd.Flags().StringVarP(&taskState, "state", "s", "", "Save task state")

}

// Open json file and return decoding byte array
func openJSON(fileName string, tasks *Tasks) {
	file, err := os.OpenFile(fileName, os.O_RDWR, os.ModeAppend)
	if err != nil {
		fmt.Println("Error in open file")
		fmt.Println(err)
	}
	defer file.Close()

	jsonData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(jsonData, tasks)

}

func writeJSON(fileName string, tasks Tasks) {
	jsonInfo, err := json.Marshal(tasks)

	if err != nil {
		fmt.Println("Error in marshal tasks")
	}

	err = os.WriteFile(fileName, jsonInfo, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
