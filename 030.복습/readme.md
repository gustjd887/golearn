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
