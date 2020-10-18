package cmd

import (
	"fmt"
	"github.com/khu-dev/khu-auth-go/auth"
	"log"
	"os"
)

func init(){
	log.Print("Start Development khu-auth-go")
	defer log.Print("Finish Development khu-auth-go")
	user, err := auth.Authenticate(os.Getenv("KHU_USERNAME"), os.Getenv("KHU_PASSWORD"))
	if err != nil {fmt.Print(err)}
	log.Printf("user %#v", user)
}

