package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
	return fmt.Sprintf("%s\n%s\n%s", border, welcomeMsg, border)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	cleanupedMsg := strings.Split(oldMsg, "\n")[1]
	return strings.Trim(cleanupedMsg, "* ")
}

/*
don't necessarily hve to represent characters가 의미하는 것.
- 문자열이 반드시 문자(char)만을 나타내는 것이 아님.
	-> binaryData := "\x00\x01\x02\x03\x04" 요런 상황처럼 바이너리 데이터를 포함 가능.
- 위와 같이 바이너리 데이터를 다룰 때는 문자열을 사용하여 처리할순 있지만 일반적으로 []byte 이용.
	-> binaryData := []byte{0, 1, 2, 3, 4}
	-> 이렇게 표현 시 명확하게 표현할 수 있고 쉽게 바이트 접근 및 수정이 가능.

- 인코딩, 디코딩 방법
	```
		- 인코딩
		text := "Hello, World!"
		data := []byte(text)
		---
		- 디코딩
		data := []byte{72, 101, 108}
		text := string(data)
		,
		encoding/binary 모듈 이용 -> 원하는 타입으로 디코딩 가능.

		data := []byte{0x00, 0x00, 0x00, 0x01}
		var num int32
		buf := bytes.NewBuffer(data)  -> 바이트 슬라이스를 읽을 수 있는 Reader 생성.
		err := binary.Read(buf, binary.BigEndian, &num)  -> 바이너리 데이터를 읽어서 num에 저장.
		if err != nil {
			fmt.Println("binary.Read failed: ", err)
		}
		fmt.Println(num) => 출력: 1
	```
*/
