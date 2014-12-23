package clip

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
*/
import "C"

func Copy(p string) {
	C.Copy(C.CString(p))
}
