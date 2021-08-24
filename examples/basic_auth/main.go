package main

import (
	"fmt"
	"log"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/poroping/libdevice42/client"
	"github.com/poroping/libdevice42/client/devices"
)

func main() {
	d42 := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host:     "api.device42.com",
		BasePath: "/",
		Schemes:  []string{"https"},
	})

	auth := httptransport.BasicAuth(os.Getenv("DEVICE42_USERNAME"), os.Getenv("DEVICE42_PASSWORD"))

	results, err := d42.Devices.GetDevices(&devices.GetDevicesParams{}, auth)

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, device := range results.GetPayload().Devices {
		fmt.Println(fmt.Sprintf("%v", device))
	}
}
