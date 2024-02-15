package database

// type GeoJson struct {
// 	Type        string    `bson:"type"`
// 	Coordinates []float64 `bson:"coordinates"`
// }

// func newPoint(longitude, latitude float64) GeoJson {
// 	return GeoJson{
// 		Type:        "Point",
// 		Coordinates: []float64{longitude, latitude},
// 	}
// }

// func fromGeolocation(g models.Geolocation) GeoJson {
// 	return newPoint(g.Longitude, g.Latitude)
// }

// func (g GeoJson) toGeolocation() models.Geolocation {
// 	return models.Geolocation{
// 		Longitude: g.Coordinates[0],
// 		Latitude:  g.Coordinates[1],
// 	}
// }
