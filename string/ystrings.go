package ystrings

import (
	"unsafe"
)

// Str2bytes string零拷贝转[]byte，注意修改[]byte时，原字符串也可能会随之变动
func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// Bytes2str []byte零拷贝转string，注意修改[]byte时，字符串也可能会随之变动
func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
