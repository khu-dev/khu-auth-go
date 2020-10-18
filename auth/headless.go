package auth

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
)

func headless(usename, password string) (string, error){
	log.Println("Start headless authentication request.")
	defer log.Println("Finish headless authentication request.")
	//this return raw user info
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	//var res string
	//var image []byte
	var rawUserInfo string
	var evalRes interface{}
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://info21.khu.ac.kr/com/LoginCtr/login.do?sso=ok`),
		chromedp.WaitReady(`input[id='userId']`, chromedp.ByQuery),
		chromedp.WaitReady(`input[id='userPw']`, chromedp.ByQuery),
		chromedp.SetValue(`#userId`, usename, chromedp.ByQuery),
		chromedp.SetValue(`#userPw`, password, chromedp.ByQuery),
		chromedp.Evaluate("loginCheck()", &evalRes),
		//chromedp.InnerHTML("body", &res),
		chromedp.WaitReady(".user_text_box", chromedp.ByQuery),
		chromedp.OuterHTML(".user_text_box", &rawUserInfo, chromedp.ByQuery),
		//chromedp.CaptureScreenshot(&image),
		)
	if err!= nil{
		log.Println(err)
		return "", err
	}
	fmt.Print("rawUserInfo", rawUserInfo)
	//fmt.Print(res)
	//err  = ioutil.WriteFile("screen.png", image, 777)
	//if err!=nil {
	//	log.Print(err)
	//  return "", err
	//}
	return rawUserInfo, nil
}

