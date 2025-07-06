package models

// InputRequest defines the structure for request body
type InputRequest struct {
	RequestJSON    map[string]interface{} `json:"requestJson"`
	RequestMapping map[string]string      `json:"requestMapping"`
}
