package model

type AllItems struct {
	Front []Item `json:"front"`
	Back  []Item `json:"back"`
}
type Item struct {
	Size string `json:"size"`
	Name string `json:"name"`
	Type string `json:"type"`
}
