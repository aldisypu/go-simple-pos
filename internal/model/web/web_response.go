package web

type WebResponse[T any] struct {
	Data   T      `json:"data"`
	Errors string `json:"errors,omitempty"`
}
