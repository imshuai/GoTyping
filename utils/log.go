package utils

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
"github.com/gorilla/websocket"
	"github.com/gin-gonic/gin"
)

var (
	lgFatal *logger
	//LogFatal 崩溃日志记录器
	LogFatal = lgFatal.Log
	lgError  *logger
	//LogError 错误日志记录器
	LogError = lgError.Log
	lgWarn   *logger
	//LogWarn 警告日志记录器
	LogWarn = lgWarn.Log
	lgInfo  *logger
	//LogInfo 详情日志记录器
	LogInfo = lgInfo.Log
	lgDebug *logger
	//LogDebug 测试日志记录器
	LogDebug = lgDebug.Log
)

var (
	isDaylyRolling bool
	isFileRolling  bool
	isConsole      bool
	logfile        *os.File
	logchan        chan []byte
)

//EnginLogger Gin框架的日志中间件
func EnginLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()
		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		clientIP := func(c *gin.Context) string {
			if ips := c.GetHeader("x-forwarded-for"); ips != "" {
				ip := strings.Split(ips, ",")
				return ip[0]
			}
			if ip := c.GetHeader("Proxy-Client-IP"); ip != "" {
				return ip
			}
			if ip := c.GetHeader("X-Real-IP"); ip != "" {
				return ip
			}
			return c.ClientIP()
		}(c)
		method := c.Request.Method
		statusCode := c.Writer.Status()

		if raw != "" {
			path = path + "?" + raw
		}

		fmt.Fprintf(lgInfo, "[GIN] %v [%3d] %13v | %15s |%-7s %s\n",
			end.Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)

	}
}

type logLevel int

const (
	ALL logLevel = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	OFF
)

type logger struct {
	level       logLevel
	levelString string
	w           chan []byte
}

func (l *logger) Write(p []byte) (n int, err error) {
	prefix := "[" + l.levelString + "] - "
	buf := bytes.NewBufferString(prefix)
	buf.Write(p)
	if isConsole {
		switch l.level {
		case ALL, DEBUG, INFO, WARN:
			buf.WriteTo(os.Stdout)
		case ERROR, FATAL:
			buf.WriteTo(os.Stderr)
		default:
		}
	}
	lenth := buf.Len()
	l.w <- buf.Bytes()
	return lenth, nil
}

func (l *logger) Log(s ...interface{}) {
	fmt.Fprintln(l, s...)
}

func logWriter(c <-chan []byte, w io.Writer) {
	for p := range c {
		if logfile != nil {
			_, err := w.Write(p)
			NoticManager("日志写入错误", err.Error())
		}
	}
}

//SetLogFile 设置日志输出文件
func SetLogFile(fpath string) error {
	f, err := os.OpenFile(fpath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	logfile = f
	return nil
}

//SetConsole 设置日志是否输出至控制台
func SetConsole(sw bool) {
	isConsole = sw
}

func loggerInit() {
	logchan = make(chan []byte, 100)
	//创建不同级别日志对象
	for i := 1; i <= 5; i++ {
		switch level := logLevel(i); level {
		case DEBUG:
			lgDebug = &logger{
				level:       level,
				levelString: "Debug",
				w:           logchan,
			}
		case INFO:
			lgInfo = &logger{
				level:       level,
				levelString: "Info",
				w:           logchan,
			}
		case WARN:
			lgWarn = &logger{
				level:       level,
				levelString: "Warn",
				w:           logchan,
			}
		case ERROR:
			lgError = &logger{
				level:       level,
				levelString: "Error",
				w:           logchan,
			}
		case FATAL:
			lgFatal = &logger{
				level:       level,
				levelString: "Fatal",
				w:           logchan,
			}
		}
	}
	go logWriter(logchan, logfile)
}
