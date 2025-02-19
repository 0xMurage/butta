package authn

type BasicAuthCredentialsDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
