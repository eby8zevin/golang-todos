package middleware

import (
	"net/http"
	"strings"

	"github.com/eby8zevin/golang-todos/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if ctx.Path() == "/auth/register" || ctx.Path() == "/auth/login" {
			return next(ctx)
		}

		authHeader := ctx.Request().Header.Get("Authorization")
		if authHeader == "" {
			return ctx.String(http.StatusUnauthorized, "token is empty")
		}

		authArr := strings.Split(authHeader, " ")
		if len(authArr) != 2 {
			return ctx.String(http.StatusUnauthorized, "token is invalid")
		}

		var tokenClaim model.AuthClaimJwt
		token, err := jwt.ParseWithClaims(authArr[1], &tokenClaim, func(t *jwt.Token) (interface{}, error) {
			return []byte("TEST"), nil
		})
		if err != nil {
			return ctx.String(http.StatusUnauthorized, err.Error())
		}
		if !token.Valid {
			return ctx.String(http.StatusUnauthorized, "token is not valid")
		}

		ctx.Set("USER", tokenClaim)
		return next(ctx)
	}
}