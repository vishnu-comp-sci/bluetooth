package main

import (
	"fmt"
	"time"

	"github.com/muka/go-bluetooth/bluez/profile/adapter"
)

func main() {
	// Get the default Bluetooth adapter
	btAdapter, err := adapter.GetDefaultAdapter()
	if err != nil {
		panic(fmt.Sprintf("Failed to get default adapter: %s", err))
	}

	// Start discovery
	err = btAdapter.StartDiscovery()
	if err != nil {
		panic(fmt.Sprintf("Failed to start discovery: %s", err))
	}

	fmt.Println("Scanning for Bluetooth devices...")

	// Discover devices for a specified duration
	discoveryTime := 10 * time.Second
	time.Sleep(discoveryTime)

	// Stop discovery
	err = btAdapter.StopDiscovery()
	if err != nil {
		panic(fmt.Sprintf("Failed to stop discovery: %s", err))
	}

	// List discovered devices
	devices, err := btAdapter.GetDevices()
	if err != nil {
		panic(fmt.Sprintf("Failed to get devices: %s", err))
	}

	fmt.Println("Discovered Devices:")
	for _, dev := range devices {
		fmt.Printf("Device: %s [%s]\n", dev.Properties.Name, dev.Properties.Address)
	}
}
