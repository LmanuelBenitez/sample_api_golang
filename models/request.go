package models

type RequestData struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
}
