# zero value
변수 타입만 지정했을 경우 출력되는 값

# var를 사용한 identifier 선언
var를 사용하여 identifier 선언 시 꼭 함수 밖에서 사용하는건 아님..

단지 identifire 선언 위치가 어디냐에 따라 package scope가 되느냐, func scope가 되냐 차이가 남

마찬가지로 단축연산자(:=)를 사용한 선언 방법도 똑같을 것 같음

# var를 사용한 선언과 단축연산자(:=)를 사용한 선언의 차이점
* var를 사용하면 zero value와 type을 미리 지정할 수 있음 
* :=를 사용하면 zero value를 사용할 수 없고, value가 지정되어야 type을 사용할 수 있다는 의미가 됨.

# 단축연산자를 사용한 변수를 선언하지만 value를 넣지 않았을 경우..
아래와 같은 에러 발생함.
```
syntax error: non-declaration statement outside function body
```
# 결론
go에서는 identifier 선언 후 value를 넣지 않은 값들에 대해 compile시 검사하게 되고
value가 없으면 에러가 남
하지만 var를 사용하여 선언하면 zero value를 넣어줌으로서 에러가 나지 않음
그런 용도의 차이인 것 같음
