package domain

type PasswordService interface {
	HashPassword(string) (string, error)
	ComparePassword(string, string) (bool, error)
}

type TokenService interface {
	GenerateAccessToken(userId string) (string, error)
	GenerateRefreshToken(userId string) (string, error)
	GenerateVerificationToken(userId string) (string, error)
	GeneratePasswordToken(userId string) (string, error)
	ValidateAccessToken(token string) (string, error)
	ValidateRefreshToken(token string) (string, error)
	ValidateVerificationToken(token string) (string, error)
	ValidatePasswordToken(token string) (string, error)
}