# make
만들어야 할 slice의 크기를 정확히 알고 있다면 make를 사용하여 slice가 사용하는 array의 크기를 고정할 수 있음

# make는 keyword?
keyword는 아니고 내장함수에 소개되어 있음.

# make cap
할당한 array의 공간이 적은데 append를 사용하여 slice를 추가할 경우, cap를 벗어나는 순간 cap가 2배로 늘어난다.
사용하던 하부 array를 버리고, 2배 더 큰 array를 만들어 사용하게 되며, 이 작업에는 시간이 좀 걸린다고 한다.

# {}를 사용하여 값을 할당하여 slice를 만들 경우.
값을 할당한 길이만큼 cap가 할당된다.
다시 말해서 append를 써서 slice의 크기가 늘어나게 되면, 기존 하부 array를 버리고 2배 더 큰 하부 array를 만들며, 작업시간이 늘어나게 된다.
어느 범위에서 slice를 사용할 지 정해져있다면 make 명령어를 사용하여 하부 array 길이를 정해놓는게 프로그램 속도를 높일 수 있다.