package response

type UserResponse struct {
	ID               uint64 `json:"id"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	IsEmailConfirmed bool   `json:"isEmailConfirmed"`
}
