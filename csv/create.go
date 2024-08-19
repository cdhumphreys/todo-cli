package csv

import (
	"encoding/csv"
	"log"
	"os"
)

func CreateCsvHeaders(file *os.File) (*os.File, error) {
	headers := []string{"id", "description", "createdAt", "completed"}

	w := csv.NewWriter(file)
	// Write any buffered data to underlying writer
	defer w.Flush()

	// Add header row
	if err := w.Write(headers); err != nil {
		log.Fatalln("error writing headers to csv:", err)
		return nil, err
	}

	// Check for other errors
	if err := w.Error(); err != nil {
		log.Fatalln("Error of some kind:", err)
		return nil, err
	}

	return file, nil
}
