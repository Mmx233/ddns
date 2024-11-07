package main

import (
	"context"
	"github.com/Mmx233/ddns/internal/config"
	"github.com/Mmx233/ddns/stun"
	log "github.com/sirupsen/logrus"
	"net"
)

func RunDDNS(ctx context.Context, network, dnsType string) (net.IP, error) {
	ctx, cancel := context.WithTimeout(ctx, config.Timeout)
	defer cancel()

	ip, err := stun.Dial(ctx, network, config.Env.STUN)
	if err != nil {
		return nil, err
	}

	return ip, config.DNS.SetDomainRecord(config.Env.Domain, dnsType, ip.String())
}

func main() {
	ctx := context.TODO()

	if config.Env.Ipv4 {
		ip, err := RunDDNS(ctx, "udp4", "A")
		if err != nil {
			log.Errorln("run ddns for ipv4 failed:", err)
		} else {
			log.Infof("successfully set %s A record to %s", ip.String(), config.Env.Domain)
		}
	}
	if config.Env.Ipv6 {
		ip, err := RunDDNS(ctx, "udp6", "AAAA")
		if err != nil {
			log.Errorln("run ddns for ipv6 failed:", err)
		} else {
			log.Infof("successfully set %s AAAA record to %s", ip.String(), config.Env.Domain)
		}
	}
}
