package hydrachat

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"sync"
	"testing"
	"time"
)

var once sync.Once

func chatServerFunc(t *testing.T) func() {
	return func(){
		t.Log("Starting Hydra Chat Server... ")
		if err := Run(":2300"); err != nil {
			t.Error("Could not start chat server ", err)
			return
		}

		t.Log("Started Hydra Chat Server...")
	}
}

func TestRun(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode... ")
	}
	t.Log("Testing hydra chat send and receive... ")
	f := chatServerFunc(t)

	go once.Do(f)

	time.Sleep(1 * time.Second)

	rand.Seed(time.Now().UnixNano())
	name := fmt.Sprintf("Anonymous%d", rand.Intn(400))

	t.Logf("Hello %s, connecting to hydra chat system... \n", name)
	conn, err := net.Dial("tcp", "127.0.0.1:2300")
	if err != nil {
		t.Fatal("Could not connect to hydra chat system", err)
	}
	t.Log("Connected to hydra chat system")
	name += ":"
	defer conn.Close()
	msgCh := make(chan string)

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			recvmsg := scanner.Text()
			sentmsg := <-msgCh
			if strings.Compare(recvmsg, sentmsg) != 0 {
				t.Errorf("Chat message %s does not match %s", recvmsg, sentmsg)
			}
		}
	}()

	for i := 0; i <= 10; i++ {
		msgBody := fmt.Sprintf("RandomMessage %d", rand.Intn(400))
		msg := name + msgBody
		_, err = fmt.Fprintf(conn, msg + "\n")
		if err != nil {
			t.Error(err)
			return
		}
		msgCh <- msg
	}
}

func TestServerConnection(t *testing.T) {
	t.Log("Test Hydra Chat receive messages... ")
	f := chatServerFunc(t)
	go once.Do(f)

	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:2300")
	if err != nil {
		t.Fatal("Could not connect to Hydra Chat System", err)
	}

	defer conn.Close()
}
