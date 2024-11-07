package config

import (
	"github.com/Mmx233/ddns/cloudflare"
	log "github.com/sirupsen/logrus"
)

var DNS *cloudflare.DnsProvider

func initDNS() {
	var err error
	DNS, err = cloudflare.New(Env.TTL, cloudflare.Cloudflare{
		Zone:  Env.Zone,
		Token: Env.Token,
	}, HttpClient)
	if err != nil {
		log.Fatalln("init dns provider failed:", err)
	}
}
