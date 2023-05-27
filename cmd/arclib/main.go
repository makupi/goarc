package main

import (
	"C"
	"github.com/makupi/goarc/pkg/decoder"
	"github.com/makupi/goarc/pkg/encoder"
)

//export decode
func decode(url *C.char, addr *C.char) *C.char {
	r, err := decoder.ParseASAUrl(C.GoString(url), C.GoString(addr))
	if err != nil {
		panic(err)
	}
	return C.CString(r)
}

//export encode
func encode(toEncode *C.char) *C.char {
	r, err := encoder.ReserveAddressFromCID(C.GoString(toEncode))
	if err != nil {
		panic(err)
	}
	return C.CString(r)
}

func main() {}
