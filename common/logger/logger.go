package logger

import (
	_f "fmt"
	_gce "io"
	_fd "os"
	_gc "path/filepath"
	_g "runtime"
)

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_dg ConsoleLogger) IsLogLevel(level LogLevel) bool { return _dg.LogLevel >= level }

// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{ LogLevel LogLevel }

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_ga WriterLogger) IsLogLevel(level LogLevel) bool { return _ga.LogLevel >= level }

// Notice logs notice message.
func (_gb ConsoleLogger) Notice(format string, args ...interface{}) {
	if _gb.LogLevel >= LogLevelNotice {
		_ba := "\u005bN\u004f\u0054\u0049\u0043\u0045\u005d "
		_gb.output(_fd.Stdout, _ba, format, args...)
	}
}

// Notice does nothing for dummy logger.
func (DummyLogger) Notice(format string, args ...interface{}) {}

// Error does nothing for dummy logger.
func (DummyLogger) Error(format string, args ...interface{}) {}

// Debug logs debug message.
func (_fbf WriterLogger) Debug(format string, args ...interface{}) {
	if _fbf.LogLevel >= LogLevelDebug {
		_bgg := "\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020"
		_fbf.logToWriter(_fbf.Output, _bgg, format, args...)
	}
}

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct {
	LogLevel LogLevel
	Output   _gce.Writer
}

// IsLogLevel returns true from dummy logger.
func (DummyLogger) IsLogLevel(level LogLevel) bool { return true }

// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger(logLevel LogLevel, writer _gce.Writer) *WriterLogger {
	logger := WriterLogger{Output: writer, LogLevel: logLevel}
	return &logger
}

// Warning logs warning message.
func (_acg WriterLogger) Warning(format string, args ...interface{}) {
	if _acg.LogLevel >= LogLevelWarning {
		_gca := "\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020"
		_acg.logToWriter(_acg.Output, _gca, format, args...)
	}
}

// Info does nothing for dummy logger.
func (DummyLogger) Info(format string, args ...interface{}) {}

// Trace does nothing for dummy logger.
func (DummyLogger) Trace(format string, args ...interface{}) {}

func _aac(_bfc _gce.Writer, _gfb string, _cca string, _fcd ...interface{}) {
	_, _fdd, _acc, _faa := _g.Caller(3)
	if !_faa {
		_fdd = "\u003f\u003f\u003f"
		_acc = 0
	} else {
		_fdd = _gc.Base(_fdd)
	}
	_eb := _f.Sprintf("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ", _gfb, _fdd, _acc) + _cca + "\u000a"
	_f.Fprintf(_bfc, _eb, _fcd...)
}

// Trace logs trace message.
func (_dd ConsoleLogger) Trace(format string, args ...interface{}) {
	if _dd.LogLevel >= LogLevelTrace {
		_df := "\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020"
		_dd.output(_fd.Stdout, _df, format, args...)
	}
}

// Trace logs trace message.
func (_aab WriterLogger) Trace(format string, args ...interface{}) {
	if _aab.LogLevel >= LogLevelTrace {
		_c := "\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020"
		_aab.logToWriter(_aab.Output, _c, format, args...)
	}
}

// DummyLogger does nothing.
type DummyLogger struct{}

// Info logs info message.
func (_fgd WriterLogger) Info(format string, args ...interface{}) {
	if _fgd.LogLevel >= LogLevelInfo {
		_efg := "\u005bI\u004e\u0046\u004f\u005d\u0020"
		_fgd.logToWriter(_fgd.Output, _efg, format, args...)
	}
}

// Logger is the interface used for logging in the unipdf package.
type Logger interface {
	Error(_ec string, _b ...interface{})
	Warning(_ff string, _bb ...interface{})
	Notice(_a string, _ac ...interface{})
	Info(_d string, _bd ...interface{})
	Debug(_bg string, _acf ...interface{})
	Trace(_ecf string, _fb ...interface{})
	IsLogLevel(_acfe LogLevel) bool
}

// SetLogger sets 'logger' to be used by the unidoc unipdf library.
func SetLogger(logger Logger) { Log = logger }

func (_cc WriterLogger) logToWriter(_ed _gce.Writer, _bab string, _ggd string, _cd ...interface{}) {
	_aac(_ed, _bab, _ggd, _cd)
}

// Info logs info message.
func (_gf ConsoleLogger) Info(format string, args ...interface{}) {
	if _gf.LogLevel >= LogLevelInfo {
		_fg := "\u005bI\u004e\u0046\u004f\u005d\u0020"
		_gf.output(_fd.Stdout, _fg, format, args...)
	}
}

// Debug does nothing for dummy logger.
func (DummyLogger) Debug(format string, args ...interface{}) {}

const (
	LogLevelTrace   LogLevel = 5
	LogLevelDebug   LogLevel = 4
	LogLevelInfo    LogLevel = 3
	LogLevelNotice  LogLevel = 2
	LogLevelWarning LogLevel = 1
	LogLevelError   LogLevel = 0
)

var Log Logger = DummyLogger{}

// Notice logs notice message.
func (_gad WriterLogger) Notice(format string, args ...interface{}) {
	if _gad.LogLevel >= LogLevelNotice {
		_ef := "\u005bN\u004f\u0054\u0049\u0043\u0045\u005d "
		_gad.logToWriter(_gad.Output, _ef, format, args...)
	}
}

// LogLevel is the verbosity level for logging.
type LogLevel int

// Warning logs warning message.
func (_bf ConsoleLogger) Warning(format string, args ...interface{}) {
	if _bf.LogLevel >= LogLevelWarning {
		_fe := "\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020"
		_bf.output(_fd.Stdout, _fe, format, args...)
	}
}

// Warning does nothing for dummy logger.
func (DummyLogger) Warning(format string, args ...interface{}) {}

// Error logs error message.
func (_ffa WriterLogger) Error(format string, args ...interface{}) {
	if _ffa.LogLevel >= LogLevelError {
		_gg := "\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020"
		_ffa.logToWriter(_ffa.Output, _gg, format, args...)
	}
}

func (_eg ConsoleLogger) output(_fa _gce.Writer, _af string, _geb string, _egg ...interface{}) {
	_aac(_fa, _af, _geb, _egg...)
}

// Error logs error message.
func (_fc ConsoleLogger) Error(format string, args ...interface{}) {
	if _fc.LogLevel >= LogLevelError {
		_ge := "\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020"
		_fc.output(_fd.Stdout, _ge, format, args...)
	}
}

// NewConsoleLogger creates new console logger.
func NewConsoleLogger(logLevel LogLevel) *ConsoleLogger { return &ConsoleLogger{LogLevel: logLevel} }

// Debug logs debug message.
func (_aa ConsoleLogger) Debug(format string, args ...interface{}) {
	if _aa.LogLevel >= LogLevelDebug {
		_ee := "\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020"
		_aa.output(_fd.Stdout, _ee, format, args...)
	}
}
