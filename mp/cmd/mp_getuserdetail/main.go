package mp_getuserdetail

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
	openid    = flag.String("o", "", "user openid")
)

func main() {
	flag.Parse()

    if *version {
		fmt.Printf("version %s", "v1")
        return
    }

    //help info
    if *help || *appid == "" || *appsecret == "" || *openid == ""{
        fmt.Printf("Usage: mp_getuserdetail -a appid -s appsecret -o openid\n")
        return
    }

	accessTokenServer := core.NewDefaultAccessTokenServer(*appid, *appsecret, nil)
	wechatClient      := core.NewClient(accessTokenServer, nil)
	detail, err := user.Get(wechatClient,*openid, "")
	if err != nil {
		fmt.Printf("get detail failed, openid:%s err:%v\n", *openid, err)
		return
	}

	fmt.Printf("detail: %v\n", detail)
}
