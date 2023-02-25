# println
평소 fmt 패키지를 사용하여 println을 사용하였는데, 그냥 println만 사용해도 되는 듯

# :=
단축 변수 선언은 함수 밖에서 사용할 수 없다.

# const
상수는 한 번 선언되고 할당되면 값을 바꿀 수 없다.

# keywords
25개의 키워드를 가지고 있으며, 변수나 상수 식별자로 선언할 수 없다.
```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

# byte, rune
byte : uint8, 바이트 코드에 사용
* 영어는 unicode 번호 출력
* 그 외 나머지 문자는 2~4바이트 조합으로 해당 문자 번호 출력
rune : uint32, 유니코드 코드포인트에 사용
* UTF-8에서 사용하는 1:1 매핑 가능한 코드포인트 출력

# ``
원시 문자열 리터럴, \n을 문자 그대로 해석한다.

# |=
이진수로 변경 후, 각 자리수에 1이 있으면 1이 된다.
4 = 100
3 =  11
4 |= 3 -> 111

# if
* if문의 조건식은 반드시 boolean 식으로 표현되어야 한다.
* if문에서 조건식 사용 이전에 간단한 문장을 같이 실행할 수 있다. 이렇게 선언된 변수는 if문 안에서만 사용 가능하다.
```
if val := i * 2; val < max {
    println(val)
}
 
// 아래 처럼 사용하면 Scope 벗어나 에러
val++
```

# switch
* switch 뒤에 expression이 없을 수 있음 : 다른 언어는 switch 키워드 뒤에 변수나 expression 반드시 두지만, Go는 이를 쓰지 않아도 된다. 이 경우 Go는 switch expression을 true로 생각하고 첫번째 case문으로 이동하여 검사한다
* case문에 expression을 쓸 수 있음 : 다른 언어의 case문은 일반적으로 리터럴 값만을 갖지만, Go는 case문에 복잡한 expression을 가질 수 있다
* No default fall through : 다른 언어의 case문은 break를 쓰지 않는 한 다음 case로 이동하지만, Go는 다음 case로 가지 않는다
* Type switch : 다른 언어의 switch는 일반적으로 변수의 값을 기준으로 case로 분기하지만, Go는 그 변수의 Type에 따라 case로 분기할 수 있다
## fallthrough
* go의 case 문 끝에는 자동으로 break문이 붙는다. 그래서 switch를 빠져나오는데, 다음 case도 실행하려면 fallthrough를 사용한다.

# for
for 초기값, 조건식, 증감 {...} 의 양식을 따른다.
* go는 반복문이 for밖에 없다.
* 조건식만 쓰는 for문은, 다른 언어에서 while과 같이 동작한다.

## range
for i, v := range 컬렉션 {...} 의 양식을 따른다.

## break, continue, goto
```
package main
func main() {
    var a = 1
    for a < 15 {
        if a == 5 {
            a += a
            continue // for루프 시작으로
        }
        a++
        if a > 10 {
            break  //루프 빠져나옴
        }
    }
    if a == 11 {
        goto END //goto 사용예
    }
    println(a)
 
END:
    println("End")
}
```

### brake Label
```
package main
 
func main() {
    i := 0
 
L1:
    for {
     
        if i == 0 {
            break L1 // L1 위치를 빠져나와 OK를 찍고 마무리.
        }
    }
 
    println("OK")
}
```
# func
* 호출되는 함수가 호출하는 함수의 앞에 있을 필요 없음.

## paas by value, pass by reference
* paas by value : 값을 전달함, 함수 내에서 그 값을 변경하더라도, 그 값을 던진 위치에서 해당 값은 변경되지 않음.
* pass by reference : 주소(&)를 전달함, 함수 내에서 그 값(*)을 변경하면 그 값을 던진 위치에서 해당 값도 변경됨.

## Variadic Function
가변인자함수. ...int 와 같은 형식으로 받는다. 보통 이런 함수에는 []int{v1, v2}... 와 같은 형식으로 args를 전달한다.

## Named Return Parameter
```
func sum(nums ...int) (count int, total int) { // count, total을 반환, 이름을 지정함. return할 파라미터 이름을 지정함.
    for _, n := range nums {
        total += n
    }
    count = len(nums)
    return // 실제 return 문에는 아무 값을 반환하지 않지만, return을 써야됨(에러발생)
}
```

## 일급함수
Go 프로그래밍 언어에서 함수는 일급함수로서 Go의 기본 타입과 동일하게 취급되며, 따라서 다른 함수의 파라미터로 전달하거나 다른 함수의 리턴값으로도 사용될 수 있다.

## Delegate
함수의 원형을 정의하고 함수를 타 메서드에 전달하고 리턴받는 기능을 타 언어에서 흔히 델리게이트(Delegate)라 부른다. Go는 이러한 Delegate 기능을 제공하고 있다.
```
// 원형 정의
type calculator func(int, int) int
 
// calculator 원형 사용
func calc(f calculator, a int, b int) int {
    result := f(a, b)
    return result
}
```

# Closure
Closure는 함수 바깥에 있는 변수를 참조하는 함수값(function value)를 일컫는데, 이때의 함수는 바깥의 변수를 마치 함수 안으로 끌어들인 듯이 그 변수를 읽거나 쓸 수 있게 된다.
```
package main
 
func nextValue() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
 
func main() {
    next := nextValue()
 
    println(next())  // 1
    println(next())  // 2
    println(next())  // 3
 
    anotherNext := nextValue()
    println(anotherNext()) // 1 다시 시작
    println(anotherNext()) // 2
}
```

# array
* 배열의 선언은 "var 변수명 [배열크기] 데이타타입" 과 같이 한다.
* Go에서 배열의 배열크기는 Type을 구성하는 한 요소이다. 즉, [3]int와 [5]int는 서로 다른 타입으로 인식된다.
* 배열이 선언된 후에 각 배열의 요소를 인덱스를 사용하여 읽거나 쓸 수 있다.
```
package main
 
func main() {
    var a [3]int  //정수형 3개 요소를 갖는 배열 a 선언
    a[0] = 1
    a[1] = 2
    a[2] = 3
    println(a[1]) // 2 출력
}
```

```
var a1 = [3]int{1, 2, 3}
var a3 = [...]int{1, 2, 3} //배열크기 자동으로
```

# slice
* 슬라이스에 별도의 길이와 용량을 지정하지 않으면, 기본적으로 길이와 용량이 0 인 슬라이스를 만드는데, 이를 Nil Slice 라 하고, nil 과 비교하면 참을 리턴한다.

## make
make() 함수로 슬라이스를 생성하면, 개발자가 슬라이스의 길이(Length)와 용량(Capacity)을 임의로 지정할 수 있는 장점이 있다.

## sub-slice
* array는 이렇게 못씀... 인덱스를 하나하나 꺼내서 사용해야 하는 듯...?
```
package main
import "fmt"
 
func main() {
    s := []int{0, 1, 2, 3, 4, 5}
    s = s[2:5]  // 부분 슬라이스
    fmt.Println(s) //2,3,4 출력
}
```

## append
내장함수 append()가 슬라이스에 데이타를 추가할 때, 내부적으로 다음과 같은 일이 일어난다. 슬라이스 용량(capacity)이 아직 남아 있는 경우는 그 용량 내에서 슬라이스의 길이(length)를 변경하여 데이타를 추가하고, 용량(capacity)을 초과하는 경우 현재 용량의 2배에 해당하는 새로운 Underlying array (주: 아래 내부구조 참조) 을 생성하고 기존 배열 값들을 모두 새 배열에 복제한 후 다시 슬라이스를 할당한다. 아래 예제는 길이 0/용량 3의 슬라이스에 1부터 15까지의 숫자를 계속 추가하면서 슬라이스의 길이와 용량이 어떻게 변하는 지를 체크하는 코드이다. 이 코드를 실행하면 1~3까지는 기존의 용량 3을 사용하고, 4~6까지는 용량 6을, 7~12는 용량 12, 그리고 13~15는 용량 24의 슬라이스가 사용되고 있음을 알 수 있다.
* slice의 capacity사 늘어나는 케이스가 반복되게 되면, underlying array를 새로 만들고 copy하는 작업이 반복되기 때문에 성능에 영향이 있을 듯.
* 사용할 길이가 어느정도 정해져 있다면, make를 사용하여 애초에 길이를 넉넉히 잡아놓고 사용하는 게 바람직할 것.
```
package main
 
import "fmt"
 
func main() {
    // len=0, cap=3 인 슬라이스
    sliceA := make([]int, 0, 3)
 
    // 계속 한 요소씩 추가
    for i := 1; i <= 15; i++ {
        sliceA = append(sliceA, i)
        // 슬라이스 길이와 용량 확인
        fmt.Println(len(sliceA), cap(sliceA))
    }
 
    fmt.Println(sliceA) // 1 부터 15 까지 숫자 출력 
}
```

## copy
이러한 추가/확장 기능과 더불어, Go 슬라이스는 내장함수 copy()를 사용하여 한 슬라이스를 다른 슬라이스로 복사할 수도 있다. 아래 예제는 3개의 요소를 갖는 소스 슬라이스를 그 2배의 크기 즉 6개를 갖는 타겟슬라이스로 복사하는 예를 보여준다. (주: 아래에서 설명하듯이 슬라이스는 실제 배열을 가리키는 포인터 정보만을 가지므로, 복사를 좀 더 정확히 표현하면, 소스 슬라이스가 갖는 배열의 데이타를 타겟 슬라이스가 갖는 배열로 복제하는 것임)

```
func main() {
    source := []int{0, 1, 2}
    target := make([]int, len(source), cap(source)*2)
    copy(target, source)
    fmt.Println(target)  // [0 1 2 ] 출력
    println(len(target), cap(target)) // 3, 6 출력
}
```

## 슬라이스의 내부구조
http://golang.site/go/article/13-Go-%EC%BB%AC%EB%A0%89%EC%85%98---Slice

참고할 것.(이해하기 쉬운 그림과 설명 있음.)

# map
map[Key타입]Value타입
## nil map
이때 선언된 변수 idMap은 (map은 reference 타입이므로) nil 값을 갖으며, 이를 Nil Map이라 부른다. Nil map에는 어떤 데이타를 쓸 수 없는데, map을 초기화하기 위해 make()함수를 사용할 수 있다 (map 리터럴을 사용할 수도 있는 이는 아래 참조).
```
var idMap map[int]string // nil map 선언, 사용 불가 상태.
```

make() 함수의 첫번째 파라미터로 map 키워드와 [키타입]값타입 을 지정하는데, 이때의 make()함수는 해시테이블 자료구조를 메모리에 생성하고 그 메모리를 가리키는 map value를 리턴한다 (map value는 내부적으로 runtime.hmap 구조체를 가리키는 포인터이다). 따라서 idMap 변수는 이 해시테이블을 가리키는 map을 가리키게 된다.
```
idMap = make(map[int]string) // map 사용 가능하도록 초기화
```

## map literal
```
//리터럴을 사용한 초기화
tickers := map[string]string{
    "GOOG": "Google Inc",
    "MSFT": "Microsoft",
    "FB":   "FaceBook",
}
```

## map 사용
처음 map이 make() 함수에 의해 초기화 되었을 때는, 아무 데이타가 없는 상태이다. 이때 새로운 데이타를 추가하기 위해서는 "map변수[키] = 값" 과 같이 해당 키에 그 값을 할당하면 된다. 예를 들면, 아래 예제에서 키 901 에 Apple을 할당하면 새 해시 키-값 쌍이 추가된다. 만약 키 901의 값이 이미 존재했다면, 추가대신 값만 갱신한다.

```
package main
 
func main() {
    var m map[int]string
 
    m = make(map[int]string)
    //추가 혹은 갱신
    m[901] = "Apple"
    m[134] = "Grape"
    m[777] = "Tomato"
 
    // 키에 대한 값 읽기
    str := m[134]
    println(str)
 
    noData := m[999] // 값이 없으면 nil 혹은 zero 리턴
    println(noData)
 
    // 삭제
    delete(m, 777)
}
```

## for range map
```
package main
 
import "fmt"
 
func main() {
    myMap := map[string]string{
        "A": "Apple",
        "B": "Banana",
        "C": "Charlie",
    }
 
    // for range 문을 사용하여 모든 맵 요소 출력
    // Map은 unordered 이므로 순서는 무작위
    for key, val := range myMap {
        fmt.Println(key, val)
    }
}
```
# package

## main package
일반적으로 패키지는 라이브러리로서 사용되지만, "main" 이라고 명명된 패키지는 Go Compiler에 의해 특별하게 인식된다. 패키지명이 main 인 경우, 컴파일러는 해당 패키지를 공유 라이브러리가 아닌 실행(executable) 프로그램으로 만든다. 그리고 이 main 패키지 안의 main() 함수가 프로그램의 시작점 즉 Entry Point가 된다. 패키지를 공유 라이브러리로 만들 때에는, main 패키지나 main 함수를 사용해서는 안된다.

## 패키지 Import
패키지를 import 할 때, Go 컴파일러는 GOROOT 혹은 GOPATH 환경변수를 검색하는데, 표준 패키지는 GOROOT/pkg 에서 그리고 사용자 패키지나 3rd Party 패키지의 경우 GOPATH/pkg 에서 패키지를 찾게 된다.

GOROOT 환경변수는 Go 설치시 자동으로 시스템에 설정되지만, GOPATH는 사용자가 지정해 주어야 한다. GOPATH 환경변수는 3rd Party 패키지를 갖는 라이브러리 디렉토리나 사용자 패키지가 있는 작업 디렉토리를 지정하게 되는데, 복수 개일 경우 세미콜론(윈도우즈의 경우)을 사용하여 연결한다.

## 패키지 Scope
패키지 내에는 함수, 구조체, 인터페이스, 메서드 등이 존재하는데, 이들의 이름(Identifier)이 첫문자를 대문자로 시작하면 이는 public 으로 사용할 수 있다. 즉, 패키지 외부에서 이들을 호출하거나 사용할 수 있게 된다. 반면, 이름이 소문자로 시작하면 이는 non-public 으로 패키지 내부에서만 사용될 수 있다.

## 패키지 init 함수와 alias
