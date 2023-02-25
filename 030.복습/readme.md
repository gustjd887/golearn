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

## append, copy
