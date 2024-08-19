/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cdh/todo-cli/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

const TimeFormat = time.RFC3339

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Add todo",
	Long: `Add a todo task, optional flags:
	--due due date
	--done set completion state (e.g. adding a task you've already complete) [bool]
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called with args", args)
		// var todo1 = Todo{1, "My first task", "2024-08-17 15:04:05", true}
		// var todo2 = Todo{2, "My second task", "2024-08-16 15:04:05", true}
		// var data = []Todo{todo1, todo2}
		file := openOrCreate()
		defer file.Close()
		currentData, err := csv.ReadCsv(file)
		if err != nil {
			log.Fatalln("Couldn't read contents")
		}
		var todoDescription = args[0]
		id := len(currentData)
		// to handle newly created / empty csv which hasn't finished writing
		if id == 0 {
			id = 1
		}
		NewTodo := csv.Todo{ID: id, Description: todoDescription, CreatedAt: time.Now().Format(TimeFormat), Completed: false}
		csv.AppendToCsv(file, NewTodo)

	},
}

func openOrCreate() *os.File {
	setHeaders := false
	if _, err := os.Stat("todos.csv"); errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does not exist
		setHeaders = true
	}
	file, err := os.OpenFile("todos.csv", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)

	if err != nil {
		log.Fatalln("Could not get or create file")
	}

	if setHeaders {
		csv.CreateCsvHeaders(file)
	}

	return file
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// func addTask(file *os.File, task_description string) {

// 	w := csv.NewWriter(file)
// 	// Write any buffered data to underlying writer
// 	defer w.Flush()

// 	// Add header row
// 	if err := w.Write(task_description); err != nil {
// 		log.Fatalln("error writing headers to csv:", err)
// 	}
// }
