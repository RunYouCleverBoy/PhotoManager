package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Geolocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Place struct {
	Name    string   `json:"name" bson:"name"`
	Aliases []string `json:"aliases" bson:"aliases"`
	City    string   `json:"city" bson:"city"`
	Country string   `json:"country" bson:"country"`
}

type geojson struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

func (g Geolocation) MarshalBSON() ([]byte, error) {
	marshalled, err := bson.Marshal(geojson{"Point", []float64{g.Longitude, g.Latitude}})
	return marshalled, err
}

func (g *Geolocation) UnmarshalBSON(data []byte) error {
	unmarshalled := geojson{}
	err := bson.Unmarshal(data, &unmarshalled)
	if err != nil {
		return err
	}
	g.Latitude = unmarshalled.Coordinates[1]
	g.Longitude = unmarshalled.Coordinates[0]
	return nil
}
