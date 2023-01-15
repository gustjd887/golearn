# git 소스 관리

* git status
* git add .
* git commit -m "message"
* git push

# package 문서

* golang.org > document > go user manual > Package Documentation

# package?
컴퓨터에서 폴더를 사용하여 파일들을 그룹화하여 관리하듯, 프로그램에서는 패키지를 디렉토리와 유사한 구조를 만들어 관리함.

# 패키지 바로가기 주소
* godoc.org/[fmt, html/template, text/template, ...]

# package 페이지의 index에서 사용할 수 있는 func를 찾을 수 있다.
https://pkg.go.dev/fmt#Println

# func 보는 법
func Println(a ...any) (n int, err error)

... = 무한한 개수만큼 매개변수를 받을 것.
any = 어떤 타입의 매개변수라도 받을 것.
 

n int, err error = 글자 수를 int로 리턴, 에러가 있을 경우 err 리턴

# fmt.Println 에서 .의 의미
fmt 패키지로부터 Println 함수를 호출한다.

# 변수 선언 후 사용하지 않으면 에러 발생
```
go run fmt.go 
# command-line-arguments
./fmt.go:16:5: e declared but not used
```

# 리턴값을 무시하는 변수를 사용하고 싶을 때는 _를 사용
_ = underscore
```
	// n, _ := fmt.Println("스트링", 42, true) // err을 버림
	// fmt.Println(n)
```