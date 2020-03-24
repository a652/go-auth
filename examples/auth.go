package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/a652/go-auth"
	"github.com/a652/go-auth/store"
	"github.com/a652/go-auth/token"
)

var appIDAndSecretMap = make(map[string]string)

func main() {
	appIDAndSecretMap["12345"] = "652899"
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	token := token.GenToken("/auth", "12345", "652899", ts)
	rawURL := fmt.Sprintf("http://charter652.com/auth?appid=12345&token=%s&ts=%s", token, ts)
	fmt.Printf("rawURL: %s\n", rawURL)
	cs := store.NewDefaultCredentialStore(appIDAndSecretMap)
	aat := auth.NewAPIAuthenticator(cs)
	fmt.Println(aat.Auth(rawURL))
}
