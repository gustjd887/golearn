# testing
test 파일을 만들 때, file_test.go, _test.go를 뒤에 붙여야 함.
func TestFunc_name(t *testing.T) 라는 함수가 들어감. Func_name은 첫 자를 대문자로 써야 함.
함수 이름과 일치시킬 필요는 없지만, 일치 시키는게 좋음.
test 파일은 동일한 패키지에 위치시키는게 좋음.