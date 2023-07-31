package hanweb

import "fmt"

const authHeader = "Authorization"

type CustomHookEvent int32

const (
	EventDefault CustomHookEvent = iota
	EventConnectionClientInfo
	EventSubscribeSuccess
	EventSubscribeFail
	EventConnectionQuality
)

func (c CustomHookEvent) String() string {
	switch c {
	case EventDefault:
		return fmt.Sprintf("%d", int(c))
	case EventConnectionClientInfo:
		return "CONNECTION_CLIENT_INFO"
	case EventSubscribeSuccess:
		return "SUBSCRIBE_SUCCESS"
	case EventSubscribeFail:
		return "SUBSCRIBE_FAIL"
	case EventConnectionQuality:
		return "CONNECTION_QUALITY"
	default:
		return fmt.Sprintf("%d", int(c))
	}
}
