package util

import (
	"github.com/spf13/viper"
)

func GetDingTalkToken() string {
	return viper.GetString("dingtalk.token")
}

func GetDingTalkSecret() string {
	return viper.GetString("dingtalk.secret")
}

type Phone struct {
	Name string
	Code string
}

func GetPhones() []*Phone {
	var phones []*Phone
	CheckErr(viper.UnmarshalKey("phones", &phones))

	return phones
}
