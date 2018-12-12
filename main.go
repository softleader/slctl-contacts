package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	api = "https://support.softleader.com.tw/leave"
)

type contactsCmd struct {
	offline bool
	verbose bool
	token   string
	name    string // 姓名, 模糊查詢
	idno    int    // 員編
	all     bool
}

func main() {
	c := contactsCmd{}
	cmd := &cobra.Command{
		Use:   "contacts NAME/IDNO",
		Short: "View contacts details in SoftLeader organization",
		Long:  longDesc,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if len := len(args); len > 0 {
				if len > 1 {
					return errors.New("this command does not accept more than 1 arguments")
				}
				if arg := strings.TrimSpace(args[0]); arg != "" {
					if c.idno, err = strconv.Atoi(arg); err != nil {
						c.name = arg
					}
				}
			}
			c.offline, _ = strconv.ParseBool(os.Getenv("SL_OFFLINE"))
			c.verbose, _ = strconv.ParseBool(os.Getenv("SL_VERBOSE"))
			c.token = os.ExpandEnv(c.token)
			return c.run()
		},
	}

	f := cmd.Flags()
	f.StringVar(&c.token, "token", "$SL_TOKEN", "github access token. Overrides $SL_TOKEN")
	f.BoolVarP(&c.all, "all", "a", false, "show all contacts (default shows just active contacts)")

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func (c *contactsCmd) run() (err error) {
	var buf bytes.Buffer
	url := fmt.Sprintf("%s/api/user/contacts?%s", api, c.queryString())
	req, err := http.NewRequest("GET", url, &buf)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	if c.verbose {
		fmt.Printf("%s %s\n", req.Method, req.URL)
		fmt.Printf("Header: %v\n", req.Header)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	contacts := contacts{}
	if err = json.NewDecoder(resp.Body).Decode(&contacts); err != nil {
		return fmt.Errorf("unable to unmarshal response from leave service: %s", err)
	}
	if len(contacts.Datas) == 0 {
		fmt.Printf("No search results")
	} else {
		table := uitable.New()
		table.AddRow(contacts.Header...)
		for _, data := range contacts.Datas {
			table.AddRow(data...)
		}
		fmt.Println(table)
	}
	return
}

func (c *contactsCmd) queryString() string {
	qs := make(map[string]string)
	qs["a"] = strconv.FormatBool(c.all)
	if c.idno > 0 {
		qs["i"] = strconv.Itoa(c.idno)
	} else {
		qs["n"] = c.name
	}
	var qss []string
	for k, v := range qs {
		if v != "" {
			qss = append(qss, k+"="+v)
		}
	}
	return strings.Join(qss, "&")
}
