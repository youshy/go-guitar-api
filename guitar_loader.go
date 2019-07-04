package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func LoadGuitarsCsv() []Guitar {
	var guitars []Guitar
	csvFile, _ := os.Open("guitars.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		strings, err := strconv.Atoi(line[1])
		pickupsno, err := strconv.Atoi(line[5])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		guitars = append(guitars, Guitar{
			Name:         line[0],
			Strings:      strings,
			BodyWood:     line[2],
			NeckWood:     line[3],
			FingerWood:   line[4],
			PickupsNo:    pickupsno,
			NeckPickup:   line[6],
			BridgePickup: line[7],
			Bridge:       line[8],
		})
	}
	return guitars
}
