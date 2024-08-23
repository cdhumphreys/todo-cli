/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
		fmt.Println("complete called")
		taskId, err := strconv.Atoi(args[0])

		if err != nil {
			panic(err)
		}

		// Open the file with read
		todosFile, err := os.OpenFile(CSV_FILENAME, os.O_RDWR, os.ModePerm)
		if err != nil {
			// panic(err)
			fmt.Println("No file exists!")
			return
		}
		defer todosFile.Close()

		todos, err := GetTodos(todosFile)

		if err != nil {
			panic(err)
		}

		newTodos := []*Todo{}

		deleteTodo, deletedTaskIdx, ok := lo.FindIndexOf(todos, func(t *Todo) bool {
			return t.ID == taskId
		})
		if !ok {
			fmt.Printf("Could not find task with ID %s", strconv.Itoa(taskId))
			return
		}
		newTodos = append(newTodos, todos[:deletedTaskIdx]...)
		newTodos = append(newTodos, todos[deletedTaskIdx+1:]...)

		fmt.Printf("Removed task %s: \"%s\"", strconv.Itoa(deleteTodo.ID), deleteTodo.Description)

		WriteTodos(todosFile, newTodos)

	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
