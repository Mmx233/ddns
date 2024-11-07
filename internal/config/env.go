package config

import (
	"github.com/Mmx233/EnvConfig"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"time"
)

func initEnvConfig() {
	EnvConfig.Load("", &Env)

	if !Env.Ipv4 && !Env.Ipv6 {
		log.Fatalln("both ipv4 and ipv6 ddns is disabled")
	}

	if Env.TTL == 0 {
		Env.TTL = 600
	}

	if Env.STUN == "" {
		Env.STUN = "stun.l.google.com:19302"
	}

	if Env.Timeout == 0 {
		Env.Timeout = 30
	}
	Timeout = time.Duration(Env.Timeout) * time.Second

	HttpClient = &http.Client{
		Transport: &http.Transport{
			Proxy:               http.ProxyFromEnvironment,
			TLSHandshakeTimeout: Timeout,
			DialContext: (&net.Dialer{
				Timeout: Timeout,
			}).DialContext,
		},
		Timeout: Timeout,
	}
}

type _EnvConfig struct {
	Ipv4    bool `config:"omitempty"`
	Ipv6    bool `config:"omitempty"`
	Domain  string
	TTL     int `config:"omitempty"`
	Timeout int `config:"omitempty"`
	Zone    string
	Token   string
	STUN    string `config:"omitempty"`
}

var Env _EnvConfig

var (
	Timeout    time.Duration
	HttpClient *http.Client
)
