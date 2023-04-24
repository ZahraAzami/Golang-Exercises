package main

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadFileCVS(filePath string) [][]string {
	//TODO
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(filePath, err)
	}
	//fmt.Println(records)
	return records
}
