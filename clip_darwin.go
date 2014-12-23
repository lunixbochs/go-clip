package clip

import (
	"unsafe"
)

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework AppKit

#import <Foundation/Foundation.h>
#import <Cocoa/Cocoa.h>

void Copy(char *p) {
	NSString *s = [NSString stringWithUTF8String:p];
	[[NSPasteboard generalPasteboard] clearContents];
	[[NSPasteboard generalPasteboard] setString:s forType:NSStringPboardType];
}

char *Paste() {
	NSString *s = [[NSPasteboard generalPasteboard] stringForType:NSStringPboardType];
	[s autorelease];
	return strdup([s UTF8String]);
}
*/
import "C"

func Copy(p string) {
	C.Copy(C.CString(p))
}

func Paste() string {
	s := C.Paste()
	defer C.free(unsafe.Pointer(s))
	return C.GoString(s)
}
