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

package gpg

import (
  "os"
  "os/exec"
  "fmt"
  "bytes"
  "strings"
  "errors"
  "github.com/ProtonMail/gopenpgp/v2/helper"
  "github.com/ProtonMail/gopenpgp/v2/crypto"
  "github.com/spf13/viper"
  "fropilr/config"
  "fropilr/utils"
  "log"
)


// ImportSecretkey - import sec key
func ImportSecretkey(passphrase string) string {
    passtr := fmt.Sprintf("--passphrase %s",passphrase)
    a := fmt.Sprintf("gpg --import --batch %s %s",passtr, config.GetPrivKey())
    args := strings.Fields(a)
    cmd := exec.Command(args[0], args[1:]...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
      panic(err)
    }
    return ""
}

// ExportPubkey this exports pubkey
func ExportPubkey(name string) (string,error) {
    a := fmt.Sprintf("gpg --armor %s",fmt.Sprintf("--export %s",name))
    args := strings.Fields(a)
    cmd := exec.Command(args[0], args[1:]...)
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    if err != nil {
      panic(err)
    }
    return string(stdout.Bytes()), errors.New(string(stderr.Bytes()))
}

// GenKey Generates RSA and Elliptic Curve public and private keys
func GenKey(name, email, keyType string, passphrase []byte) string {
    bits := config.RsaBits
    if keyType == "x25519" {
        bits = 0
    }
    rsaKey, err := helper.GenerateKey(name, email, passphrase, keyType, bits)
    utils.Errormsg(err)
    fmt.Println("Private Key Generated")
    return rsaKey
}

// GetPassphrase gets your Passphrase
func GetPassphrase() string {
    armorPassphrase := utils.ReadFile(config.GetPassphraseFile())
    passphrase,_ := DecryptCustomValue(armorPassphrase)
    return passphrase
}

// DecryptArmorFile decrypts and writes out a base64 string into a file
func DecryptArmorFile(path string){
    privkey := utils.ReadFile(config.GetPrivKey())
    armor := utils.ReadFile(path)
    // armor := ReadFile(path)
    passphrase := GetPassphrase()
    decrypted, err := helper.DecryptMessageArmored(privkey, []byte(passphrase), armor)
    if err != nil {
        log.Fatal(err)
    }
    out := viper.GetString("out")
    utils.Write2file(out,decrypted) // detach this
    fmt.Println(decrypted)
    return
}

// DecryptBinaryFile decrypts binary file to string file
func DecryptBinaryFile(path string){
    privkey := utils.ReadFile(config.GetPrivKey())
    cypherbytes := utils.ReadBinaryFile(path)
    // armor := ReadFile(path)
    passphrase := GetPassphrase()
    decrypted, err := DecryptMessageBinary(privkey, []byte(passphrase), cypherbytes)
    if err != nil {
        fmt.Println(err)
        return
    }
    out := viper.GetString("out")
    utils.Write2file(out,decrypted) // detach this
    fmt.Println(decrypted)
    return
}


// EncryptCustomValue encrypts string with custom string
func EncryptCustomValue(message string) (string, error) {
    armor, err := helper.EncryptMessageWithPassword(config.GetSystemPassword(), message)
    return armor, err
}

// DecryptCustomValue decrypts string with custom encryption
func DecryptCustomValue(message string) (string, error) {
    decrypted, err := helper.DecryptMessageWithPassword(config.GetSystemPassword(), message)
    return decrypted, err
}

// EncryptMessageBinary - encrypt in binary format
func EncryptMessageBinary(key, plaintext string) (cipherbytes []byte, err error) {
	var pubKey *crypto.Key
	var pubKeyRing *crypto.KeyRing
	var pgpMessage *crypto.PGPMessage

	var message = crypto.NewPlainMessageFromString(plaintext)

	if pubKey, err = crypto.NewKeyFromArmored(key); err != nil {
		return nil, err
	}

	if pubKeyRing, err = crypto.NewKeyRing(pubKey); err != nil {
		return nil, err
	}

	if pgpMessage, err = pubKeyRing.Encrypt(message, nil); err != nil {
		return nil, err
	}

	cipherbytes = pgpMessage.GetBinary()

	return cipherbytes, nil
}

// DecryptMessageBinary - this decrypts a binary file
func DecryptMessageBinary(
	privKey string, passphrase []byte, cipherbytes []byte,
) (plaintext string, err error) {
	var privKeyObj, privKeyUnlocked *crypto.Key
	var privKeyRing *crypto.KeyRing
	var pgpMessage *crypto.PGPMessage
	var message *crypto.PlainMessage

	if privKeyObj, err = crypto.NewKeyFromArmored(privKey); err != nil {
		return "", err
	}

	if privKeyUnlocked, err = privKeyObj.Unlock(passphrase); err != nil {
		return "", err
	}
	defer privKeyUnlocked.ClearPrivateParams()

	if privKeyRing, err = crypto.NewKeyRing(privKeyUnlocked); err != nil {
		return "", err
	}

	pgpMessage = crypto.NewPGPMessage(cipherbytes)

	if message, err = privKeyRing.Decrypt(pgpMessage, nil, 0); err != nil {
		return "", err
	}

	return message.GetString(), nil
}
