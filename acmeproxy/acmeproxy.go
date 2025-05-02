package main

import (
	"flag"
	"github.com/hugojosefson/acmeproxy/acmeproxy/server"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

var (
	version = "dev"
)

var port uint
var provider string
var allowedDomainsString string
var allowedDomains []string
var debugLogging bool

func main() {
	Init()

	srv, err := server.NewServer(port, provider, allowedDomains)
	if err != nil {
		log.Fatalf("unable to create server, %v", err)
	}

	srv.Run()
}

func Init() {
	envPrefix := "ACMEPROXY_"
	envPort := os.Getenv(envPrefix + "PORT")
	envProvider := os.Getenv(envPrefix + "PROVIDER")
	envAllowedDomains := os.Getenv(envPrefix + "DOMAINS")

	if envPort != "" {
		p, err := strconv.ParseUint(envPort, 10, 32)
		if err != nil {
			log.Panicf("unable to parse port %s, %v", envPort, err)
		}

		port = uint(p)
	} else {
		port = 9096
	}

	flag.UintVar(&port, "port", port, "server port")
	flag.StringVar(&provider, "provider", envProvider, "DNS provider")
	flag.StringVar(&allowedDomainsString, "domains", envAllowedDomains, "comma seperated list of allowed domains")
	flag.BoolVar(&debugLogging, "debug", false, "enable debug logging")

	flag.Parse()

	allowedDomains = strings.Split(allowedDomainsString, ",")
	for i := range allowedDomains {
		allowedDomains[i] = strings.TrimSpace(allowedDomains[i])
	}

	if debugLogging {
		log.SetLevel(log.DebugLevel)
	}
}
