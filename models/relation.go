package models

type Relation struct {
	ObjectID string  `json:"objectid"`
	Child    *string `json:"child"`
	Parent   *string `json:"parent"`
}
