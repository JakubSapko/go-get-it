package parser

import (
	"fmt"
	"os"

	"github.com/jackpal/bencode-go"
)

type bencodedTorrent struct {
	Announce     string `bencode:"announce"`
	Comment      string `bencode:"comments"`
	CreationDate int    `bencode:"creation date"`
	Info         bencodedTorrentInfo
}

type bencodedTorrentInfo struct {
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`
}

type Torrent struct {
	Announce    string
	InfoHash    [20]byte   //SHA-1 hashes are 20 bytes long
	Pieces      [][20]byte //Torrent pieces is a byte stream of SHA-1 hashes, each representing a piece of a file we want to download
	PieceLength int
	Length      int
	Name        string
}

func (bt *bencodedTorrent) toString() string {
	return fmt.Sprint("Announce: ", bt.Announce, "\nComment: ", bt.Comment, "\nCreation Date: ", bt.CreationDate, "\nLength: ", bt.Info.Length, "\nName: ", bt.Info.Name, "\nPiece Length: ", bt.Info.PieceLength, "\nPieces (head): ", bt.Info.Pieces[0:20])
}

func (bt *bencodedTorrent) ToTorrent() (Torrent, error) {
	torrent := Torrent{}
	fmt.Println(bt.toString())
	return torrent, nil
}

func OpenTorrent(path string) (*bencodedTorrent, error) {
	torrent := bencodedTorrent{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	err = bencode.Unmarshal(file, &torrent)
	return &torrent, err
}
