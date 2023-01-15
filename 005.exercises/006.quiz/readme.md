# golang.org
공식사이트

# godoc.org
서드파티 사이트

# 단축연산자
* 단축연산자(:=)는 package scope에서는 사용 할 수 없다.
* block scope(func)에서 사용 가능...
* 한번도 예제에서 package scope에서 사용하질 않아서 몰랐는데, 실제 사용해보면 에러남.

```
non-declaration statement outside function body
```

# 좋은 패키지 이름.
짧고(short), 간결하고(concise), 연상적인(evocative)

# go 질문이 필요할때는
https://forum.golangbridge.org

# 관용적 go 코드?
어떤 지역, 사람들 또는 언어에서 말하는 패턴을 말한다. go 프로그래밍에서 '관용적' 코드란 언어 규약을 지키는 코드를 말한다.
(정적 프로그래밍)

# 리턴한 값을 버리고 싶을때는
_ 사용

# ...interface{} 의 의미?
가변 매개변수. 어떤 값이던 얼마든지 입력 가능
옛날 강의 기준으로는 ...interface{}라고 적혀있는거 같지만, 최근 문서를 보면 ...any로 바뀐 것 같은데...?(아래와 같이)
```
func Appendln(b []byte, a ...any) []byte
```

# enpty interface   interface{}
Go에서 모든 값은 ‘빈 인터페이스(empty interface)’ 타입이기도 하며, ‘interface{}’라고 표현한다.
찾아보면 메서드들의 집합을 만들수 있다고 한다.(클래스...같은건가?)
일단 넘어가기로 함.

# go에서는 캐스팅(casting)이라고 하지 않고, 변환(conversion)이라고 한다.
ㅇㅇ

# {}
컬리(curly -> 곱슬...), 브레이스(brace -> 중괄호)
두 단어 모두 사용함

# ()
퍼렌스(parens -> 줄여서 말한 단어인듯, parentheses -> 소괄호)

# []
브라켓(bracket -> 대괄호) ˈbrakit

# expression
expression = operator + operland

# golang spec
매번 google에 검색해서 들어가는데, go 언어 자체 스펙을 볼때 유용해서 즐겨찾기 추가해 둠.
https://go.dev/ref/spec

# fmt.Println()
계속 사용해온 fmt와 Println같은 것도 결국 누군가가 만들어놓은 package(fmt)이며.
identifier(Println)이다.

# interface{}
go에서 모든 값은 type interface{}로 이루어져 있다.
모든 type도 type interface{}로 이루어져 있다.

