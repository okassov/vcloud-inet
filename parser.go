package main

import (
	"regexp"
)

type GuestInfo struct {
	ip       string
	prefix   string
	gateway  string
	dns1     string
	dns2     string
	suffix   string
	hostname string
}

func parseOutput(outs []byte, guestinfo *GuestInfo) {
	var ip = regexp.MustCompile(`<P.*="guestinfo.ip".*="(.*)"`)
	var prefix = regexp.MustCompile(`<P.*="guestinfo.netmask".*="(.*)"`)
	var gateway = regexp.MustCompile(`<P.*="guestinfo.gateway".*="(.*)"`)
	var dns1 = regexp.MustCompile(`<P.*="guestinfo.dns1".*="(.*)"`)
	var dns2 = regexp.MustCompile(`<P.*="guestinfo.dns2".*="(.*)"`)
	var suffix = regexp.MustCompile(`<P.*="guestinfo.suffix".*="(.*)"`)
	var hostname = regexp.MustCompile(`<P.*="guestinfo.hostname".*="(.*)"`)

	ip_match := ip.FindStringSubmatch(string(outs))
	prefix_match := prefix.FindStringSubmatch(string(outs))
	gateway_match := gateway.FindStringSubmatch(string(outs))
	dns1_match := dns1.FindStringSubmatch(string(outs))
	dns2_match := dns2.FindStringSubmatch(string(outs))
	suffix_match := suffix.FindStringSubmatch(string(outs))
	hostname_match := hostname.FindStringSubmatch(string(outs))

	if len(ip_match) > 0 {
		guestinfo.ip = ip_match[1]
	}
	if len(prefix_match) > 0 {
		guestinfo.prefix = prefix_match[1]
	}
	if len(gateway_match) > 0 {
		guestinfo.gateway = gateway_match[1]
	}
	if len(dns1_match) > 0 {
		guestinfo.dns1 = dns1_match[1]
	}
	if len(dns2_match) > 0 {
		guestinfo.dns2 = dns2_match[1]
	}
	if len(suffix_match) > 0 {
		guestinfo.suffix = suffix_match[1]
	}
	if len(hostname_match) > 0 {
		guestinfo.hostname = hostname_match[1]
	}
}
