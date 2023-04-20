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

const REDIS_PREFIX_MAIN string = "ms_example"
const REDIS_SEPARATOR string = ":"
const REDIS_EXPIRE_TIME_DEFAULT uint32 = 604800 // (seconds) <=> 1 week
const REDIS_SUCCESS_STATUS string = "successfully"
const REDIS_FAIL_STATUS string = "failure"
const REDIS_STATUS_WAIT string = "wait"
const REDIS_STATUS_END string = "end"
const REDIS_STATUS_READY string = "ready"
const REDIS_STATUS_RECONNECT string = "reconnecting"
const REDIS_RECONNECT_TIME uint8 = 100 // milisecond
const REDIS_PING_SUCCESS_MESSAGE string = "PONG"
