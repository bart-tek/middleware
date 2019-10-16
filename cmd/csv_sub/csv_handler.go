package main

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"time"
)

// WriteCsv is used to write date in a csv file
//
// @params date (string), aeroportID (string), capteurID (string), nature (string), valeur (string)
//
// @return void
//
func WriteCsv(date string, aeroportID string, capteurID string, nature string, valeur string) {

	layout := "2006-01-02 15:04:05 -0700 MST"
	timestamp, err := time.Parse(layout, date)
	checkError("Error while parsing date", err)

	hour := timestamp.Format("15:04:05")
	day := timestamp.Format("2006-01-02")
	dataToWrite := [][]string{
		[]string{hour, valeur},
	}

	filePath, err := filepath.Abs("../../data/" + aeroportID + "-" + day + "-" + nature + ".csv")
	checkError("Error in the path", err)

	// Reads the csv file if it exists
	if fileExists(filePath) {
		appendData(filePath, dataToWrite[0])
	} else {
		createData(filePath, dataToWrite)
	}

}

// fileExists verify that a file exists
//
// @params filePath (string)
//
// @return bool
//
func fileExists(filePath string) bool {
	ret := true

	_, err := os.Open(filePath)
	if err != nil {
		ret = false
	}
	return ret
}

// appendData functions appends data to a csv file
//
// @params filePath (string), dataToWrite ([]string)
//
// @return void
//
func appendData(filePath string, dataToWrite []string) {
	file, err := os.Open(filePath)

	reader := csv.NewReader(file)
	reader.TrailingComma = true

	lines, err := reader.ReadAll()
	checkError("Cannot read file", err)

	// Add column
	lines = append(lines, dataToWrite)

	// Write the file
	file, err = os.Create(filePath)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)

	for _, line := range lines {
		err = writer.Write(line)
		checkError("Cannot write data into csv", err)
	}

	writer.Flush()
}

// createData creates a new csv file with data in it
//
// @params filePath (string), dataToWrite ([]string)
//
// @return void
//
func createData(filePath string, dataToWrite [][]string) {
	file, err := os.Create(filePath)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(dataToWrite)
	checkError("Cannot write to file", err)
}

// checkError is the basic error handler function to use
//
// @params message (string), err (error)
//
// @return void
//
func checkError(message string, err error) {
	if err != nil {
		log.Fatal("\n"+message, err)
	}
}
