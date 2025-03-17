package main

import (
	"context"
	"log"
	"time"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
	"github.com/razzle131/hs316go/internal/opc"
	"github.com/razzle131/hs316go/internal/tags"
)

const connString = "opc.tcp://10.160.160.61:4840"

func main() {
	ctx := context.TODO()

	client := opcua.NewClient(connString, opcua.SecurityMode(ua.MessageSecurityModeNone), opcua.DialTimeout(time.Second*5))

	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	opc.GetNodeValue(tags.InputGripperStart, client)

	opc.WriteNodeValue(tags.OutputConveyorRight, "true", client)
	time.Sleep(time.Second * 5)
	opc.WriteNodeValue(tags.OutputConveyorRight, "false", client)
}
