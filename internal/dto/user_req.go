package dto

type UserReq struct {
	Name string `json:"name"`
	Deps string `json:"deps"`
	Age  int    `json:"age"`
}
