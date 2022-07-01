package models

//Person model
type Person struct {
	ObjectID string  `json:"objectid"`
	Name     *string `json:"name"`
}

//PersonTree model
type PersonTree struct {
	ObjectID  string    `json:"objectid"`
	Name      *string   `json:"name"`
	Relations []*Person `json:"relations"`
}
