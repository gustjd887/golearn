# ` (백틱)
``으로 감싸면, 원시 문자열 리터럴 가능

# type string은 byte slice로 되어있음
int를 배울 때, byte는 uint8의 alias로 배웠고, 그 alias를 쓰는 이유는 UTF-8을 담을 수 있는 크기이기 때문이라고 배웠음

# 한글의 경우
 3byte = unicode = utf-8

 # 영어의 경우
 1byte = ascii 호환 = utf-8

# UTF-8
10진수(byte) : 코드 포인트(rune) : 16진수(hex) -> 각 매칭되는 문자들이 있음(code scheme)

# rune
rune = int32 = 4byte = UTF-8 표현을 위한 alias

# codepoint
문자열을 대표하는 정수값

유니코드 인코딩 방식인 UTF-8은, 문자열을 표시하는 방법이 상이함.
UTF-8 : 1~4 바이트 가변
UTF-16 : 2~4 바이트 가변
UTF-32 : 4 바이트 고정

https://miaow-miaow.tistory.com/37

# 문자열 변경
새 값을 지정할 순 있지만, 문자열을 변경할 수는 없음

# 