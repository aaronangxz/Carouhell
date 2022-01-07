package models

type API struct {
	APIName        string `json:"api_name"`
	APIMethod      string `json:"api_method"`
	APIDescription string `json:"api_description"`
}
