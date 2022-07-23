package helper

import (
	"fmt"
	"runtime"
	"strings"
)

//const template = "Error: %s:%s:L%d: %v"
const template = "Error: %s:L%d: \"%v\""

func PrintTrace(err error) string {
	errMsg := strings.ReplaceAll(err.Error(), "\n", "")
	//if pc, file, line, ok := runtime.Caller(1); !ok {
	if _, file, line, ok := runtime.Caller(1); !ok {
		//return fmt.Sprintf(template, "?", "?", 0, err)
		return fmt.Sprintf(template, "?", 0, errMsg)
	} else {
		//return fmt.Sprintf(template, file, runtime.FuncForPC(pc).Name(), line, err)
		return fmt.Sprintf(template, file, line, errMsg)
	}
}
