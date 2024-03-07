package models

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

type geoJson struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}
