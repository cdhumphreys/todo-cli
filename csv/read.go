package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
)

func ReadCsv(file *os.File) ([][]string, error) {

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	data, err := reader.ReadAll()

	if err != nil {
		log.Fatalln("Could not read contents of todos.csv", err)
		return nil, err
	}

	return data, nil
}

func ListCsvContents(data [][]string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.AlignRight|tabwriter.Debug)
	defer w.Flush()

	// header row
	fmt.Fprintln(w, strings.Join(data[0], " \t "))
	for _, row := range data[1:] {
		// set time to/from text for date time
		time, err := time.Parse(TimeFormat, row[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		row[2] = timediff.TimeDiff(time)
		fmt.Fprintln(w, strings.Join(row, " \t "))
	}

}
