package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-projects/chatroom/server"
)

var (
	addr   = ":2022"
//	banner = `
//    ____              _____
//   |    |    |   /\     |
//   |    |____|  /  \    |
//   |    |    | /----\   |
//   |____|    |/      \  |
//ChatRoom，start on：%s
//`
	banner = "ChatRoom，start on：%s"
)



func main() {
	fmt.Printf(banner, addr)

	server.RegisterHandle()

	log.Fatal(http.ListenAndServe(addr, nil))
}
