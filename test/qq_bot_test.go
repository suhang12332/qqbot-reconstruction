package test

import (
	"fmt"
	"net/url"
	"testing"
)

func LoadConfigs()  {
	
}


func TestLoadConfigs(t *testing.T) {
	unescape, _ := url.QueryUnescape("trojan://i79sd&HbL97s%y@los.pasi.cat:443#los.pasi.cat%3A443")
	fmt.Println(unescape)
}