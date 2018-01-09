package fn

import (
	"github.com/labstack/echo"

	jwt "github.com/dgrijalva/jwt-go"
)

// obtener el account_id desde el token
func GetAccountID(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	accountID := int(claims["account_id"].(float64))
	return accountID
}
