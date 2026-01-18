package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

const (
	longDesc = `
列出所有公司員工通訊錄

	$ slctl contacts

可以使用員工姓名(模糊查詢), 或員工編號(完整查詢)過濾資料

	$ slctl contacts matt
	$ slctl contacts 33

傳入 '--all' 可以查詢包含非 active 的員工通訊錄, e.g. 已離職員工

	$ slctl contacts -a
`
)

var (
	api = "https://support.softleader.com.tw/erp"
)

type contactsCmd struct {
	offline    bool
	verbose    bool
	token      string
	out        io.Writer
	cli        string
	version    string
	name       string // 姓名, 模糊查詢
	id         int    // 員編
	all        bool
	horizontal bool // 水平列表
}

func main() {
	c := contactsCmd{}
	cmd := &cobra.Command{
		Use:   "slctl contacts NAME/ID",
		Short: "view contacts details in SoftLeader organization",
		Long:  longDesc,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if c.offline {
				return fmt.Errorf("can not run the command in offline mode")
			}
			if c.token = os.ExpandEnv(c.token); c.token == "" {
				return fmt.Errorf("require GitHub access token to run the command")
			}
			if len := len(args); len > 0 {
				if len > 1 {
					return errors.New("this command does not accept more than 1 arguments")
				}
				if arg := strings.TrimSpace(args[0]); arg != "" {
					if c.id, err = strconv.Atoi(arg); err != nil {
						c.name = arg
					}
				}
			}
			return c.run()
		},
	}

	c.out = cmd.OutOrStdout()
	c.cli = os.Getenv("SL_CLI")
	c.version = os.Getenv("SL_VERSION")
	c.offline, _ = strconv.ParseBool(os.Getenv("SL_OFFLINE"))
	c.verbose, _ = strconv.ParseBool(os.Getenv("SL_VERBOSE"))

	f := cmd.Flags()
	f.BoolVarP(&c.offline, "offline", "o", c.offline, "work offline, Overrides $SL_OFFLINE")
	f.BoolVarP(&c.verbose, "verbose", "v", c.verbose, "enable verbose output, Overrides $SL_VERBOSE")
	f.StringVar(&c.token, "token", "$SL_TOKEN", "github access token. Overrides $SL_TOKEN")
	f.BoolVarP(&c.all, "all", "a", false, "show all contacts (default shows just active contacts)")
	f.BoolVarP(&c.horizontal, "horizontal", "H", false, "show contacts horizontally")

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func (c *contactsCmd) run() (err error) {
	client := resty.New()
	client.SetDebug(c.verbose)
	resp, err := client.R().
		SetQueryParams(c.queryParams()).
		SetAuthToken(c.token).
		SetHeader("User-Agent", fmt.Sprintf("%s/%s %s/%s", c.cli, c.version, "contacts", ver())).
		Get(fmt.Sprintf("%s/api/user/contacts", api))
	if err != nil {
		return
	}
	if !resp.IsSuccess() {
		return fmt.Errorf(`expected response status code 2xx, but got %v.
Use the '--verbose' flag to see the full stacktrace
`, resp.StatusCode())
	}
	err = print(c.out, resp.Body(), c.horizontal)
	return
}

func print(out io.Writer, data []byte, horizontal bool) (err error) {
	contacts := contacts{}
	if err = json.Unmarshal(data, &contacts); err != nil {
		return fmt.Errorf("unable to unmarshal response: %s", err)
	}
	if len(contacts.Datas) == 0 {
		fmt.Fprintf(out, "No search results")
		return
	}
	if horizontal {
		fmt.Fprintln(out, contacts.horizontalTable())
	} else {
		fmt.Fprintln(out, contacts.verticalTable())
	}
	return
}

func (c *contactsCmd) queryParams() (qp map[string]string) {
	qp = make(map[string]string)
	qp["a"] = strconv.FormatBool(c.all)
	if id := c.id; id > 0 {
		qp["i"] = strconv.Itoa(id)
	} else if name := c.name; name != "" {
		qp["n"] = name
	}
	return
}
