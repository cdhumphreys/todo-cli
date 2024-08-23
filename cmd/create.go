/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/spf13/cobra"
)

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

		if len(args) == 0 {
			fmt.Println("No description given")
			return
		}
		var todoDescription = args[0]

		if todoDescription == "" {
			fmt.Println("No description given")
			return
		}

		// Open the file with read
		todosFile, err := os.OpenFile(CSV_FILENAME, os.O_CREATE|os.O_RDWR, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer todosFile.Close()

		// Reset the file reader
		if _, err := todosFile.Seek(0, io.SeekStart); err != nil {
			panic(err)
		}

		newTodos := []*Todo{}

		todos, err := GetTodos(todosFile)

		if err == nil {
			newTodos = append(newTodos, todos...)
		}

		newTodos = append(newTodos, &Todo{len(newTodos), todoDescription, time.Now().Format(TimeFormat), false})

		WriteTodos(todosFile, newTodos)
	},
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
