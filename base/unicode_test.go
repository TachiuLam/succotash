package base

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestUTF8Length(t *testing.T) {
	{
		s := "hello"
		t.Logf("len(%s) = %d", s, len(s))
	}
	{
		s := "你好"
		t.Logf("len(%s) = %d", s, len(s))

	}
}

func TestUnicodeLength(t *testing.T) {
	{
		s := "hello"
		t.Logf("unicodeLen(%s) = %d", s, utf8.RuneCountInString(s))
	}
	{
		s := "你好"
		t.Logf("unicodeLen(%s) = %d", s, utf8.RuneCountInString(s))
	}
	{
		s := "😄"
		t.Logf("unicodeLen(%s) = %d", s, utf8.RuneCountInString(s))
	}
}

func TestUTF8(t *testing.T) {
	{
		s := "hello z"
		t.Logf("print %s", s)
		RangePrint(s)
	}
	fmt.Println()
	{
		s := "你好"
		t.Logf("print %s", s)
		RangePrint(s)
	}
	fmt.Println()
	{
		s := "😄"
		t.Logf("print %s", s)
		RangePrint(s)
	}
}

func TestUtf8Bytes(t *testing.T) {
	{
		s := "😄"
		for _, b := range []byte(s) {
			fmt.Printf("%#02x ", b)
		}
	}
}

func RangePrint(s string) {
	for _, char := range s {
		fmt.Printf("%U ", char)
	}
}
