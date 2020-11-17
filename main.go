package main

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/secret"
	"github.com/ssst0n3/docker_secret/cert"
	"github.com/ssst0n3/docker_secret/lib"
	"os"
	"strconv"
	"strings"
)

const (
	EnvSecret      = "SECRET"
	EnvDirSecret   = "DIR_SECRET"
	EnvDevelopment = "DEVELOPMENT"
)

var (
	Secret      []string
	Development bool
	DirSecret   string
)

func GetEnv() {
	var err error

	envSecret := os.Getenv(EnvSecret)
	Secret = strings.Split(envSecret, ",")

	DirSecret = os.Getenv(EnvDirSecret)

	envDevelopment := os.Getenv(EnvDevelopment)
	Development, err = strconv.ParseBool(envDevelopment)

	awesome_error.CheckFatal(err)
}

func main() {
	GetEnv()
	for _, e := range Secret {
		if strings.HasPrefix(e, "CERT_") {
			certName := cert.ParseEnv(e)
			awesome_error.CheckFatal(cert.LoadCertificate(certName))
		} else {
			_, _, err := secret.LoadKey(e)
			awesome_error.CheckFatal(err)
		}
	}
	// COPY all secrets into /tmp/secret
	if Development {
		var sourceFilenameList []string
		for _, e := range Secret {
			if strings.HasPrefix(e, "CERT_") {
				caFileName, keyFileName := cert.CertificateFileName(cert.ParseEnv(e))
				sourceFilenameList = append(sourceFilenameList, caFileName)
				sourceFilenameList = append(sourceFilenameList, keyFileName)
			} else {
				sourceFilenameList = append(sourceFilenameList, e)
			}
		}

		err := lib.CopyFiles(sourceFilenameList, DirSecret, "/tmp/secret")
		awesome_error.CheckFatal(err)
	}
}
