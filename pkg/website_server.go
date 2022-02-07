package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	//import yours modules / packages here
)

func RequestWebsite(w http.ResponseWriter, r *http.Request) {

	var c config
	//fmt.Fprintf(w, "Give me a json with the url of a leclerc product\n")

	// Read received request's body and retrieve json
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&c)
	if err != nil {
		fmt.Fprintf(w, "Wrong JSON format \nThe format is : \n{\n    \"url\" : \"https://www.e.leclerc/fp/..........\"\n}\n")
	} else {
		// Get data from JSON and store it in config structure
		err = getDataFromJSON(&c, w)
		if err != nil {
			fmt.Println(err)
		} else {
			// Call main function of the Website
			err = Website(&c)
			if err != nil {
				//fmt.Println(err)
				fmt.Fprintf(w, err.Error())
			} else {
				fmt.Fprintf(w, "Product added to the cart\n")
			}
		}
	}
}

// Halt function to stop the service
func Halt(w http.ResponseWriter, r *http.Request) {

}

// Checking the service function (ping the service to check if it's up)
func CheckService(w http.ResponseWriter, r *http.Request) {
	return
}

// Processes the JSON received by the server and hydrates a config structure to be used by other functions
func getDataFromJSON(c *config, w http.ResponseWriter) error {
	tmp := c.Url
	j := strings.Split(tmp, "/")
	sku := strings.Split(j[len(j)-1], "-")
	c.ProductSku = sku[len(sku)-1]
	c.Slug = strings.Join(sku[:len(sku)-1], "-")

	return nil
}
