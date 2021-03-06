package dto

import "github.com/dgrijalva/jwt-go"

type JwtReq struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
type JwtClaim struct {
	*jwt.StandardClaims
	Username    string `json:"username"`
	RefestToken string `json:"refesh_token"`
	UserID      uint   `json:"user_id"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
}
type JwtConfig struct {
	jwt.StandardClaims
	SigningMethod string `json:"signing_method"`
	PublicKey     string `json:"public_key"`
	PrivateKey    string `json:"private_key"`
}
