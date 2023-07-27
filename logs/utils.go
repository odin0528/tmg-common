package logs

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
)

func getFuncCallerName(skipFrames int) string {
	funNameList := getAllFuncCallerNameList(skipFrames + 1) //Add skip this function (getFuncCallerName) layer

	if 0 >= len(funNameList) {
		return UNKNOWN_CALLER
	}

	return funNameList[0]
}

func getAllFuncCallerNameList(skipFrames int) []string {
	defer handlePanic()

	funNameList := []string{}

	programCounters := make([]uintptr, skipFrames+MAX_CALLER_COUNT)

	n := runtime.Callers(CALLER_SKIP_LAYER, programCounters) // Skip "runtime.Callers" and "getAllFuncCallerNameList"

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
		Error(LOG_TYPE_PANIC_RECOVER, LOG_KEY_FUNC_CALLER, fmt.Sprintln(err)+string(debug.Stack()), map[string]interface{}{})
	}
}

func combinFuncCallerName(callers []string) string {
	return strings.Join(getAllFuncCallerNameList(BASE_SKIP_LAYER+1), " < ") //Add skip this function (combinFuncCallerName) layer
}
