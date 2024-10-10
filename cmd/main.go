package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/razdrom/m3u8"
	"github.com/razdrom/m3u8/decoder"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("ошибка при получении рабочей директории: %w", err)
	}

	file, err := os.Open(path.Join(dir, "test/assets", "media_001.m3u8"))
	if err != nil {
		log.Fatal("ошибка при открытии файла: %w", err)
	}

	defer file.Close()

	media := &m3u8.MediaPlaylist{}
	decoder := decoder.NewDecoder(media)
	err = decoder.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(media)
}
