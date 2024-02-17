package dto

type Response struct {
	Data    interface{} "json:data"
	Message string      "json:message"
}
