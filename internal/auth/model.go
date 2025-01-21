package auth

type OAuthUser struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Provider string `json:"provider"`
}
