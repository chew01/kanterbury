package log

import (
	"fmt"
	"github.com/chew01/kanterbury/utils"
	"github.com/logrusorgru/aurora"
	"github.com/mattn/go-colorable"
	"io"
	"log"
	"os"
	"path"
	"runtime"
)

type Logger interface {
	Printf(string, ...interface{})
	Println(...interface{})
}

type kLog struct {
	fileLogger   *log.Logger
	stdoutLogger *log.Logger
}

// New initializes and returns a new instance of Logger
func New(stdout bool, filePath string, flags int) Logger {
	logger := &kLog{}

	// Add color compatibility for Windows
	var output io.Writer
	if runtime.GOOS == "windows" {
		output = colorable.NewColorableStdout()
	} else {
		output = os.Stdout
	}

	// If stdout is enabled, initialize new stdout logger
	if stdout {
		logger.stdoutLogger = log.New(output, "", flags)
	}

	// If filepath is "/dev/null", return. Else create logfile and its parent directories, then initialize logger
	if filePath == "/dev/null" {
		return logger
	} else if filePath == "" {
		filePath = path.Join(utils.BinDir, "/logs/proxy.log")
	} else if !path.IsAbs(filePath) {
		filePath = path.Join(utils.BinDir, filePath)
	}

	dir := path.Dir(filePath)
	utils.Must(os.MkdirAll(dir, 0755))
	logFile, err := os.Create(filePath)
	utils.Must(err)
	logger.fileLogger = log.New(logFile, "", flags)

	return logger
}

func (log *kLog) output(calldepth int, prefixColor func(interface{}) aurora.Value, prefix, str string) {
	if log == nil {
		return
	}
	calldepth++
	if log.fileLogger != nil {
		utils.Must(log.fileLogger.Output(calldepth, prefix+str))
	}
	if log.stdoutLogger != nil {
		if prefixColor != nil {
			prefix = prefixColor(prefix).String()
		}
		// Don't print long strings in stdout, truncate them to 400 chars.
		if len(str) > 403 {
			str = str[0:400] + "..."
		}
		utils.Must(log.stdoutLogger.Output(calldepth, prefix+str))
	}
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (log *kLog) Printf(format string, v ...interface{}) {
	log.output(2, aurora.Green, "INFO ", fmt.Sprintf(format, v...))
}

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func (log *kLog) Println(v ...interface{}) {
	log.output(2, aurora.Green, "INFO ", fmt.Sprintln(v...))
}
