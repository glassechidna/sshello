package main

import (
	"github.com/gliderlabs/ssh"
	gossh "golang.org/x/crypto/ssh"
	"io"
	"log"
	"github.com/davecgh/go-spew/spew"
	"os"
)

type PrintableCertFields struct {
	Serial          uint64
	CertType        uint32
	KeyId           string
	ValidPrincipals []string
	Permissions		gossh.Permissions
}

func main() {
	spew.Config.SortKeys = true

	ssh.Handle(func(s ssh.Session) {
		cert, ok := s.PublicKey().(*gossh.Certificate)
		if !ok {
			io.WriteString(s, "not an ssh certificate!\n")
		} else {
			spew.Fdump(s, printableFields(cert))
		}
	})

	publicKeyOption := ssh.PublicKeyAuth(func(ctx ssh.Context, key ssh.PublicKey) bool {
		return true
	})

	addr := ":2222"
	if envval, found := os.LookupEnv("LISTEN_ADDRESS"); found {
		addr = envval
	}

	log.Fatal(ssh.ListenAndServe(addr, nil, publicKeyOption))
}

func printableFields(c *gossh.Certificate) PrintableCertFields {
	return PrintableCertFields{
		Serial: c.Serial,
		CertType: c.CertType,
		KeyId: c.KeyId,
		ValidPrincipals: c.ValidPrincipals,
		Permissions: c.Permissions,
	}
}
