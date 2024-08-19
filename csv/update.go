package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func AppendToCsv(file *os.File, newTodo Todo) (*os.File, error) {

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err := writer.Write(newTodo.String())

	if err != nil {
		log.Fatalln("Error writing to CSV")
		return nil, err
	}

	fmt.Println("Data appended to csv file successfully.")

	return file, nil
}
