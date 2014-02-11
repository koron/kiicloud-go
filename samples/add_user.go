package main

import (
	kc "../kiicloud"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("USAGE: %s {loginName} {password}\n", os.Args[0])
		os.Exit(1)
	}
	loginName := os.Args[1]
	password := os.Args[2]

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

	b, err := c.RegisterUser(loginName, password, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b)
}
