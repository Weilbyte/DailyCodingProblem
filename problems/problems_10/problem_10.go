package problems10

/* Problem Info
Asked by: Apple

Implement a job scheduler which takes in a function f and an integer n, and calls f after n milliseconds.

*/

import "time"

var runAfterWork = false

func runAfterExample() {
	runAfterWork = true
}

func runAfter(function interface{}, timeMs int) bool {
	callfunc := function.(func())
	time.Sleep(time.Duration(timeMs) * time.Millisecond)
	callfunc()
	return runAfterWork
}
