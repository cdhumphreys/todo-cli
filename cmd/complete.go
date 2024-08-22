/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gocarina/gocsv"
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("complete called")
		taskId, err := strconv.Atoi(args[0])

		if err != nil {
			panic(err)
		}

		// Open the file with read
		todosFile, err := os.OpenFile(CSV_FILENAME, os.O_RDWR|os.O_TRUNC, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer todosFile.Close()

		todos, err := GetTodos(todosFile)

		if err != nil {
			panic(err)
		}

		for _,todo := range todos {
			fmt.Println(todo)
			if todo.ID == taskId {
				todo.Completed = true
			}
		}
		
		// Save the CSV back to the file
		if err := gocsv.MarshalFile(&todos, todosFile); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
