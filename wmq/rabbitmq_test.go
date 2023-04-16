package wmq

import (
	"fmt"
	"os"
	"testing"
)

func TestInitRabbitMQClient(t *testing.T) {
	conn, err := InitRabbitMQClient("guest", "guest", "localhost", "5672")
	if err != nil {
		t.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("mq connect success...")
	fmt.Println(conn)
}
