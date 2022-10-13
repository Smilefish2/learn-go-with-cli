package gobyexample_com

import (
	"fmt"
	"go-labs/cmd"

	"github.com/spf13/cobra"
)

// Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.
// 通道(channels) 是连接多个协程的管道。 你可以从一个协程将值发送到通道，然后在另一个协程中接收。

// channelsCmd represents the channels command
var channelsCmd = &cobra.Command{
	Use:   "go-by-example:channels",
	Short: "https://gobyexample.com/channels",
	Run: func(cmd *cobra.Command, args []string) {

		// Create a new channel with make(chan val-type). Channels are typed by the values they convey.
		// 使用 make(chan val-type) 创建一个新的通道。 通道类型就是他们需要传递值的类型。
		messages := make(chan string)

		// Send a value into a channel using the channel <- syntax. Here we send "ping" to the messages channel we made above, from a new goroutine.
		// 使用 channel <- 语法 发送 一个新的值到通道中。 这里我们在一个新的协程中发送 "ping" 到上面创建的 messages 通道中。
		go func() { messages <- "ping1" }()

		// The <-channel syntax receives a value from the channel. Here we’ll receive the "ping" message we sent above and print it out.
		// 使用 <-channel 语法从通道中 接收 一个值。 这里我们会收到在上面发送的 "ping" 消息并将其打印出来。
		msg := <-messages
		fmt.Println(msg)

		// When we run the program the "ping" message is successfully passed from one goroutine to another via our channel.
		// 我们运行程序时，通过通道， 成功的将消息 "ping" 从一个协程传送到了另一个协程中。

		// By default sends and receives block until both the sender and receiver are ready. This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.
		// 默认发送和接收操作是阻塞的，直到发送方和接收方都就绪。 这个特性允许我们，不使用任何其它的同步操作， 就可以在程序结尾处等待消息 "ping"。
	},
}

func init() {
	cmd.RootCmd.AddCommand(channelsCmd)
}
