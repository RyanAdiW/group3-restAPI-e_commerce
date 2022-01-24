package auth

type Auth interface {
	LoginUserName(userName, password string) (string, error)
	GetPasswordByUsername(userName string) (string, error)
	GetIdByUsername(userName string) (int, error)
}
