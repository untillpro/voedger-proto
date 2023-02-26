/*
 * Copyright (c) 2020-present unTill Pro, Ltd.
 * @author Maxim Geraskin
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package logger

import (
	"fmt"
	"runtime"
	"strings"
	"sync/atomic"
	"time"
)

const (
	skipStackFramesCount = 4
	normalLineLength     = 60
)

const (
	errorPrefix   = "*****"
	warningPrefix = "!!!"
	infoPrefix    = "==="
	verbosePrefix = "---"
	tracePrefix   = "..."
)

var globalLogPrinter = logPrinter{logLevel: LogLevelInfo}

type logPrinter struct {
	logLevel TLogLevel
}

func isEnabled(logLevel TLogLevel) bool {
	curLogLevel := TLogLevel(atomic.LoadInt32((*int32)(&globalLogPrinter.logLevel)))
	return curLogLevel >= logLevel
}

func (p *logPrinter) getFuncName(skipCount int) (funcName string, line int) {
	var fn string
	pc, _, line, ok := runtime.Caller(skipCount)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		elems := strings.Split(details.Name(), "/")
		if len(elems) > 0 {
			fn = elems[len(elems)-1]
		}
	}
	return fn, line
}

func (p *logPrinter) getFormattedMsg(msgType string, funcName string, line int, args ...interface{}) string {
	t := time.Now()
	out := fmt.Sprint(t.Format("01/02 15:04:05.000"))
	out += fmt.Sprint(": " + msgType)
	out += fmt.Sprintf(": [%v:%v]:", funcName, line)
	if len(args) > 0 {
		var s string
		for _, arg := range args {
			s = s + fmt.Sprint(" ", arg)
		}
		for i := len(s); i < normalLineLength; i++ {
			s = s + " "
		}
		out += fmt.Sprint(s)
	}
	return out
}

func (p *logPrinter) print(msgType string, args ...interface{}) {
	funcName, line := p.getFuncName(skipStackFramesCount)
	out := p.getFormattedMsg(msgType, funcName, line, args...)
	fmt.Println(out)
}

func getLevelPrefix(level TLogLevel) string {
	switch level {
	case LogLevelError:
		return errorPrefix
	case LogLevelWarning:
		return warningPrefix
	case LogLevelInfo:
		return infoPrefix
	case LogLevelVerbose:
		return verbosePrefix
	case LogLevelTrace:
		return tracePrefix
	}
	return ""
}

func printIfLevel(level TLogLevel, args ...interface{}) {
	if isEnabled(level) {
		globalLogPrinter.print(getLevelPrefix(level), args...)
	}
}
