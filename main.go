package main

import (
	"github.com/keonjeo/i18n"
	"log"
)

func main() {
	result, err := i18n.T("zh-CN", "err.msg.kitty")
	if err != nil {
		log.Fatalf("Fail to T, err: %v", err)
	}
	log.Println("result => ", result)

	result, err = i18n.T("en", "err.msg.hello")
	if err != nil {
		log.Fatalf("Fail to T, err: %v", err)
	}
	log.Println("result => ", result)
}
