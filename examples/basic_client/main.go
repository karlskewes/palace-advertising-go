package main

import (
	"fmt"
	"log"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	apiclient "github.com/karlskewes/palace-advertising-go/client"
)

func main() {
	fmt.Println("Palace Advertising Client Basic Example")
	fmt.Println("Usage: API_USER=<user> API_PASSWORD=<password> go run main.go")

	// create a runtime transport using the project defaults
	r := httptransport.New(apiclient.DefaultHost, apiclient.DefaultBasePath, apiclient.DefaultSchemes)
	// enable Basic Auth as required by the Palace API
	r.DefaultAuthentication = httptransport.BasicAuth(os.Getenv("API_USER"), os.Getenv("API_PASSWORD"))

	// enable Debug to see all HTTP request/response details
	// r.SetDebug(true)

	// create the API client
	client := apiclient.New(r, strfmt.Default)

	// make an API request to get all Available Properties
	resp, err := client.Developers.V2AvailableProperties(nil)
	if err != nil {
		log.Fatal(err)
	}

	// range over results slice printing each listing
	for listing, property := range resp.Payload {
		fmt.Printf("listing: %v\n", listing)
		fmt.Printf("property: %+v\n", property)
	}

}
