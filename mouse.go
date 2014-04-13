// +build darwin

package userinput
// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include <AppKit/NSEvent.h>
//
// NSPoint mouseLocation(){
//  return [NSEvent mouseLocation];
// }
import "C"

//NSPoint
type Point struct {
  X float64
  Y float64
}

func MouseClick( x float64, y float64 ){
  
}

func MouseMove( x float64, y float64 ){
  var move = C.CGEventCreateMouseEvent(
                nil, 
                C.kCGEventMouseMoved, 
                C.CGPointMake(C.CGFloat(x),C.CGFloat(y)),
                0)
  // defer C.CFRelease( move )

  C.CGEventPost(C.kCGHIDEventTap, move)
}

func MousePosition() Point {
  var event = C.mouseLocation()
  return Point{ X: float64(event.x), Y: float64(event.y) }

}