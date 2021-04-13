package main

//server
import (
	"fmt"
	"net"
	"time"
)

const ip string = "localhost:13487"
const Maxsize = 512

// TCP service

func initlisten() {
	wg.Add(1)
	// 1.监听本地端口
	linser, err := net.Listen("tcp", ip)
	if err != nil {
		fmt.Println("faild, error of:", err)
		return
	} else {
		fmt.Println("--------Load Successed!--------\nListening at ", ip)
	}
	for {
		fmt.Println("^_^")
		// 2.等待客户端连接
		client_conn, err := linser.Accept()
		if err != nil {
			fmt.Println("conn error of:", err)
			return
		} else {
			now := time.Now()
			mytime := now.Format("2006-01-02 15:04:05")
			s := client_conn.RemoteAddr().String()
			fmt.Println("-->Connect from :", s, "\nAt:", mytime)
		}

		// 3.接收客户端信息
		var msg [Maxsize]byte

		n, err := client_conn.Read(msg[:])
		if err != nil {
			fmt.Println("Read error of:", err)
			return
		} else {
			myfun(msg, n)
		}
		// checkconn(client_conn)
		// 4.给客户端回应消息
		_, err = client_conn.Write([]byte("Accept Successed!"))
		if err != nil {
			fmt.Println("Write err of:", err)
			return
		}
		client_conn.Close()
	}
	linser.Close()
	wg.Done()
}

//对数据做一些什么吧
func myfun(msg [Maxsize]byte, n int) {
	fmt.Println("The client msg of:", string(msg[:n]))
	fmt.Println("------------------------------------")
}

func checkconn(conn net.Conn) {
	s := conn.RemoteAddr().String()
	fmt.Println(s)
}
