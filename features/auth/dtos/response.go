package dtos

type ResUser struct {
	Name         string `json:"name"`
	UserType     string `json:"user_type`
	AccessToken  string `json:"access_token`
	RefreshToken string `json:"refresh_token`
}
