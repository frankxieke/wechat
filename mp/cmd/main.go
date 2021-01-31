package main

import (
	"flag"
    "fmt"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/user"
)
var (
    version  = flag.Bool("v", false, "show version information")
    help     = flag.Bool("h", false, "show help information")
    appid     = flag.String("a", "", "mp appid")
    appsecret     = flag.String("s", "", "mp appsecret")
)

func main() {
	flag.Parse()

    if *version {
		fmt.Printf("version %s", "v1")
        return
    }

    //help info
    if *help || *appid == "" || *appsecret == "" {
        fmt.Printf("Usage: mp-user -a appid -s appsecret\n")
        return
    }

	accessTokenServer := core.NewDefaultAccessTokenServer(*appid, *appsecret, nil)
	wechatClient      := core.NewClient(accessTokenServer, nil)
	iter, err := user.NewUserIterator(wechatClient, "")
	if err != nil {
		fmt.Printf("new iterator failed, err:%v", err)
		return
	}

	var result []string

	for iter.HasNext() {
		openids, err := iter.NextPage()
		if err != nil {
			fmt.Printf("new iterator next page failed, err:%v", err)
			return
		}
		result = append(result, openids...)
	}

	fmt.Printf("get total user list num:%d",len(result))
	for i := 0;i < len(result);i ++ {
		fmt.Printf("[%d] : %s\n",result[i])
	}
}
