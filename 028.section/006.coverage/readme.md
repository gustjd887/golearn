# cover
go test -cover
go test -coverprofile=c.out
go tool cover -html=c.out : 포함된 커버 소스를 html으로 출력. 브라우저가 지원되는 환경이면 창을 열 수 있을듯...
