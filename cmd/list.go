/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")

		// Open the file with read
		todosFile, err := os.OpenFile(CSV_FILENAME, os.O_RDONLY, os.ModePerm)
		if err != nil {
			fmt.Println("No file exits!")
			return

		}
		defer todosFile.Close()

		todos, err := GetTodos(todosFile)

		if err != nil {
			fmt.Println(err)
			return
		}

		ListCsvContents(todos)
	},
}

func makeDividers(word string, dividerCharacter string) string {
	s := make([]string, len(word))
	for i := range s {
		s[i] = dividerCharacter
	}

	return strings.Join(s, "")
}

func ListCsvContents(todos []*Todo) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.TabIndent)
	defer w.Flush()

	// header row
	header_row := "ID\tDescription\tCreated At\tCompleted?"
	fmt.Fprintln(w, header_row)

	// divider row
	dividers := []string{makeDividers("ID", "-"), makeDividers("Description", "-"), makeDividers("Created At", "-"), makeDividers("Completed?", "-")}
	fmt.Fprintln(w, strings.Join(dividers, "\t"))

	// print data as tabbed columns
	for _, todo := range todos {
		fmt.Fprintln(w, todo.String())
	}

}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
