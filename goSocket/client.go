package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/micmonay/keybd_event"
)

func kbEvent(inp string) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	// kb.AddKey(keybd_event.VK_B)
	lower := strings.ToLower(inp)
	split := strings.Split(lower, "")
	for i := 0; i < len(inp); i++ {
		if split[i] != "\n" {
			kb.AddKey(Binds[split[i]])
		}
	}
	err = kb.Launching()
	if err != nil {
		panic(err)
	}
}

func SocketClient(ip string, port int) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	fmt.Print("Connected! PogChamp")
	defer conn.Close()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
		kbEvent(message)
	}

}

func main() {

	var (
		ip   = "192.168.1.69"
		port = 80
	)

	SocketClient(ip, port)

}
