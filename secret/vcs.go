package secret

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	vault "github.com/hashicorp/vault/api"
)

type Environment struct {
	VcsIdRsa         string
	VaultToken       string
	VaultTlsCert     string
	Home             string
	VaultAddress     string
	VcsSecretName    string
	VaultTlsCertPath string
}

func NewEnvironment() *Environment {
	var env = &Environment{}
	env.VaultToken = os.Getenv("VAULT_TOKEN")
	env.Home = os.Getenv("HOME")
	env.VaultAddress = os.Getenv("VAULT_ADDRESS")
	env.VcsSecretName = os.Getenv("VCS_SECRET_NAME")
	env.VaultTlsCertPath = os.Getenv("VAULT_TLS_CERT_PATH")

	config := vault.DefaultConfig()
	// VAULT_ADDRESS env variable should be set up

	f, err := os.Open(fmt.Sprintf("%s/%s", env.Home, env.VaultTlsCertPath))
	if err != nil {
		log.Fatalf("secret: %s", err)
	}

	buf := bytes.Buffer{}
	buf.ReadFrom(f)
	buf.WriteString(env.VaultTlsCert)

	VaultTlsConfig := &vault.TLSConfig{CAPath: fmt.Sprintf("%s/%s", env.Home, env.VaultTlsCertPath), TLSServerName: env.VaultAddress}
	config.ConfigureTLS(VaultTlsConfig)
	vaultclient, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("secret error: %s", err)
	}

	vaultclient.SetToken(env.VaultToken)

	secret, err := vaultclient.Logical().Read(env.VcsSecretName)
	if err != nil {
		log.Fatalf("secret: getting vault client: %T %#v\n", err)

	}

	m, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		log.Fatalf("%T %#v\n", secret.Data["data"], secret.Data["data"])
	}

	//getting the key name from vault path for VCS key:
	pathslice := strings.Split(env.VcsSecretName, "/")
	key := pathslice[len(pathslice)-1]
	env.VcsIdRsa = fmt.Sprintf("%s", m[key])

	fmt.Println(env)
	return env

}
