package echo

import (
	"strings"

	"github.com/gonebot-dev/gonebot/message"
	"github.com/gonebot-dev/gonebot/plugin"
	"github.com/gonebot-dev/gonebot/plugin/handler"
)

var EchoPlugin plugin.GonePlugin

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
	EchoPlugin.Name = "Echo"
	EchoPlugin.Version = "v2.0.alpha"
	EchoPlugin.Description = "Reply the same message of what you have sent"

	EchoPlugin.Handlers = append(EchoPlugin.Handlers, handler.GoneHandler{
		Matcher: echoMatcher,
		Handler: echoHandler,
	})
}

func GetPlugin() plugin.GonePlugin {
	return EchoPlugin
}
