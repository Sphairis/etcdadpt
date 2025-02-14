/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package log

import (
	"fmt"
	"runtime"

	"github.com/coreos/pkg/capnslog"

	"github.com/go-chassis/foundation/fileutil"
	"github.com/go-chassis/openlog"
)

const grpcCallerSkip = 2

// Logger implement from grcplog.LoggerV2s and capnslog.Formatter
type Logger struct {
	Logger openlog.Logger
}

func (l *Logger) Format(pkg string, level capnslog.LogLevel, depth int, entries ...interface{}) {
	format := l.getCaller(depth+1+grpcCallerSkip) + " " + pkg + " %s"
	switch level {
	case capnslog.NOTICE, capnslog.DEBUG, capnslog.TRACE:
		l.Logger.Debug(fmt.Sprintf(format, entries...))
	default:
		l.Logger.Error(fmt.Sprintf(format, entries...), nil)
	}
}

func (l *Logger) getCaller(depth int) string {
	_, file, line, ok := runtime.Caller(depth + 1)
	if !ok {
		return "???"
	}
	return fmt.Sprintf("%s:%d", fileutil.LastNameOf(file), line)
}

func (l *Logger) Flush() {
}

func (l *Logger) Debug(args ...interface{}) {
	l.Logger.Debug(fmt.Sprint(args...))
}

func (l *Logger) Debugln(args ...interface{}) {
	l.Logger.Debug(fmt.Sprint(args...))
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Logger.Debug(fmt.Sprintf(format, args...))
}

func (l *Logger) Info(args ...interface{}) {
	l.Logger.Info(fmt.Sprint(args...))
}

func (l *Logger) Infoln(args ...interface{}) {
	l.Logger.Info(fmt.Sprint(args...))
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.Logger.Info(fmt.Sprintf(format, args...))
}

func (l *Logger) Warning(args ...interface{}) {
	l.Logger.Warn(fmt.Sprint(args...))
}

func (l *Logger) Warningln(args ...interface{}) {
	l.Logger.Warn(fmt.Sprint(args...))
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	l.Logger.Warn(fmt.Sprintf(format, args...))
}

func (l *Logger) Error(args ...interface{}) {
	l.Logger.Error(fmt.Sprint(args...), nil)
}

func (l *Logger) Errorln(args ...interface{}) {
	l.Logger.Error(fmt.Sprint(args...), nil)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Logger.Error(fmt.Sprintf(format, args...), nil)
}

// V reports whether verbosity level l is at least the requested verbose level.
func (l *Logger) V(_ int) bool {
	return true
}

func (l *Logger) Fatal(args ...interface{}) {
	l.Logger.Fatal(fmt.Sprint(args...), nil)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatal(fmt.Sprintf(format, args...), nil)
}

func (l *Logger) Fatalln(args ...interface{}) {
	l.Logger.Fatal(fmt.Sprint(args...), nil)
}

func (l *Logger) Print(args ...interface{}) {
	l.Logger.Error(fmt.Sprint(args...), nil)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.Logger.Error(fmt.Sprintf(format, args...), nil)
}

func (l *Logger) Println(args ...interface{}) {
	l.Logger.Error(fmt.Sprint(args...), nil)
}
