package main

// import (
// 	"flag"
// 	"fmt"
// 	"os"

// 	tl "github.com/JoelOtter/termloop"
// )

// type MovingText struct {
// 	*tl.Text
// }

// func (m *MovingText) Tick(ev tl.Event) {
// 	if ev.Type == tl.EventKey {
// 		x, y := m.Position()

// 		switch ev.Key {
// 		case tl.KeyArrowRight:
// 			c.Send("moved", m)
// 			x += 1
// 		case tl.KeyArrowLeft:
// 			c.Send("moved", m)
// 			x -= 1
// 		case tl.KeyArrowUp:
// 			c.Send("moved", m)
// 			y -= 1
// 		case tl.KeyArrowDown:
// 			c.Send("moved", m)
// 			y += 1
// 		}

// 	}
// }

// func main() {
// 	if len(os.Args) == 1 {
// 		fmt.Printf("<Welcome to Forward!>\n")
// 		fmt.Printf("Please use the following flags:\n\n")
// 		fmt.Printf("- port: 자신이 원하는 이름을 입력해주세요.\n")
// 		fmt.Printf("- server: 서버의 연결 주소를 입력해주세요.\n\n")
// 		os.Exit(1)
// 	}

// 	nameP := flag.String("name", "default", "자신이 원하는 이름을 입력해주세요.")
// 	// serverP := flag.String("server", "localhost:9190", "서버의 연결 주소를 입력해주세요.")
// 	flag.Parse()

// 	g := tl.NewGame()
// 	g.Screen().SetFps(60)
// 	g.Screen().AddEntity(&MovingText{tl.NewText(0, 0, os.Args[1], tl.ColorWhite, tl.ColorBlue)})
// 	g.Start()
// }
