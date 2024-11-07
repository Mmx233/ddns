package main

import (
	"context"
	"github.com/Mmx233/ddns/internal/config"
	"github.com/Mmx233/ddns/stun"
	log "github.com/sirupsen/logrus"
)

func RunDDNS(ctx context.Context, network, dnsType string) error {
	ctx, cancel := context.WithTimeout(ctx, config.Timeout)
	defer cancel()

	ip, err := stun.Dial(ctx, network, config.Env.STUN)
	if err != nil {
		return err
	}

	return config.DNS.SetDomainRecord(config.Env.Domain, dnsType, ip.String())
}

func main() {
	ctx := context.TODO()

	if config.Env.Ipv4 {
		err := RunDDNS(ctx, "udp4", "A")
		if err != nil {
			log.Fatalln("run ddns for ipv4 failed:", err)
		}
	}
	if config.Env.Ipv6 {
		err := RunDDNS(ctx, "udp6", "AAAA")
		if err != nil {
			log.Fatalln("run ddns for ipv6 failed:", err)
		}
	}
}
