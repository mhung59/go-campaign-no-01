package dto

type UserRes struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Deps string `json:"deps"`
	Age  int    `json:"age"`
}
