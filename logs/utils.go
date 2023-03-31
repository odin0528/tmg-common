package logs

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func getFuncCallerName(skipFrames int) string {
	// skip getFuncCallerName
	funNameList := getAllFuncCallerNameList(skipFrames + 1)

	if 0 >= len(funNameList) {
		return "unknown"
	}

	return funNameList[0]
}

// getAllFuncCallerNameList - get 20 of caller func names stack except for runtime.Callers and getAllFuncCallerNameList
func getAllFuncCallerNameList(skipFrames int) []string {
	defer handlePanic()

	funNameList := []string{}

	// Set size to skipFrames+20 to ensure we have room for 20 more callers than we need
	programCounters := make([]uintptr, skipFrames+20)
	// Skip 2, runtime.Callers and getAllFuncCallerNameList
	n := runtime.Callers(2, programCounters)

	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more; frameIndex++ {
			var frame runtime.Frame
			frame, more = frames.Next()
			if frameIndex < skipFrames {
				continue
			}
			funNameList = append(funNameList, frame.Function)
		}
	}

	return funNameList
}

func handlePanic() {
	if err := recover(); err != nil {
		Error(PANIC, FUNC_CALLER, fmt.Sprintln(err)+string(debug.Stack()), map[string]interface{}{})
	}
}
