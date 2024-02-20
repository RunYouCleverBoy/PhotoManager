package database

import (
	"go.mongodb.org/mongo-driver/bson"
)

// bsonBuilder is a struct that represents a BSON builder.
// It is used to construct BSON documents.
type bsonBuilder struct {
	data bson.D
}

func (builder *bsonBuilder) addVal(key string, value interface{}) *bsonBuilder {
	builder.data = append(builder.data, bson.E{Key: key, Value: value})

	return builder
}

func (builder *bsonBuilder) addValIf(key string, condition bool, value interface{}) *bsonBuilder {
	if condition {
		builder.addVal(key, value)
	}
	return builder
}

// addOr appends one or more conditions to the $or operator in the BSON document.
// It takes a variadic parameter of bson.D, representing the conditions to be added.
// The conditions are added to the existing data in the bsonBuilder.
// Returns a pointer to the bsonBuilder for method chaining.
func (builder *bsonBuilder) addOr(conditions ...bson.D) *bsonBuilder {
	builder.data = append(builder.data, bson.E{Key: "$or", Value: conditions})
	return builder
}

// addIfContains adds a key-value pair to the bsonBuilder if the value is not nil.
// It adds the key-value pair as a regex expression in the form of bson.D{{Key: "$regex", Value: value}}.
// The method returns the bsonBuilder to allow for method chaining.
func (builder *bsonBuilder) addIfContains(key string, value *string) *bsonBuilder {
	if value != nil {
		builder.addVal(key, bson.D{{Key: "$regex", Value: *value}})
	}
	return builder
}

// build returns the built BSON document as a pointer to bson.D.
// It returns a pointer to the underlying data of the bsonBuilder.
func (builder *bsonBuilder) build() *bson.D {
	return &builder.data
}
