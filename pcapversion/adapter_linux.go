// Taken and modified from golang/src/pkg/net

package main

import (
	"errors"
	"net"
)

func localMac(dev string) (net.HardwareAddr, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, inter := range interfaces {
		if inter.Name == dev {
			return inter.HardwareAddr, nil
		}
	}

	return nil, errors.New("Could not find adapter")
}
