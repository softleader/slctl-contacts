package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
)

type contactsCmd struct {
	offline bool
	verbose bool
	token   string
}

func main() {
	c := contactsCmd{}
	cmd := &cobra.Command{
		Use:   "contacts",
		Short: "contacts",
		Long:  "The contacts plugin",
		RunE: func(cmd *cobra.Command, args []string) error {
			c.offline, _ = strconv.ParseBool(os.Getenv("SL_OFFLINE"))
			c.verbose, _ = strconv.ParseBool(os.Getenv("SL_VERBOSE"))
			c.token = os.ExpandEnv(c.token)
			return c.run()
		},
	}

	f := cmd.Flags()
	f.StringVar(&c.token, "token", "$SL_TOKEN", "github access token. Overrides $SL_TOKEN")

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func (c *contactsCmd) run() error {
	// use os.LookupEnv to retrieve the specific value of the environment (e.g. os.LookupEnv("SL_TOKEN"))
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, "SL_") {
			fmt.Println(env)
		}
	}
	fmt.Printf("%+v\n", c)
	return nil
}
