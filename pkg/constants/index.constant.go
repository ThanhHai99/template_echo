package constants

type Code string
type CachingEnum string

const (
	SUCCESS Code = "SUCCESS"
	FAILURE Code = "FAILURE"
)

const (
	ON  CachingEnum = "ON"
	OFF CachingEnum = "OFF"
)

func (c CachingEnum) String() string {
	switch c {
	case ON:
		return "ON"
	case OFF:
		return "OFF"
	}
	return ""
}

const REDIS_PREFIX_MAIN = "ms_example"
const REDIS_SEPARATOR = ":"
const REDIS_EXPIRE_TIME_DEFAULT = 604800 // (seconds) <=> 1 week
const REDIS_SUCCESS_STATUS = "successfully"
const REDIS_FAIL_STATUS = "failure"
const REDIS_STATUS_WAIT = "wait"
const REDIS_STATUS_END = "end"
const REDIS_STATUS_READY = "ready"
const REDIS_STATUS_RECONNECT = "reconnecting"
const REDIS_RECONNECT_TIME = 100 // milisecond
const REDIS_PING_SUCCESS_MESSAGE = "PONG"
