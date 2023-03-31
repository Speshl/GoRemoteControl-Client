package main

import (
	"context"
	"flag"
	"log"

	"github.com/Speshl/GoRemoteControl_Client/client"
	"golang.org/x/sync/errgroup"
)

func main() {
	listJoysticks := flag.Bool("listjoys", false, "List available joysticks")
	showJoyStats := flag.Bool("joystats", false, "Shows states of connected joysticks")

	udpPort := flag.String("joystickport", "1053", "Joystick Port")
	host := flag.String("host", "localhost", "IP of udp server")

	controlDeviceCfg := flag.String("cfg", "./configs/g27.json", "Path to cfg json")
	invertESC := flag.Bool("invertesc", false, "Invert ESC value")
	invertSteering := flag.Bool("invertsteer", false, "Invert Steer value")
	trimSteering := flag.Int("strim", 0, "Steering trim: between -15000 and 15000") //5500 for little truck

	flag.Parse()

	if listJoysticks != nil && *listJoysticks {
		_, err := client.GetJoysticks()
		if err != nil {
			log.Fatal(err)
		}
	} else if showJoyStats != nil && *showJoyStats {
		_, err := client.ShowJoyStats()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		errorGroup, ctx := errgroup.WithContext(context.Background())

		c := client.NewClient(*host+":"+*udpPort, *controlDeviceCfg, *invertESC, *invertSteering, *trimSteering)
		errorGroup.Go(func() error { return c.RunClient(ctx) })

		err := errorGroup.Wait()
		if err != nil {
			log.Fatalf("error: %s", err.Error())
		}
	}
}
