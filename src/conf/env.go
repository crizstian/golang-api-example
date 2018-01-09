package conf

import "os"

// EasyPayHost ...
var EasyPayHost = selectEaspay()

func selectEaspay() string {
	if os.Getenv("EASYPAY") != "" {
		return os.Getenv("EASYPAY")
	}

	return "http://localhost:3000"
}
