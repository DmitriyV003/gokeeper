package core

type LoginSecret struct {
	Username       string
	Website        string
	Password       string
	AdditionalData string
	ID             int64
	UserID         int64
}

type CardSecret struct {
	ID             int64
	CardholderName string
	Type           string
	ExpireDate     string
	ValidFrom      string
	Number         string
	AdditionalData string
	SecretCode     string
	UserID         int64
}
