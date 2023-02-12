# json encode

* NewEncoder : w Writer Type을 파라미터로 가지고, *encoder 를 반환한다.
* os.Stdout : File Type에 Write 메서드가 구현되어 있으며, *File 을 반환한다.(화면에 출력하는 포인터를 반환)
* Encode : *Encoder를 리시버로 받는 메서드, err를 반환

# NewEncoder(os.Stdout).Encode(users)
* 리시버는 *Encoder가 필요하고, 파라미터는 아무거나 와도 된다.
* 에러만 반환하는데, 만약 저장이 필요하면 Encoder를 미리 지정해놔야 할듯....