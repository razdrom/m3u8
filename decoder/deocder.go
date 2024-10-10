package decoder

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"m3u8"
)

type Decoder struct {
	playlist m3u8.Playlist
}

func NewDecoder(playlist m3u8.Playlist) *Decoder {
	return &Decoder{playlist: playlist}
}

func (d *Decoder) Decode(reader io.Reader) error {
	bufreader := bufio.NewReader(reader)
	for {
		line, _, err := bufreader.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return err
		}

		err = d.decodeLine(line)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Decoder) decodeLine(line []byte) error {
	line = bytes.TrimSpace(line)

	if d.isBlank(line) {
		return nil
	}

	if d.isTag(line) {
		k, v := d.splitTag(line)
		d.playlist.ParseTag(k, v)
		return nil
	}

	if d.isURI(line) {
		uri := string(line)
		d.playlist.HandleUri(uri)
		return nil
	}

	return fmt.Errorf("unrecognized m3u8 format")
}

func (d *Decoder) splitTag(line []byte) (string, string) {
	index := bytes.IndexByte(line, ':')
	if index == -1 {
		return "", ""
	}
	return string(line[1:index]), string(line[index+1:])
}

func (d *Decoder) isBlank(line []byte) bool {
	return len(line) == 0
}

func (d *Decoder) isTag(line []byte) bool {
	return bytes.HasPrefix(line, []byte{'#'})
}

func (d *Decoder) isURI(line []byte) bool {
	return !d.isBlank(line) && !d.isTag(line)
}
