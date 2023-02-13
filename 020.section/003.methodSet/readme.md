# method set
변수의 타입이 Value인 경우, 인터페이스의 Method Sets을 충족하려면 인터페이스가 정의하는 모든 메서드의 receiver 타입은 오직 Value여야 한다.
변수의 타입이 Pointer인 경우, receiver 타입은 Value든 Pointer든 상관 없다.

https://velog.io/@undefcat/Go-Method-Sets

VALUE RECEIVER로 구현된 메서드는, 무조건 VALUE를 리시버에 넘겨야 한다.

# 무조건 값 리시버 메서드를 구현하는게 더 좋은가...?
value든 pointer든 둘다 사용 가능해지니...

# 기본적으로 go는 value를 던짐
하지만 포인터 리시버로 구현된 메서드에는 주소를 던지나...