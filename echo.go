package echo

import (
	"github.com/gonebot-dev/gonebot/adapter"
	"github.com/gonebot-dev/gonebot/message"
	"github.com/gonebot-dev/gonebot/plugin"
	"github.com/gonebot-dev/gonebot/rule"
)

var Echo plugin.GonePlugin

func init() {
	Echo.Name = "Echo"
	Echo.Version = "v0.0.1"
	Echo.Description = "Reply the same message of what you have sent"

	Echo.Handlers = append(Echo.Handlers, plugin.GoneHandler{
		Rules: []rule.FilterBundle{{Filters: []rule.FilterRule{rule.ToMe(), rule.Command([]string{"echo"})}}},
		Handler: func(a *adapter.Adapter, msg message.Message) bool {
			reply := msg
			reply.Sender = reply.Self
			reply.Receiver = msg.Sender
			a.SendChannel.Push(reply, false)
			return true
		},
	})
}
