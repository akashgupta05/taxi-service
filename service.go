package main

import (
	"math"
)

type TripCompletion struct {
	CustomerID string `json:"customer_id"`
}

type TaxiAssigned struct {
	TaxiID     string `json:"id"`
	CustomerID string `json:"customer_id"`
}

type Taxi struct {
	ID        string   `json:"id"`
	Type      string   `json:"type"`
	Location  Location `json:"location"`
	Available bool
}

type Location struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Input struct {
	Taxi           *Taxi           `json:"taxi"`
	TaxiRequest    *TaxiRequest    `json:"taxi_request"`
	TripCompletion *TripCompletion `json:"trip_completion"`
}

type TaxiRequest struct {
	CustomerID    string   `json:"customer_id"`
	PreferredType string   `json:"preferred_type"`
	SecondaryType string   `json:"secondary_type"`
	Location      Location `json:"location"`
	Destination   Location `json:"destination"`
}

type TaxiTypeList map[string][]*Taxi

func (tr *TaxiRequest) GetMinDistanceTaxi(taxiList []*Taxi) (*Taxi, float64) {
	minDistance := 10000.0
	var availableTaxi *Taxi
	for _, taxi := range taxiList {
		if taxi.Available {
			if distance := getdistance(tr.Location.X, tr.Location.Y,
				taxi.Location.X, taxi.Location.Y); distance < minDistance {
				minDistance = distance
				availableTaxi = taxi
			}
		}
	}
	return availableTaxi, minDistance
}

func (tr *TaxiRequest) GetAvailableTaxi(ttl TaxiTypeList) *Taxi {
	if len(ttl[tr.PreferredType]) > 0 {
		if t, _ := tr.GetMinDistanceTaxi(ttl[tr.PreferredType]); t != nil {
			return t
		}
	}

	if len(ttl[tr.SecondaryType]) > 0 {
		if t, _ := tr.GetMinDistanceTaxi(ttl[tr.SecondaryType]); t != nil {
			return t
		}
	}
	minDistance := 1000.0
	var availableTaxi *Taxi
	for key, taxiList := range ttl {
		if key == tr.PreferredType || key == tr.SecondaryType {
			continue
		}

		if t, d := tr.GetMinDistanceTaxi(taxiList); t != nil {
			if t != nil && d < minDistance {
				minDistance = d
				availableTaxi = t
			}
		}
	}

	return availableTaxi
}

func (tr *TaxiRequest) AssignTaxi(taxi *Taxi) *TaxiAssigned {
	return &TaxiAssigned{CustomerID: tr.CustomerID, TaxiID: taxi.ID}
}

func getdistance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}

func modifyTaxiList(ttl TaxiTypeList, taxiID string, available bool) {
	for key := range ttl {
		for i := range ttl[key] {
			if ttl[key][i].ID == taxiID {
				ttl[key][i].Available = available
				return
			}
		}
	}
}
