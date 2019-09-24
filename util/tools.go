package util

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
	"regexp"
	"strings"
)

func ToHyphenateCamel(str string) string {
	reg := regexp.MustCompile(`\B([A-Z])`)
	return reg.ReplaceAllString(str, "-$1")
}

func XenvDecode(str string) (xenv map[string]string, err error) {
	bty, err := base64.StdEncoding.DecodeString(str)
	var _xenv Xenv
	if err = json.Unmarshal(bty, &_xenv); err != nil {
		return
	}
	elem := reflect.ValueOf(&_xenv).Elem()
	relType := elem.Type()
	xenv = make(map[string]string)
	for i := 0; i < relType.NumField(); i++ {
		xenv[relType.Field(i).Name] = elem.Field(i).String()
	}
	return
}

func Ucwords(s string) (str string) {
	for _, key := range strings.Split(s, "_") {
		str += strings.ToUpper(string(key[0])) + key[1:]
	}
	return
}

type Xenv struct {
	SiteId      string `json:"siteid"`
	FromId      string `json:"fromid"`
	UserId      string `json:"user_id"`
	GuniqId     string `json:"guniqid"`
	AppType     int    `json:"app_type"`
	UtmSource   string `json:"utm_source"`
	UtmMedium   string `json:"utm_medium"`
	UtmCampaign string `json:"utm_campaign"`
	UtmContent  string `json:"utm_content"`
	UtmTerm     string `json:"utm_term"`
	UtmKey      string `json:"utm_key"`
	Pp          string `json:"pp"`
	Pd          string `json:"pd"`
	Pk          string `json:"pk"`
	UserIds     string `json:"userids"`
	Tracker     string `json:"tracker"`
	Channel     string `json:"channel"`
}
