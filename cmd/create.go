/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

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
		// var todoDescription = args[0]

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

// func addTask(file *os.File, task_description string) {

// 	w := csv.NewWriter(file)
// 	// Write any buffered data to underlying writer
// 	defer w.Flush()

// 	// Add header row
// 	if err := w.Write(task_description); err != nil {
// 		log.Fatalln("error writing headers to csv:", err)
// 	}
// }
