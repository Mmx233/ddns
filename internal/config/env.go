package config

import (
	"github.com/Mmx233/EnvConfig"
	log "github.com/sirupsen/logrus"
)

func initEnvConfig() {
	EnvConfig.Load("", &Env)

	if !Env.Ipv4 && !Env.Ipv6 {
		log.Fatalln("both ipv4 and ipv6 ddns is disabled")
	}

	if Env.STUN == "" {
		Env.STUN = "stun.l.google.com:19302"
	}
}

type _EnvConfig struct {
	Ipv4  bool
	Ipv6  bool
	TTL   int
	Zone  string
	Token string
	STUN  string
}

var Env _EnvConfig
