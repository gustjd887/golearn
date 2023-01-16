# 정수
소수점 아래 숫자가 없음

# 소수
수수점 아래 숫자가 있음

# numeric type
https://go.dev/ref/spec#Numeric_types
여러가지 타입이 있지만, 원하는 경우에 찾아서 사용

# type 선택
작은 자료형만 선택해서 효율적으로 관리할 수 있음

# int, float64
두개 사용할 수 있다고 함.

# uint8, int8
uint : 부호가 없음
int : 부호가 있음

# int, uint 최적화
컴파일 시, 사용할 CPU 아키텍처(32, 64)에 따라 적합한 타입을 선택하여 컴파일한다고 함

# alias
byte = uint8
rune(문자라는 뜻) = int32 (UTF-8을 위한 4byte)

type byte로 사용 가능

type byte로 지정해도 Printf로 확인해보면 결과물은 uint8로 찍힘
rune도 마찬가지일듯

# runtime package
현재 사용중인 OS의 종류나 아키텍처 출력 가능한 패키지

# 결론
일반적으로는 단축연산자나 int, float64 두가지만 사용하면 됨.