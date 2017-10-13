package main

import (
	"flag"
	"fmt"
	"os"

	common "github.com/dokku/dokku/plugins/common"
	"github.com/dokku/dokku/plugins/config"
	"github.com/dokku/dokku/plugins/config/src/configenv"
)

// get the given entries from the specified environment
func main() {
	args := flag.NewFlagSet("config:get", flag.ExitOnError)
	global := args.Bool("global", false, "--global: use the global environment")
	quoted := args.Bool("quoted", false, "--quoted: get the value quoted")
	args.Parse(os.Args[2:])

	appName, keys := config.GetCommonArgs(*global, args.Args())
	if len(keys) > 1 {
		common.LogFail(fmt.Sprintf("Unexpected argument(s): %v", keys[1:]))
	}
	if len(keys) == 0 {
		common.LogFail("Expected: key")
	}
	if !config.HasKey(appName, keys[0]) {
		os.Exit(1)
	}

	value := config.GetWithDefault(appName, keys[0], "")
	if *quoted {
		fmt.Printf("'%s'", configenv.SingleQuoteEscape(value))
	} else {
		fmt.Printf("%s", value)
	}
}