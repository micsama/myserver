package main

import (
	"fmt"
	"net"
	"runtime"
	"time"
)

const Maxsize = 512

func connectip(bot string, ip string) {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
	// 1.连接服务端
	fmt.Println("start tcp")
	conn_service, err := net.Dial("tcp", ip+":8086")
	// conn_service, err := net.Dial("tcp", ip)
	fmt.Println("tcpok")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 2.给服务端发生消息
	// fmt.Printf("args[%v]=[%v]\n", k, v)
	v := "hi"
	message := []byte(v)
	_, err = conn_service.Write(message)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Print("Send success!")
	}

	// 3.接收服务端消息
	var msg [Maxsize]byte // 声明一个接收信息的变量
	for {
		n, err := conn_service.Read(msg[:])
		if err != nil {
			fmt.Println("Read error of:", err)
			return
		}
		myfun(msg, n)
	}
	// 4.关闭连接
	conn_service.Close()
}

// TCP service
// func initlisten() {
// 	wg.Add(1)
// 	// 1.监听本地端口
// 	linser, err := net.Listen("tcp", ip)
// 	if err != nil {
// 		fmt.Println("faild, error of:", err)
// 		return
// 	} else {
// 		fmt.Println("--------Load Successed!--------\nListening at ", ip)
// 	}
// 	for {
// 		fmt.Println("^_^")
// 		// 2.等待客户端连接
// 		client_conn, err := linser.Accept()
// 		if err != nil {
// 			fmt.Println("conn error of:", err)
// 			return
// 		} else {
// 			now := time.Now()
// 			mytime := now.Format("2006-01-02 15:04:05")
// 			s := client_conn.RemoteAddr().String()
// 			fmt.Println("-->Connect from :", s, "\nAt:", mytime)
// 		}
// 		// 3.接收客户端信息
// 		var msg [Maxsize]byte
// 		n, err := client_conn.Read(msg[:])
// 		if err != nil {
// 			fmt.Println("Read error of:", err)
// 			return
// 		} else {
// 			myfun(msg, n)
// 		}
// 		// checkconn(client_conn)
// 		// 4.给客户端回应消息
// 		_, err = client_conn.Write([]byte("Accept Successed!"))
// 		if err != nil {
// 			fmt.Println("Write err of:", err)
// 			return
// 		}
// 		client_conn.Close()
// 	}
// 	linser.Close()
// 	wg.Done()
// }

//对数据做一些什么吧
func myfun(msg [Maxsize]byte, n int) {
	fmt.Println("The client msg of:", string(msg[:n]))
	fmt.Println("------------------------------------")
	type mytdlog struct {
		Date string
		Time string
		Name string
		Num  int
		Area string
	}
	var a mytdlog
	s := string(msg[:n])
	fmt.Println(s)
	data := time.Now().Local().Format("2006-01-02")
	userdb = userdb.Table("tdlog")
	a.Area = s
	a.Date = data
	a.Time = time.Now().Local().Format("15:04")
	a.Num = 100
	a.Name = "疫苗"
	if err := userdb.Create(a).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	} else {
		fmt.Println("sucess")
	}
}
func checkconn(conn net.Conn) {
	type mytdlog struct {
		Date string
		Time string
		Name string
		Num  int
		Area string
	}
	var a mytdlog
	s := conn.RemoteAddr().String()
	data := time.Now().Local().Format("2006-01-02")
	userdb = userdb.Table("tdlog")
	a.Area = s
	a.Date = data
	a.Time = time.Now().Local().Format("15:04")
	a.Num = 100
	a.Name = "疫苗"
	if err := userdb.Create(a).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}
}
