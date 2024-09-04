package echo

import (
	"encoding/json"
	"log"

	"github.com/gonebot-dev/gonebot/messages"
	"github.com/gonebot-dev/gonebot/plugins"
)

func echoHandler(msg messages.IncomingStruct) (result messages.ResultStruct) {
	result.Text = msg.Text

	dResult, _ := json.Marshal(result)
	log.Printf("Echo: %s\n", dResult)

	return result
}
func init() {
	echo := plugins.GonePlugin{}
	echo.Name = "Echo"
	echo.Description = "Echo text message."
	echo.Handlers = append(echo.Handlers,
		plugins.GoneHandler{
			Command: []string{"echo"},
			Handler: echoHandler,
		})
	plugins.LoadPlugin(echo)
}
