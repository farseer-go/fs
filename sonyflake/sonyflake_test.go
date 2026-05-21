package sonyflake

import (
	"net"
	"testing"
)

func TestPrivateIPv4FallsBackToPublicIPv4(t *testing.T) {
	ip, err := privateIPv4(func() ([]net.Addr, error) {
		return []net.Addr{
			&net.IPNet{IP: net.ParseIP("130.118.128.21"), Mask: net.CIDRMask(24, 32)},
		}, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if ip.String() != "130.118.128.21" {
		t.Fatalf("expected public IPv4 fallback, got %s", ip.String())
	}
}

func TestPrivateIPv4PrefersPrivateIPv4(t *testing.T) {
	ip, err := privateIPv4(func() ([]net.Addr, error) {
		return []net.Addr{
			&net.IPNet{IP: net.ParseIP("130.118.128.21"), Mask: net.CIDRMask(24, 32)},
			&net.IPNet{IP: net.ParseIP("192.168.3.45"), Mask: net.CIDRMask(24, 32)},
		}, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if ip.String() != "192.168.3.45" {
		t.Fatalf("expected private IPv4 first, got %s", ip.String())
	}
}
