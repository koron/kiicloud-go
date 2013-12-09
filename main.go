package main

import (
	"./kii"
	"fmt"
)

func main() {
	conf, err := kii.DefaultConfig()
	if err != nil {
		fmt.Println("kii.DefaultConfig() failed:", err)
		return
	}
	cx, err := kii.NewContext(conf.AppInfo)
	if err != nil {
		fmt.Println("kii.NewContext() failed:", err)
		return
	}
	ax, err := cx.Admin(conf.AdminInfo)
	if err != nil {
		fmt.Println("cx.Admin() failed:", err)
		return
	}
	fmt.Println("cx.Admin() succeeded:", ax)
}
