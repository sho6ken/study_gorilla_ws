package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// 入口
func main() {
	// 註冊業務
	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Println(writer, "hello")
	})

	// 容器大小
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// 註冊業務
	http.HandleFunc("/ws", func(writer http.ResponseWriter, req *http.Request) {
		// 升級成ws
		conn, err := upgrader.Upgrade(writer, req, nil)

		if err != nil {
			log.Println(`upgrade to ws failed`, err)
			return
		}

		defer conn.Close()

		// 讀內文
		var value interface{}
		err = conn.ReadJSON(&value)

		if err != nil {
			log.Println(err)
			return
		}

		// 收到訊息
		fmt.Println("rcv msg from ws", value)

		if err := conn.WriteJSON("hello ws"); err != nil {
			log.Println(`write json to ws failed`, err)
			return
		}
	})

	// 開始監聽
	log.Fatal(http.ListenAndServe(":9527", nil))
}
