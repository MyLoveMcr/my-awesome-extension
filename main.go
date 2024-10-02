package main

import (
	_ "github.com/MyLoveMcr/my-awesome-extension/hiddify_extension"

	"github.com/hiddify/hiddify-core/extension/server"
)

func main() {
	server.StartTestExtensionServer()
}
