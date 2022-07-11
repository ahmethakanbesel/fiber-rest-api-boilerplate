package dto

type AuthDTO struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
