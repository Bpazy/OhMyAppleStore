package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/Bpazy/oh-my-apple-store/util"
	"github.com/blinkbean/dingtalk"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

type JsonCookie struct {
	Domain         string  `json:"domain"`
	ExpirationDate float64 `json:"expirationDate,omitempty"`
	HostOnly       bool    `json:"hostOnly"`
	HttpOnly       bool    `json:"httpOnly"`
	Name           string  `json:"name"`
	Path           string  `json:"path"`
	SameSite       string  `json:"sameSite"`
	Secure         bool    `json:"secure"`
	Session        bool    `json:"session"`
	StoreId        string  `json:"storeId"`
	Value          string  `json:"value"`
	Id             int     `json:"id"`
}

type Phone struct {
	Name string
	Code string
}

// loopCmd represents the loop command
var loopCmd = &cobra.Command{
	Use:   "loop",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := resty.New()
		c.SetCookieJar(buildCookiejar())
		dk := dingtalk.InitDingTalkWithSecret(util.GetDingTalkToken(), util.GetDingTalkSecret())

		for {
			for _, phone := range util.GetPhones() {
				r, err := c.R().Get("https://www.apple.com.cn/shop/pickup-message-recommendations?mt=compact&searchNearby=true&store=R493&product=" + phone.Code)
				if err != nil {
					fmt.Printf("请求失败: %+v\n", err)
					continue
				}

				s := r.String()
				n := time.Now().String()

				if strings.Contains(s, "并无") {
					fmt.Printf("%s: %s未检测到库存\n", n, phone.Name)
				} else {
					msg := fmt.Sprintf("%s: %s检测到库存，发送通知\n", n, phone.Name)
					fmt.Print(msg)
					if err := dk.SendTextMessage(msg); err != nil {
						fmt.Println("发送钉钉消息失败: " + msg)
					}
				}
			}
			time.Sleep(30 * time.Second)
		}
	},
}

func buildCookiejar() *cookiejar.Jar {
	jar, err := cookiejar.New(nil)
	util.CheckErr(err)

	d, err := ioutil.ReadFile("cookies.json")
	util.CheckErr(err)
	var jCookies []JsonCookie
	util.CheckErr(json.Unmarshal(d, &jCookies))

	var cookies []*http.Cookie
	for _, c := range jCookies {
		cookies = append(cookies, &http.Cookie{
			Name:   c.Name,
			Value:  c.Value,
			Path:   c.Path,
			Domain: c.Domain,
			MaxAge: -1,
		})
	}

	u, err := url.Parse("https://www.apple.com.cn/")
	util.CheckErr(err)
	jar.SetCookies(u, cookies)

	return jar
}

func init() {
	rootCmd.AddCommand(loopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
