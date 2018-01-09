package ctrls

import (
	"easycast/src/db"
	"easycast/src/errs"
	"easycast/src/fn"
	"easycast/src/models"

	"github.com/labstack/echo"
)

// This file is for demostration only

// TotalMain ...
func TotalMain(c echo.Context) (echo.Map, errs.Custom) {
	Db := db.New()
	defer Db.Close()

	// get account_id from user token
	accountID := fn.GetAccountID(c)

	total := 0

	// this is an example of how to implement a switch in golang
	switch item := c.Param("item"); item {
	case "live":
		asset := new(models.Asset)
		total = asset.GetTotalAsset(Db, accountID, 0) // 0 => live
	case "vod":
		asset := new(models.Asset)
		total = asset.GetTotalAsset(Db, accountID, 1) // 1 => vod
	case "playlist":
		asset := new(models.Asset)
		total = asset.GetTotalAsset(Db, accountID, 2) // 1 => vod
	case "user":
		u := new(models.User)
		total = u.GetTotalUser(Db, accountID) // 1 => vod
	case "categories":
		c := new(models.Categories)
		total = c.GetTotalCategories(Db, accountID)
	case "instances":
		lci := new(models.LivechannelsInstances)
		total = lci.GetTotalInstances(Db, accountID)
	case "invoices":
		customerID := fn.GetCustomerID(c)
		cc := new(models.CustomersCharges)
		total = cc.GetTotalInvoices(Db, customerID)
	}

	res := echo.Map{
		"total":   total,
		"success": true,
	}

	return res, nil
}
