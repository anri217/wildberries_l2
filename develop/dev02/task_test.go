package main

import "testing"

func TestUnpackString(t *testing.T) {
	answerStr := "aaaabccddddde"
	str, err := UnpackStr("a4bc2d5e")
	if err != nil {
		t.Fatalf("Can't unpack your string: %v", err)
	} else if str != answerStr {
		t.Fatalf("The function is not working: %v", err)
	}
}

func TestUnpackEscapeString(t *testing.T) {
	answerStr := "qwe45"
	str, err := UnpackStr(`qwe\4\5`)
	if err != nil {
		t.Fatalf("Can't unpack your string: %v", err)
	} else if str != answerStr {
		t.Fatalf("The function is not working: %v", err)
	}
}
