# Fropilr

A method of managing linux profiles


## Getting Started 



### Download or install package



### Usage

~~~
This is a tool built to automate some of the crazier things
one needs to do when doing gpg. I'm betting there will be more than just
encryption in this but for now, it's pretty basic.

Usage:
  fropilr [command]

Available Commands:
  backup      Backup the current user's info
  decrypt     This decrypts encrypted files sent to the server.
  encrypt     Using a pubkey, this encrypts the message
  get         Get config values
  help        Help about any command
  import      Import all the profile
  list        List Saved Profiles
  remove      Remove the Current Profile
  restore     Restores a Profile from a profile string provided by the user
  server      Server for profile pacakge management
  setpass     Setup Passphrase for PrivKey
  setup       Sets up the gnupg environment
  switch      Switches user profile to selected profile
  testing     The test part of the command out

Flags:
      --config string   config file (default is $HOME/.fropilr.yaml)
  -h, --help            help for fropilr
~~~


### Profiles

####  Setup - new installs

```
fropilr setup
# then answer name, email and passphrase
```

#### Restore - if the profile exists locally

```
fropilr restore <profile string>
```

#### Switch - between profiles, this will also backup the current and save locally.

```
fropilr switch <profile string>
```

#### Uninstall

```
fropilr remove
```

### Encrypt Decrypt

```
fropilr encrypt <plain_text_file> -f <outputfile> -p <full_path_to_pubkey>
fropilr decrypt <encrypted_file> -o/--out <outputfile>
```


## Development

Add this test script to your local environment so you can test different things.
it's ignored in git so it shouldn't copy up to the build server.

```
package cmd

import (
	"fmt"
	"fropilr/config"
	"github.com/spf13/cobra"
)

// testingCmd represents the testing command
var testingCmd = &cobra.Command{
	Use:   "testing",
	Short: "Use this for testing various things",
	Long: `Its a good idea to add this when you need it
and then delete it when you are ready to push to production`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("testing called")
		fmt.Println(config.SystemPasswd)
	},
}

func init() {
	rootCmd.AddCommand(testingCmd)
}
```
