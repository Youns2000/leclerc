package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
)

// Website is the main function
func Website(c *config) error {
	var d Detail

	//Getting Infos and store them in the Detail structure
	errInfos := getInfos(&d, c)
	if errInfos != nil {
		return errInfos
	}

	//Checking if the product is available
	exist := checkProductExist(&d, c)
	if exist == false {
		b := randomProductChosen(&d, c)
		if b == false {
			return errors.New("product doesn't exist")
		}
	}

	//Adding to Cart
	errATC := atc(c)
	if errATC != nil {
		return errATC
	}
	return nil
}

// Getting Infos function and store them in Detail
func getInfos(d *Detail, c *config) error {
	resp, err := http.Get("https://www.e.leclerc/api/rest/live-api/product-details-by-sku/" + c.ProductSku)
	if err != nil {
		return err
	}

	detailsValue, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(detailsValue, &d)

	return nil
}

//Checking if the product with the given size and color is available and exists
func checkProductExist(d *Detail, c *config) bool {
	for i := 0; i < len(d.Variants); i++ {
		if d.Variants[i].Sku == c.ProductSku {
			if len(d.Variants[i].Offers) == 0 {
				return false
			} else {
				c.ExternalId = d.Variants[i].Offers[0].ExternalId
			}
			return true
		}
	}
	return false
}

//Pick randomly another size (and/or) another color for the product
func randomProductChosen(d *Detail, c *config) bool {
	if len(d.Variants) != 0 {
		c.ExternalId = d.Variants[rand.Intn(len(d.Variants)-1)].Offers[0].ExternalId
		return true
	} else {
		return false
	}
}

// Adding to Cart function
func atc(c *config) error {
	//Setting the body and the header of the POST query
	data := fmt.Sprintf(`{"operationName": "computeLocalCartFromOffers",
    "variables": {
        "offers": [
            {
                "offerId": "%s",
                "productSku": "%s",
                "quantity": 1
            }
        ]
    },"query": "query computeLocalCartFromOffers($offers: [ICartItemIds]) { computeLocalCartFromOffers(offers: $offers) {_id}}"
	}`, c.ExternalId, c.ProductSku)
	body := strings.NewReader(data)

	//Preparing the request
	req, err := http.NewRequest("POST", "https://www.e.leclerc/api/graphql", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	//Sending the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("wrong POST response status code : %d", resp.StatusCode)
	}

	return nil
}
