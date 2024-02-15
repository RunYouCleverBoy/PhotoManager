package models

import "time"

type WorkFlow struct {
	IsRejected    bool   `json:"is_rejected" bson:"is_rejected"`
	IsSelected    bool   `json:"is_selected" bson:"is_selected"`
	WorkflowStage string `json:"workflow_stage" bson:"workflow_stage"`
}

type PhotoMetadata struct {
	ShotDate time.Time   `json:"shot_date"`
	Camera   string      `json:"camera"`
	Location Geolocation `json:"location"`
	Place    Place       `json:"place"`
	Exposure string      `json:"exposure"`
	FNumber  float32     `json:"f_number"`
	ISO      int         `json:"iso"`
}

type PhotoModel struct {
	ID          string        `json:"id"`
	Url         string        `json:"url"`
	Description string        `json:"description"`
	IsPublic    bool          `json:"is_public"`
	Metadata    PhotoMetadata `json:"metadata"`
	WorkFlow    WorkFlow      `json:"workflow_stage"`
}
