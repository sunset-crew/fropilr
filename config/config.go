/*
Copyright © 2021 Joe Siwiak <joe@unherd.info>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os/user"
	"strings"
)

const (
	// RsaBits bit amount
	RsaBits = 4096
	// Version
	Version = "0.1.8"
)

var (
	// SystemPasswd - System password determined by the compiler
	SystemPasswd string
)

// GetSystemPassword Returns Master Password
func GetSystemPassword() []byte {
	return []byte(SystemPasswd)
}

// GetFropileDirectory - gets the root directory for the app
func GetFropileDirectory() string {
	home, _ := GetHomeDir()
	return fmt.Sprintf("%s/.fropilr", home)
}

// GetConfigDirectory Returns Configuration Directory
func GetConfigDirectory() string {
	return fmt.Sprintf("%s/config", GetFropileDirectory())
}

// GetBackupDirectory profile backups
func GetBackupDirectory() string {
	return fmt.Sprintf("%s/store", GetFropileDirectory())
}

// GetDataDirectory app data directory
func GetDataDirectory() string {
	return fmt.Sprintf("%s/data", GetFropileDirectory())
}

// GetConfigFile Returns Configuration File
func GetConfigFile() string {
	return fmt.Sprintf("%s/config.yaml", GetConfigDirectory())
}

// GetDataFile Returns Configuration File
func GetDataFile() string {
	return fmt.Sprintf("%s/data.json", GetConfigDirectory())
}

// GetPrivKey Returns the Location of your Private Key
func GetPrivKey() string {
	return fmt.Sprintf("%s/privkey", GetConfigDirectory())
}

// GetPubKey Returns the Location of your Public Key
func GetPubKey() string {
	return fmt.Sprintf("%s/pubkey", GetConfigDirectory())
}

// GetPassphraseFile Returns the Location of your passphrase file
func GetPassphraseFile() string {
	return fmt.Sprintf("%s/passphrase", GetConfigDirectory())
}

// GetBackupName retrns backup name and stuff
func GetBackupName() string {
	email := viper.GetString("email")
	name := viper.GetString("name")
	return TransformNameEmail(name, email)
}

// TransformNameEmail - transforms the name and email into a profile string
func TransformNameEmail(name, email string) string {
	name = strings.ReplaceAll(name, " ", "-")
	email = strings.Replace(email, "@", "_at_", -1)
	return fmt.Sprintf("%s_%s", name, email)
}

// GetHomeDir returns the home directory as a string
func GetHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, err
}
