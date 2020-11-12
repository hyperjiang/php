package php

import (
	"encoding/binary"
	"net"
	"os"
	"strings"
)

// GetHostByAddr gets the Internet host name corresponding to a given IP address
func GetHostByAddr(ipAddress string) (string, error) {
	names, err := net.LookupAddr(ipAddress)
	if names != nil && len(names) > 0 {
		return strings.TrimRight(names[0], "."), nil
	}
	return "", err
}

// GetHostByName gets the IPv4 address corresponding to a given Internet host name
func GetHostByName(hostname string) (string, error) {
	ips, err := net.LookupIP(hostname)
	if ips != nil {
		for _, v := range ips {
			if v.To4() != nil {
				return v.String(), nil
			}
		}
	}
	return "", err
}

// GetHostByNamel gets a list of IPv4 addresses corresponding to a given Internet host name
func GetHostByNamel(hostname string) ([]string, error) {
	ips, err := net.LookupIP(hostname)
	if ips != nil {
		var ipstrs []string
		for _, v := range ips {
			if v.To4() != nil {
				ipstrs = append(ipstrs, v.String())
			}
		}
		return ipstrs, nil
	}
	return nil, err
}

// GetHostName gets the host name
func GetHostName() (string, error) {
	return os.Hostname()
}

// IP2Long converts a string containing an (IPv4) Internet Protocol dotted address into a long integer
func IP2Long(ipAddress string) uint32 {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return 0
	}
	ipByte := ip.To4()
	if ipByte == nil {
		return 0
	}

	return binary.BigEndian.Uint32(ipByte)
}

// Long2IP converts an long integer address into a string in (IPv4) Internet standard dotted format
func Long2IP(properAddress uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, properAddress)
	ip := net.IP(ipByte)
	return ip.String()
}
