package main

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import <Foundation/Foundation.h>
//#include "application.h"
//#include "window.h"
import "C"
import "unsafe"

func main() {
	C.InitApplication()
	C.Window_MakeKeyAndOrderFront(unsafe.Pointer(C.Window_Create(1200, 750, 300, 200, C.CString("21 School"))))
	C.RunApplication()
}
