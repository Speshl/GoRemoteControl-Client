package main

import (
	"context"
	"flag"
	"log"

	"github.com/Speshl/GoRemoteControl/client"
	"golang.org/x/sync/errgroup"
)

func main() {
	listJoysticks := flag.Bool("listjoys", false, "List available joysticks")
	showJoyStats := flag.Bool("joystats", false, "Shows states of connected joysticks")

	udpPort := flag.String("joystickport", "1053", "Joystick Port")

	controlDeviceCfg := flag.String("cfg", "./configs/g27.json", "Path to cfg json")

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

		c := client.NewClient(":"+*udpPort, *controlDeviceCfg)
		errorGroup.Go(func() error { return c.RunClient(ctx) })

		err := errorGroup.Wait()
		if err != nil {
			log.Fatalf("Errorgroup had error: %s", err.Error())
		}
	}
}
