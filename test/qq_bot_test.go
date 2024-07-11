package test

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestLoadConfigs(tt *testing.T) {
	unescape, _ := url.QueryUnescape("trojan://i79sd&HbL97s%y@los.pasi.cat:443#los.pasi.cat%3A443")
	fmt.Println(unescape)
}

type man interface {
	say()
}

type student struct {
}

func (s student) say() {
	fmt.Println("what can i say!")
}

func TestStudent(t *testing.T) {
	registry := make(map[string]reflect.Type)
	registry["student"] = reflect.TypeOf(student{})

	if typ, ok := registry["student"]; ok {
		instance := reflect.New(typ).Interface()
		// [student] is an instance of interface [man]
		if handler, ok := instance.(man); ok {
			say_sth(handler)
		} else {
			fmt.Println("damn")
		}
	}
}

func say_sth(m man) {
	m.say()
}
