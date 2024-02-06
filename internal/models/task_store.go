package models

type InputTaskData struct {
	TagId    int  `json:"tag_id"`
	DueYear  int  `json:"year"`
	DueMonth *int `json:"month"`
	DueDay   *int `json:"day"`
}
