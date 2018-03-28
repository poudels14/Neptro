// Provides utility method to load required keys from
// environment variables
package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	yaml "gopkg.in/yaml.v2"
)

type Credentials map[string]string
type Services map[string]Credentials

func LoadAllKeys() error {
	homePath, err := homedir.Dir()
	if err != nil {
		return err
	}

	keysFile, err := ioutil.ReadFile(filepath.Join(homePath, "keys_neptro.yaml"))
	if err != nil {
		return err
	}

	var services Services
	err = yaml.Unmarshal(keysFile, &services)
	if err != nil {
		return err
	}

	for _, s := range services {
		for k, v := range s {
			os.Setenv(k, v)
		}
	}

	return nil
}
