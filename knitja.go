package main

//int samu(int argc, char *argv[]);
import "C"
import (
	"os"
	"unsafe"
)

func main() {
	arr := []*C.char{C.CString(os.Args[0]), C.CString("-t"), C.CString("toknit"), nil}
	C.samu(C.int(len(arr)-1), (**C.char)(unsafe.Pointer(&arr[0])))
}
