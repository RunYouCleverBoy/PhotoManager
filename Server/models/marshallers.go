package models

import "go.mongodb.org/mongo-driver/bson"

func (g *Geolocation) MarshalBSON() ([]byte, error) {
	marshalled, err := bson.Marshal(geoJson{"Point", []float64{g.Longitude, g.Latitude}})
	return marshalled, err
}

func (g *Geolocation) UnmarshalBSON(data []byte) error {
	unmarshalled := geoJson{}
	err := bson.Unmarshal(data, &unmarshalled)
	if err != nil {
		return err
	}
	g.Latitude = unmarshalled.Coordinates[1]
	g.Longitude = unmarshalled.Coordinates[0]
	return nil
}
