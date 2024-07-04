// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import (
	"fmt"
	"strings"
)

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	// also use strings.Replace function
	name = strings.TrimSpace(name)
	if name == "" {
		name = "you"
		// 인자로 받은 변수의 값을 수정하는 것은 일반적으로 허용되는 행위.
		// -> 특히, 문자열과 같은 immutable type의 경우에 원본 데이터에 영향을 주지 않고 새로운 값 생성.
		// -> 사이드 이펙트 걱정 X
		// 가동성과 명확성, 불별성 유지, 효율성 등과 같은 부분을 고려하여 함수 내에서 인자 변수 변경.
	}

	return "One for " + name + ", one for me."
	// 기본적인 방식 -> go에서 문자열을 immutable 그래서 문자열을 조작할 때마다 새로운 문자열을 생성.
}

func ShareWithOtherWay(name string) string {
	name = strings.Replace(name, " ", "", -1)
	if name == "" {
		defalutValue := "you"
		return fmt.Sprintf("One for %s, one for me.", defalutValue)
	}

	var builder strings.Builder
	builder.WriteString("One for ")
	builder.WriteString(name)
	priorStr := builder.String()
	return strings.Join([]string{priorStr, "one for me."}, ", ")
}
