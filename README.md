# KHU Authenticator in Golang

경희대학교 인증서비스를 제공하는 Golang package

## 사용 방법

```go
import "github.com/khu-dev/khu-auth-go/auth"

func main(){
    userInfo, err := auth.Authenciate(당신의 info21 ID, 당신의 info21 PW)
    if err != nil { fmt.Println(err) }
    fmt.Printf("%#v", userInfo)
}
```

## 개발 명령어

```bash
KHU_USERNAME={{MY_USERNAME}} KHU_PASSWORD={{MY_PASSWORD}} go run .
```