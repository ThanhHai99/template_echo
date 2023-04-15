package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/color"
	"github.com/rs/zerolog"
	"github.com/valyala/fasttemplate"
	"gopkg.in/natefinch/lumberjack.v2"
)

var xlog *zerolog.Logger

func Log() *zerolog.Logger {
	if xlog == nil {
		folderPath := "log/"
		filePath := folderPath + "output.log"
		makeSureDirExists(folderPath)

		ljLogger := &lumberjack.Logger{
			Filename:   filePath,
			MaxBackups: 28, // files
			MaxSize:    8,  // megabytes
			MaxAge:     1,  // days
		}
		multiWriter := io.MultiWriter(
			zerolog.ConsoleWriter{
				Out:        os.Stderr,
				NoColor:    false,
				TimeFormat: time.RFC3339,
			},
			ljLogger,
		)
		logger := zerolog.New(multiWriter).With().Timestamp().Logger()
		xlog = &logger
		xlog.Info().
			Str("logDirectory", folderPath).
			Str("fileName", filePath).
			Int("maxSizeMB", ljLogger.MaxSize).
			Int("maxBackups", ljLogger.MaxBackups).
			Int("maxAgeInDays", ljLogger.MaxAge).
			Msg("logging configured")
	}
	return xlog
}

func makeSureDirExists(location string) {
	loc := location
	if loc[len(loc)-1:] == "/" {
		loc = loc[:len(loc)-1]
	}
	if !isDirectoryExisted(loc) {
		err := os.MkdirAll(loc, os.ModePerm)
		if err != nil {
			Log().Fatal().Err(err)
		}
	}
}

func isDirectoryExisted(location string) bool {
	_, err := os.Stat(location)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

type (
	// LoggerConfig defines the config for Logger middleware.
	LoggerConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper middleware.Skipper

		// Tags to construct the logger format.
		//
		// - time_unix
		// - time_unix_nano
		// - time_rfc3339
		// - time_rfc3339_nano
		// - time_custom
		// - id (Request ID)
		// - remote_ip
		// - uri
		// - host
		// - method
		// - path
		// - protocol
		// - referer
		// - user_agent
		// - status
		// - error
		// - latency (In nanoseconds)
		// - latency_human (Human readable)
		// - bytes_in (Bytes received)
		// - bytes_out (Bytes sent)
		// - header:<NAME>
		// - query:<NAME>
		// - form:<NAME>
		//
		// Example "${remote_ip} ${status}"
		//
		// Optional. Default value DefaultLoggerConfig.Format.
		Format string `yaml:"format"`

		// Optional. Default value DefaultLoggerConfig.CustomTimeFormat.
		CustomTimeFormat string `yaml:"custom_time_format"`

		// Output is a writer where logs in JSON format are written.
		// Optional. Default value os.Stdout.
		Output io.Writer

		template *fasttemplate.Template
		colorer  *color.Color
		pool     *sync.Pool
	}
)

var (
	// DefaultLoggerConfig is the default Logger middleware config.
	DefaultLoggerConfig = LoggerConfig{
		Skipper: middleware.DefaultSkipper,
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
		colorer:          color.New(),
	}
)

// LoggerWithConfig returns a Logger middleware with config.
// See: `Logger()`.
func LoggerWithConfig(config LoggerConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = DefaultLoggerConfig.Skipper
	}
	if config.Format == "" {
		config.Format = DefaultLoggerConfig.Format
	}

	config.template = fasttemplate.New(config.Format, "${", "}")
	config.colorer = color.New()
	config.colorer.SetOutput(config.Output)
	config.pool = &sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 256))
		},
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()
			buf := config.pool.Get().(*bytes.Buffer)
			buf.Reset()
			defer config.pool.Put(buf)

			dumplogger := Log().With()
			msg := "Request"

			if _, err = config.template.ExecuteFunc(buf, func(w io.Writer, tag string) (int, error) {
				switch tag {
				case "time_unix":
					dumplogger = dumplogger.Str("time", strconv.FormatInt(time.Now().Unix(), 10))
					return buf.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
				case "time_unix_nano":
					dumplogger = dumplogger.Str("time", strconv.FormatInt(time.Now().UnixNano(), 10))
					return buf.WriteString(strconv.FormatInt(time.Now().UnixNano(), 10))
				case "time_rfc3339":
					dumplogger = dumplogger.Str("time", time.Now().Format(time.RFC3339))
					return buf.WriteString(time.Now().Format(time.RFC3339))
				case "time_rfc3339_nano":
					dumplogger = dumplogger.Str("time", time.Now().Format(time.RFC3339Nano))
					return buf.WriteString(time.Now().Format(time.RFC3339Nano))
				case "time_custom":
					dumplogger = dumplogger.Str("time", time.Now().Format(config.CustomTimeFormat))
					return buf.WriteString(time.Now().Format(config.CustomTimeFormat))
				case "id":
					id := req.Header.Get(echo.HeaderXRequestID)
					if id == "" {
						id = res.Header().Get(echo.HeaderXRequestID)
					}
					dumplogger = dumplogger.Str("id", id)
					return buf.WriteString(id)
				case "remote_ip":
					dumplogger = dumplogger.Str("ip", c.RealIP()+"|"+c.Request().RemoteAddr)
					return buf.WriteString(c.RealIP())
				case "host":
					dumplogger = dumplogger.Str("host", req.Host)
					return buf.WriteString(req.Host)
				case "uri":
					dumplogger = dumplogger.Str("uri", req.RequestURI)
					return buf.WriteString(req.RequestURI)
				case "method":
					dumplogger = dumplogger.Str("method", req.Method)
					return buf.WriteString(req.Method)
				case "path":
					p := req.URL.Path
					if p == "" {
						p = "/"
					}
					dumplogger = dumplogger.Str("path", p)
					return buf.WriteString(p)
				case "protocol":
					dumplogger = dumplogger.Str("proto", req.Proto)
					return buf.WriteString(req.Proto)
				case "referer":
					dumplogger = dumplogger.Str("ref", req.Referer())
					return buf.WriteString(req.Referer())
				case "user_agent":
					dumplogger = dumplogger.Str("agent", req.UserAgent())
					return buf.WriteString(req.UserAgent())
				case "status":
					n := res.Status
					s := config.colorer.Green(n)
					switch {
					case n >= 500:
						s = config.colorer.Red(n)
					case n >= 400:
						s = config.colorer.Yellow(n)
					case n >= 300:
						s = config.colorer.Cyan(n)
					}
					dumplogger = dumplogger.Str("status", s)
					return buf.WriteString(s)
				case "error":
					if err != nil {
						// Error may contain invalid JSON e.g. `"`
						b, _ := json.Marshal(err.Error())
						b = b[1 : len(b)-1]
						msg = "|ERROR|"
						// print details information when error happen
						return buf.Write(b)
					}
				case "latency":
					l := stop.Sub(start)
					dumplogger = dumplogger.Str("latency", strconv.FormatInt(int64(l), 10))
					return buf.WriteString(strconv.FormatInt(int64(l), 10))
				case "latency_human":
					dumplogger = dumplogger.Str("latency", stop.Sub(start).String())
					return buf.WriteString(stop.Sub(start).String())
				case "data_in_out":
					byteIn := req.Header.Get(echo.HeaderContentLength)
					if byteIn == "" {
						byteIn = "0"
					}
					byteOut := strconv.FormatInt(res.Size, 10)
					dumplogger = dumplogger.Str("data", fmt.Sprintf("(I: %sb|O: %sb)", byteIn, byteOut))
					return buf.WriteString(stop.Sub(start).String())
				case "bytes_in":
					cl := req.Header.Get(echo.HeaderContentLength)
					if cl == "" {
						cl = "0"
					}
					dumplogger = dumplogger.Str("in", cl)
					return buf.WriteString(cl)
				case "bytes_out":
					dumplogger = dumplogger.Str("out", strconv.FormatInt(res.Size, 10))
					return buf.WriteString(strconv.FormatInt(res.Size, 10))
				default:
					switch {
					case strings.HasPrefix(tag, "header:"):
						dumplogger = dumplogger.Str("header:"+tag[7:], c.Request().Header.Get(tag[7:]))
						return buf.Write([]byte(c.Request().Header.Get(tag[7:])))
					case strings.HasPrefix(tag, "query:"):
						dumplogger = dumplogger.Str("query:"+tag[6:], c.QueryParam(tag[6:]))
						return buf.Write([]byte(c.QueryParam(tag[6:])))
					case strings.HasPrefix(tag, "form:"):
						dumplogger = dumplogger.Str("form:"+tag[5:], c.FormValue(tag[5:]))
						return buf.Write([]byte(c.FormValue(tag[5:])))
					case strings.HasPrefix(tag, "cookie:"):
						cookie, err := c.Cookie(tag[7:])
						if err == nil {
							dumplogger = dumplogger.Str("cookie:"+tag[7:], cookie.Value)
							return buf.Write([]byte(cookie.Value))
						}
					}
				}
				return 0, nil
			}); err != nil {
				return
			}

			localLogger := dumplogger.Logger()

			switch {
			case res.Status >= http.StatusBadRequest && res.Status < http.StatusInternalServerError:
				{
					localLogger.Warn().Msg(msg)
				}
			case res.Status >= http.StatusInternalServerError:
				{
					localLogger.Error().Msg(msg)
				}
			default:
				localLogger.Info().Msg(msg)
			}

			return
		}
	}
}
