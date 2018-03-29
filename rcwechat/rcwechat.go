package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
)

func rcservice(ctx *context.Context) {
	//配置微信参数
	config := &wechat.Config{
		AppID:          "wx08b29a151ffdaec1",
		AppSecret:      "b133de6ed5b61b10515c87a988608a2e",
		Token:          "levelup",
		EncodingAESKey: "ECA86A47BF33BC640B0E04F0EBC8827A083100F92101518C1FAEFE8AB46AC6CC",
	}
	wc := wechat.NewWechat(config)

	// 传入request和responseWriter
	server := wc.GetServer(ctx.Request, ctx.ResponseWriter)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

		if msg.MsgType == message.MsgTypeEvent {

			switch msg.Event {
			case message.EventClick:

				text := message.NewText(msg.EventKey)
				return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

			default:

			}

		}

		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}

func main() {
	beego.Any("/", rcservice)
	beego.Run(":8001")
}
