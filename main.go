package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	f, err := ioutil.ReadFile("./sample-input.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	input := []Input{}
	json.Unmarshal(f, &input)
	taxiTypeList := TaxiTypeList{}
	sampleOutput := []*TaxiAssigned{}
	for _, item := range input {
		if item.Taxi != nil {
			item.Taxi.Available = true
			taxiTypeList[item.Taxi.Type] = append(taxiTypeList[item.Taxi.Type], item.Taxi)
			continue
		}

		if item.TaxiRequest != nil {
			availableTaxi := item.TaxiRequest.GetAvailableTaxi(taxiTypeList)
			if availableTaxi == nil {
				continue
			}
			assignedTaxi := item.TaxiRequest.AssignTaxi(availableTaxi)
			sampleOutput = append(sampleOutput, assignedTaxi)
			modifyTaxiList(taxiTypeList, availableTaxi.ID, true)
			continue
		}

		if item.TripCompletion != nil {
			for _, t := range sampleOutput {
				if t.CustomerID == item.TripCompletion.CustomerID {
					modifyTaxiList(taxiTypeList, t.TaxiID, true)
					break
				}
			}
		}
	}

	bytesData, err := json.MarshalIndent(sampleOutput, "", "\t")
	if err != nil {
		fmt.Printf("Error while marshaling the json: %v", err)
		return
	}

	err = ioutil.WriteFile("./sample-output.json", bytesData, 0644)
	if err != nil {
		fmt.Printf("Error while writing the json: %v", err)
		return
	}
}
