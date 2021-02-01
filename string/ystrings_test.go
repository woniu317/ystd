package ystrings

import (
	"reflect"
	"strings"
	"testing"
)

// TestStr2Byte 字符串转[]byte
func TestStr2Byte(t *testing.T) {
	for i, tc := range []struct {
		str string
		ret []byte
	}{
		{"abc中国", []byte{97, 98, 99, 228, 184, 173, 229, 155, 189}},
		{"abc", []byte{97, 98, 99}},
		{"中国", []byte{228, 184, 173, 229, 155, 189}},
	} {
		b := Str2bytes(tc.str)
		if !reflect.DeepEqual(b, tc.ret) {
			t.Errorf("No.%d Str2bytes error", i)
		}
	}
}

// TestStr2Byte2 修改转换后的[]byte, 原字符串会同步更改
func TestStr2Byte2(t *testing.T) {
	s := strings.Repeat("abc", 1)
	b := Str2bytes(s)
	b[1] = 'e'
	if !reflect.DeepEqual(s, "aec") {
		t.Error("change byte error")
	}
}

// TestBytes2Str []byte转string
func TestBytes2Str(t *testing.T) {
	b := []byte{97, 98, 99, 228, 184, 173, 229, 155, 189}
	s := Bytes2str(b)
	if !reflect.DeepEqual(s, "abc中国") {
		t.Error("Bytes2Str error")
	}
}
