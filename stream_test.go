package zdpgo_redis

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRedis_PubStream(t *testing.T) {
	r := prepareRedis()

	for i := 0; i < 3000; i++ {
		values := map[string]interface{}{
			"whatHappened": string("ticket received"),
			"ticketID":     int(rand.Intn(100000000)),
			"ticketData":   string("some ticket data"),
		}
		r.PubStream(PubStreamConfig{
			Subject: "test",
			Values:  values,
		})
	}
}

func TestRedis_SubStream(t *testing.T) {
	r := prepareRedis()
	r.SubStream(SubStreamConfig{
		Subject:           "test",
		ConsumerGroupName: "test_group",
		HandStreamFunc:    handleNewTicket,
	})
}

func handleNewTicket(values map[string]interface{}) error {
	fmt.Println("正在消费数据：", "values", values)
	return nil
}
