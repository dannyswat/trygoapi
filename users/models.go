package users

type LoginStatus string

const (
	Success      LoginStatus = "success"
	InvalidCred  LoginStatus = "invalid"
	AcctLocked   LoginStatus = "account_locked"
	AcctDisabled LoginStatus = "account_disabled"
	Need2FA      LoginStatus = "proceed_2fa"
)

type LoginModel struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResult struct {
	AccessToken string      `json:"token"`
	Status      LoginStatus `json:"status"`
}
