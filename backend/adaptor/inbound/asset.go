package inbound

import (
	"finance-go/adaptor/outbound"
	"fmt"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
)

func AssetList() ([]*outbound.Asset, error) {

	var assets []*outbound.Asset
	res := outbound.DB.Find(&assets)
	if res.Error != nil {
		return nil, res.Error
	}

	return assets, nil
}

func AssetListWithExchangeRate(c *gin.Context) ([]*outbound.Asset, error) {
	currencyType := outbound.CurrencyType(c.Query("currency_type"))

	if !slices.Contains(outbound.CollectionCurrencyTypes, currencyType) {
		return nil, fmt.Errorf("currency_type not vaild")
	}

	list, err := AssetList()
	if err != nil {
		return nil, err
	}

	exchangeRate, err := outbound.GetExchangeRate(time.Now(), currencyType)
	if err != nil {
		exchangeRate, err = outbound.GetExchangeRate(time.Now().Add(-24*time.Hour), currencyType)
		if err != nil {
			return nil, err
		}
	}

	fmt.Printf("%v\n", exchangeRate)

	for _, r := range list {
		r.Currency = r.Currency / exchangeRate[r.CurrencyType]
		r.CurrencyType = currencyType
	}

	return list, nil
}

func AssetCreateOrUpdate(c *gin.Context) error {
	var asset outbound.Asset
	err := c.ShouldBindJSON(&asset)
	if err != nil {
		return err
	}

	res := outbound.DB.Save(&asset)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
