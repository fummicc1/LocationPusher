package main

import (
	"fmt"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/sideshow/apns2/payload"
)

func pushNotification() {
	cert, err := certificate.FromP12File("pusher.p12", "")

	fmt.Printf("start.")

	if err != nil {
		panic(err)
	}
	fmt.Printf("cert created.")
	notification := &apns2.Notification{}
	notification.DeviceToken = "2c00793251098ed2349830f8f3474167ec903d2b4b1a4dd65a7c900803c08028"
	notification.Topic = "com.fumiya.LocationShareAppWithAPNS"
	notification.Payload = payload.NewPayload().ContentAvailable()
	client := apns2.NewClient(cert).Development()
	res, err := client.Push(notification)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
	wg.Done()
}
