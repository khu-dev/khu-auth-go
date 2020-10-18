package auth

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func parse(rawUserInfo string) *UserInfo {
	log.Println("Start parse raw user information.")
	defer log.Println("Finish parse raw user information.")
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(rawUserInfo))
	if err !=nil {
		log.Panic(err)
	}
	name := strings.TrimSpace(doc.Find(".user_text01").Text())
	stuNum := strings.TrimSpace(doc.Find(".user_text02").Eq(0).Text())
	dept := strings.TrimSpace(doc.Find(".user_text03").Text())

	user := UserInfo{}
	user.Name = name
	user.StuNumber = stuNum
	user.Department = dept
	return &user
}