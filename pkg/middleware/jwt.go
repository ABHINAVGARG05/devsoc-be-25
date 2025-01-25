package middleware

import (
	"fmt"
	"net/http"

	logger "github.com/CodeChefVIT/devsoc-be-24/pkg/logging"
	"github.com/CodeChefVIT/devsoc-be-24/pkg/models"
	"github.com/CodeChefVIT/devsoc-be-24/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Protected() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(utils.Config.JwtSecret),
	})
}

func JWTMiddleware() echo.MiddlewareFunc {
	config := echojwt.Config{
		SigningKey:  []byte(utils.Config.JwtSecret),
		TokenLookup: "cookie:jwt",
		SuccessHandler: func(c echo.Context) {
			reqID := uuid.New().String()
			token := c.Get("user").(*jwt.Token)
			claims := token.Claims.(jwt.MapClaims)

			userId, _ := uuid.Parse(claims["user_id"].(string))
			logger.Infof("handling request with request id %s, uesr id is %s", reqID, userId)

			c.Set("req_id", reqID)
			user, err := utils.Queries.GetUserByID(c.Request().Context(), userId)
			if err != nil {
				logger.Errorf(logger.InternalError, err.Error())
			}

			c.Set("user", user)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			fmt.Println(err)
			if err == echojwt.ErrJWTMissing {
				return c.JSON(http.StatusUnauthorized, &models.Response{
					Status: "fail",
					Data: map[string]string{
						"error": "Missing or malformed JWT",
					},
				})
			}

			return c.JSON(http.StatusUnauthorized, &models.Response{
				Status: "fail",
				Data: map[string]string{
					"error": "Invalid or expired token",
				},
			})
		},
	}

	return echojwt.WithConfig(config)
}
