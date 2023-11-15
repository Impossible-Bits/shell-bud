package utils

import (
	"net"
	"time"
)

// IsMachineOnline checks if a machine is online by pinging its IP address.
func IsMachineOnline(ip string) bool {
	timeout := time.Second * 1
	_, err := net.DialTimeout("tcp", ip+":22", timeout)
	return err == nil
}
