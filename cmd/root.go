/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

type Todo struct {
	ID          int    `csv:"id"`
	Description string `csv:"description"`
	CreatedAt   string `csv:"created_at"`
	Completed   bool   `csv:"completed"`
}

func convertTimeToHumanReadable(datetime string) string {
	// set time to/from text for date time
	time, err := time.Parse(TimeFormat, datetime)
	if err != nil {
		fmt.Println(err)
		// default to the original string as a backup
		return datetime
	}
	humanReadableTime := timediff.TimeDiff(time)

	return humanReadableTime
}

func (todo Todo) String() string {
	humanReadableTime := convertTimeToHumanReadable(todo.CreatedAt)
	var values = []string{strconv.Itoa(todo.ID), todo.Description, humanReadableTime, strconv.FormatBool(todo.Completed)}

	return strings.Join(values, " \t ")
}

const TimeFormat = time.RFC3339

const CSV_FILENAME = "todos.csv"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todo-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func GetTodos(file *os.File) ([]*Todo, error) {
	// Todos slice
	todos := []*Todo{}

	// gocsv.UnmarshalFile alters the slice above, so can combo the call & error
	if err := gocsv.UnmarshalFile(file, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}
