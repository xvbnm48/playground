// You can edit this code!
// Click here and start typing.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type CreateRoomBody struct {
	TopicID      string
	ReferenceID  string   `json:"reference_id"`
	Participants []string `json:"participants"`
}

func main() {
	payload := CreateRoomBody{
		TopicID:      "112",
		ReferenceID:  "1212",
		Participants: []string{"as", "dr"},
	}

	fmt.Println(payload.TopicID)
	bytePayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	_ := bytes.NewBuffer(bytePayload)
	net, _ := net.Listen("tcp", "localhost:8080")
}
