package main

/*
#cgo LDFLAGS: -L . -lc_test -lstdc++
#cgo CFLAGS: -I ./ -IE:/c/mingw64/lib/gcc/x86_64-w64-mingw32/8.1.0/include/c++
#cgo CFLAGS: -I ./ -IE:/c/mingw64/lib/gcc/x86_64-w64-mingw32/8.1.0/include/c++/bits/stl_algobase.h
#include "lunar.cpp"
*/
import "C"

import "fmt"

func GoSum() {
	s := C.getYearCal(2019)
	fmt.Println(s)
}

func main()  {
	GoSum()
}