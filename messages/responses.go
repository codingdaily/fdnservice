package messages

//GenericResponse structure for generic response
type GenericResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
