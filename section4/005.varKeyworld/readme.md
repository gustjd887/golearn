# 단축선언연산자 사용과 var를 사용한 선언의 차이점?

https://go.dev/ref/spec#Keywords

var는 keyword임

var를 사용하여 변수를 선언하면 함수 밖에서도 선언하여 함수 안에서도 사용 가능

단축연산자를 사용하여 함수 안에서 선언하면, 함수 안에서만 사용 가능

최대한 단축연산자를 사용하여 코드를 작성하고, 변수의 사용 범위를 제한하는 것이 좋음

# var
var = variables = 변수

# var int z
값을 지정하지 않은 변수 z를 지정하면 기본값인 0이 출력됨

# var int z = 2
var를 사용하여 변수 지정하면서 값 지정

# var string g
string type g 변수 지정, default는 ""으로 출력될 것.

