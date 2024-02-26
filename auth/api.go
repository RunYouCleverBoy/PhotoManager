package auth

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginWithTokenRequest struct {
	OldToken     string `json:"oldToken"`
	RefreshToken string `json:"refreshToken"`
}
