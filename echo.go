package echo

import (
	"strings"

	"github.com/gonebot-dev/gonebot/message"
	"github.com/gonebot-dev/gonebot/plugin"
	"github.com/gonebot-dev/gonebot/plugin/handler"
)

var Echo plugin.GonePlugin

var echo_prefix string = "echo"

func echoMatcher(msg message.Message) bool {
	for _, seg := range msg.Segments {
		if seg.Type == "text" {
			if strings.HasPrefix(seg.Content, echo_prefix) {
				return true
			}
		}
	}
	return false
}

func echoHandler(incomingMsg message.Message, resultMsg *message.Message) (block bool) {
	for _, seg := range incomingMsg.Segments {
		if seg.Type == "text" {
			if strings.HasPrefix(seg.Content, echo_prefix) {
				seg.Content = strings.Trim(seg.Content, echo_prefix)
				break
			}
		}
	}
	resultMsg.Segments = append(resultMsg.Segments, incomingMsg.Segments...)
	return true
}

func init() {
	Echo.Name = "Echo"
	Echo.Version = "v2.0.alpha"
	Echo.Description = "Reply the same message of what you have sent"

	Echo.Handlers = append(Echo.Handlers, handler.GoneHandler{
		Matcher: echoMatcher,
		Handler: echoHandler,
	})
}
