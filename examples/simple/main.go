package main

import (
	"fmt"
	"log"

	"github.com/poroping/libdevice42/client"
	"github.com/poroping/libdevice42/client/devices"
)

func main() {
	d42 := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host:     "api.device42.com",
		BasePath: "/",
		Schemes:  []string{"https"},
	})

	results, err := d42.Devices.GetDevices(&devices.GetDevicesParams{}, nil)

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, device := range results.GetPayload().Devices {
		fmt.Println(fmt.Sprintf("%v", device))
	}
}
