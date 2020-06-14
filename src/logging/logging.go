package logging

import (
	"fmt"
	// TODO Copy default format and just remove import
	"log"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/kylelemons/godebug/pretty"
)

var (
	errLog = log.New(os.Stderr, "", 0)
	outLog = log.New(os.Stdout, "", 0)

	LevelSending          Level = 2
	terminalSupportsColor       = false

	DebugEnabled = false
)

const (
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DEBUG = 0
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	INFO = 1
	// WarnLevel level. Non-critical entries that deserve eyes.
	WARN = 2
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ERROR = 3
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FATAL = 4
)

type Level uint32

func init() {
	var outString string = "0"

	// Check whether the current terminal supports colours
	out, _ := exec.Command("/usr/bin/tput", "colors").Output()

	outString = strings.TrimSpace(string(out))
	colours, err := strconv.Atoi(outString)
	if err != nil {
		sendLog(err, ERROR)
	}

	if colours > 8 {
		terminalSupportsColor = true
	}

	// Check the current Debug status
	if os.Getenv("DEBUG") == "true" {
		DebugEnabled = true
	}
}

func sendLog(i interface{}, l Level) {
	// Get the runtime caller
	pc, _, line, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	if !ok && details == nil {
		return
	}

	// Check the current debug status
	if l == DEBUG && !DebugEnabled {
		return
	}

	// Manage the current time (if enabled)
	timestring := time.Now().Format("Mon Jan _2 15:04:05") + " "

	// If the type is a structure format it nicely for output
	if t := reflect.TypeOf(i); t.Kind() == reflect.Struct {
		i = "\n" + pretty.Sprint(i)
	}

	// Include tracing information if debug mode is enabled
	if DebugEnabled {
		if l <= LevelSending {
			outLog.Printf("%s[%s] [%s#%d] %s \n", timestring, levelToString(l), details.Name(), line, i)
			return
		} else {
			errLog.Printf("%s[%s] [%s#%d] %s \n", timestring, levelToString(l), details.Name(), line, i)
			return
		}
	}

	// Keep the logs simple if debug is not enabled
	if l <= LevelSending {
		outLog.Printf("%s[%s] %s \n", timestring, levelToString(l), i)
		return
	} else {
		errLog.Printf("%s[%s] %s \n", timestring, levelToString(l), i)
		return
	}
}

func Error(err interface{}) {
	sendLog(err, ERROR)
}
func Debug(debug interface{}) {
	sendLog(debug, DEBUG)
}
func Info(info interface{}) {
	sendLog(info, INFO)
}
func Warn(warn interface{}) {
	sendLog(warn, WARN)
}
func Fatal(fatal interface{}) {
	sendLog(fatal, FATAL)
	os.Exit(1)
}

func Errorf(form string, err ...interface{}) {
	sendLog(fmt.Sprintf(form, err...), ERROR)
}
func Debugf(form string, debug ...interface{}) {
	sendLog(fmt.Sprintf(form, debug...), DEBUG)
}
func Infof(form string, info ...interface{}) {
	sendLog(fmt.Sprintf(form, info...), INFO)
}
func Warnf(form string, warn ...interface{}) {
	sendLog(fmt.Sprintf(form, warn...), WARN)
}
func Fatalf(form string, fatal ...interface{}) {
	sendLog(fmt.Sprintf(form, fatal...), FATAL)
	os.Exit(1)
}

func FatalNoExit(fatal interface{}) {
	sendLog(fatal, FATAL)
}

// RecoveryHandler in case it crashes
func RecoveryHandler(next http.Handler) http.Handler {
	return handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(next)
}

func levelToString(l Level) string {
	switch l {
	case DEBUG:
		if terminalSupportsColor {
			return " \u001b[32mDEBUG\u001b[0m "
		}
		return " DEBUG "
	case INFO:
		if terminalSupportsColor {
			return " \u001b[34mINFO\u001b[0m  "
		}
		return "INFO  "
	case WARN:
		if terminalSupportsColor {
			return " \u001b[33mWARN\u001b[0m  "
		}
		return " WARN  "
	case ERROR:
		if terminalSupportsColor {
			return " \u001b[31mERROR\u001b[0m "
		}
		return " ERROR "
	case FATAL:
		if terminalSupportsColor {
			return " \u001b[31mFATAL\u001b[0m "
		}
		return " FATAL "
	default:
		return " N/A "
	}
}
