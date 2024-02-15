package models

import "time"

type WorkFlow struct {
	IsRejected    bool   `json:"is_rejected" bson:"is_rejected"`
	IsAccepted    bool   `json:"is_accepted" bson:"is_accepted"`
	WorkflowStage string `json:"workflow_stage" bson:"workflow_stage"`
}

type PhotoMetadata struct {
	ShotDate time.Time   `json:"shot_date" bson:"shot_date"`
	Camera   string      `json:"camera" bson:"camera"`
	Location Geolocation `json:"location" bson:"location"`
	Place    Place       `json:"place" bson:"place"`
	Exposure string      `json:"exposure" bson:"exposure"`
	FNumber  float32     `json:"f_number" bson:"f_number"`
	ISO      int         `json:"iso" bson:"iso"`
}

type PhotoModel struct {
	ID          string        `json:"id" bson:"id"`
	Url         string        `json:"url" bson:"url"`
	Description string        `json:"description" bson:"description"`
	IsPublic    bool          `json:"is_public" bson:"is_public"`
	Metadata    PhotoMetadata `json:"metadata" bson:"metadata"`
	WorkFlow    WorkFlow      `json:"workflow_stage" bson:"workflow_stage"`
}
