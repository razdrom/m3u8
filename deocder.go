package m3u8

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type Decoder struct {
	playlist Playlist
}

func NewDecoder(playlist Playlist) *Decoder {
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

		line = bytes.TrimSpace(line)

		if d.IsBlank(line) {
			continue
		}

		if d.IsTag(line) {
			k, v := d.SplitTag(line)
			d.playlist.ParseTag(k, v)
			continue
		}

		if d.IsURI(line) {
			uri := string(line)
			d.playlist.HandleUri(uri)
			continue
		}

		return fmt.Errorf("unrecognized m3u8 format")
	}

	return nil
}

func (d *Decoder) SplitTag(line []byte) (string, string) {
	index := bytes.IndexByte(line, ':')
	if index == -1 {
		return string(line), ""
	}
	return string(line[1:index]), string(line[index+1:])
}

func (d *Decoder) IsBlank(line []byte) bool {
	return len(line) == 0
}

func (d *Decoder) IsTag(line []byte) bool {
	return bytes.HasPrefix(line, []byte{'#'})
}

func (d *Decoder) IsURI(line []byte) bool {
	return !d.IsBlank(line) && !d.IsTag(line)
}
