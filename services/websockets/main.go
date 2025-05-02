package main

import (
	"github.com/nitrictech/go-sdk/nitric"
	"github.com/nitrictech/go-sdk/nitric/keyvalue"
)

func main() {
	// create a WS API name "public"
	ws := nitric.NewWebsocket("public")

	// init a KV store named "connections" with Get, Set, and Delete permissions
	connections := nitric.NewKv("connections").Allow(keyvalue.KvStoreGet, keyvalue.KvStoreSet, keyvalue.KvStoreDelete)

	// add event handlers here

	nitric.Run()
}
