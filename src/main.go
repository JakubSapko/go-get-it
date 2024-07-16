package main

import (
	"gogetit/parser"
)

func main() {
	fileData, err := parser.OpenTorrent("parser/files/warzone2100_macOS_universal.zip.torrent")

	if err != nil {
		panic(1)
	}

	fileData.ToTorrent()

	//fmt.Print(fileData)
}
