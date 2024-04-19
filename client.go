package main

// func main() {
// 	conn, err := net.Dial("tcp", "localhost:8080")
// 	if err != nil {
// 		log.Fatal("서버에 요청하던 중 에러가 발생했습니다: ", err)
// 	}
// 	defer conn.Close()

// 	go func(conn net.Conn) {
// 		bufi := bufio.NewReader(os.Stdin)

// 		s, _ := bufi.ReadString('\n')
// 		_, err = conn.Write([]byte(s))
// 		if err != nil {
// 			log.Fatal("서버에게 메시지를 날리는 동안 에러가 발행했습니다: ", err)
// 		}

// 	}(conn)

// 	go func(conn net.Conn) {
// 		buf := make([]byte, 1024)
// 		conn.Read(buf)
// 		if err != nil {
// 			log.Fatal("서버의 메시지를 수신하다 에러가 발생했습니다.")
// 			return
// 		}

// 		fmt.Printf("server: %s", buf)
// 	}(conn)
// 	forever := make(chan struct{})

// 	<-forever
// }
