


type UserResponse struct {
}

type ClaimResponse struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool
	jwt.RegisteredClaims
}