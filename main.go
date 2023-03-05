package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/0xAX/notificator"
	"github.com/tsivinsky/bluetoothd/bluetooth"
)

var notify = notificator.New(notificator.Options{
	AppName: "Bluetooth daemon",
})

var (
	interval int
)

func main() {
	flag.IntVar(&interval, "i", 5, "Interval in minutes")
	flag.Parse()

	for {
		devices, err := bluetooth.GetConnectedDevices()
		if err != nil {
			log.Fatal(err)
		}

		for _, device := range devices {
			title := fmt.Sprintf("%s is low on power", device.Name)
			text := fmt.Sprintf("%s has only %d%% of power", device.Name, device.Percentage)

			if device.Percentage <= 30 {
				notify.Push(title, text, "", "critical")
			}
		}

		time.Sleep(time.Duration(interval * int(time.Minute)))
	}
}
