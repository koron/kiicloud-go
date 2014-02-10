package main

import (
	kc "./kiicloud"
	"fmt"
)

func main() {
	conf, err := kc.DefaultConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	c, err := conf.NewAdminClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := c.RegisterUser("foo", "bar")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b)
}
