package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ekachaikeaw/ecomm/token"
)

func GetAuthMiddlewareFunc(tokenMaker *token.JWTMaker) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// read the authorization header
			// verify the token
			verifyClaimsFromAuthHeader(r, tokenMaker)	
			// pass the payload/claims down the context
		})
	}
}

func verifyClaimsFromAuthHeader(r *http.Request, tokenMaker *token.JWTMaker) (*token.UserClaims, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("authorization header is missing")
	}

	fields := strings.Fields(authHeader)
	if len(fields) != 2 || fields[0] != "Bearer" {
		return nil, fmt.Errorf("invalid authorization header")
	}

	token := fields[1]
	claims, err := tokenMaker.VerifyToken(token)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	return claims, nil
}