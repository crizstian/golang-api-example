package fn

import (
	"github.com/labstack/echo"

	jwt "github.com/dgrijalva/jwt-go"
)

// GetCustomerID returns customer_id from token
func GetCustomerID(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	customerID := claims["customer_id"].(string)
	return customerID
}
