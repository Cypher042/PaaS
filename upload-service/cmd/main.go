package main

import (
	"github.com/Cypher042/PaaS/upload-service/internal/uploader"
)


func main() {

	uploader.Upload("https://github.com/Cypher042/SadServers.md", "testingdelete")
}
