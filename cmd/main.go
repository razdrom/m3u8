package main

import (
	"fmt"
	"log"
	"m3u8"
	"m3u8/decoder"
	"os"
	"path"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("ошибка при получении рабочей директории: %w", err)
	}

	file, err := os.Open(path.Join(dir, "pkg/hls/mocks", "media.m3u8"))
	if err != nil {
		log.Fatal("ошибка при открытии файла: %w", err)
	}

	defer file.Close()

	master := &m3u8.MasterPlaylist{}
	decoder := decoder.NewDecoder(master)
	err = decoder.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(master)
}
