package main

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/secret"
	"github.com/ssst0n3/docker_secret/cert"
	"os"
	"strings"
)

const (
	EnvSecret = "SECRET"
)

func GetEnv() []string {
	env := os.Getenv(EnvSecret)
	return strings.Split(env, ",")
}

func main() {
	env := GetEnv()
	for _, e := range env {
		if strings.HasPrefix(e, "CERT_") {
			certName := cert.ParseEnv(e)
			awesome_error.CheckFatal(cert.LoadCertificate(certName))
		} else {
			_, _, err := secret.LoadKey(e)
			awesome_error.CheckFatal(err)
		}
	}
}
