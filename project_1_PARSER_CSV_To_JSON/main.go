package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Vehicle struct {
	Model       string
	VehicleType string
	Year        int
	MaxSpeed    int
}

func main() {
	csvFile, err := os.Open("./testData.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var veh Vehicle
	var vehArray []Vehicle

	for _, each := range csvData {
		veh.Model = each[0]
		veh.VehicleType = each[1]
		veh.Year, _ = strconv.Atoi(each[2])
		veh.MaxSpeed, _ = strconv.Atoi(each[3])
		vehArray = append(vehArray, veh)
	}

	// Convert to JSON
	jsonData, err := json.Marshal(vehArray)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))

	jsonFile, err := os.Create("./parsedData.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
}
