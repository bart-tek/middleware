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

	layout := "2001-01-01 00:00:00.00000000 +0200 CEST m=+0.000000000"
	timestamp, err := time.Parse(layout, date)
	checkError("Error while parsing date", err)

	hour := timestamp.Format("12:50:55.10341964")
	day := timestamp.Format("2019-10-11")
	dataToWrite := []string{hour, valeur}

	filePath, err := filepath.Abs("../../data/" + aeroportID + "-" + day + "-" + nature + ".csv")
	checkError("Error in the path", err)

	// Reads the csv file
	file, err := os.Open(filePath)
	checkError("Cannot open file", err)

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	checkError("Cannot read file", err)

	// Add column
	l := len(lines)
	if len(dataToWrite) < l {
		l = len(dataToWrite)
	}
	for i := 0; i < l; i++ {
		lines[i] = append(lines[i], dataToWrite[i])
	}

	// Write the file
	file, err = os.Create(filePath)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(lines)
	checkError("Cannot write data into csv", err)
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal("\n"+message, err)
	}
}
