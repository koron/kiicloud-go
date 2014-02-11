package main

import (
	kc "../kiicloud"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("USAGE: %s {loginName}\n", os.Args[0])
		os.Exit(1)
	}
	loginName := os.Args[1]

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

	b, err := c.UnregisterUser(loginName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b)
}
