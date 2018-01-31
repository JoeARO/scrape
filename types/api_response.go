package types

type ApiResponse struct {
	Data Data `json:"data"`
}

type Data struct {
	Children []Post `json:"children"`
}
