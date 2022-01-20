package auth

type Auth interface {
	LoginUserName(userName, password string) (string, error)
	// LoginEmail(email, password string) (string, error)
}
