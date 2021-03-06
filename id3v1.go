package id3v1

import (
	"bytes"
	"io"
	"errors"
)

type Id3v1Tag struct {
	title, artist, album, comment, year string
	trackNumber                         int
}

func (tag *Id3v1Tag) Title() string {
	return tag.title
}

func (tag *Id3v1Tag) Artist() string {
	return tag.artist
}

func (tag *Id3v1Tag) Album() string {
	return tag.album
}

func (tag *Id3v1Tag) Comment() string {
	return tag.comment
}

func (tag *Id3v1Tag) Year() string {
	return tag.year
}

func (tag *Id3v1Tag) TrackNumber() int {
	return tag.trackNumber
}

func ReadTag(r io.ReadSeeker) (Id3v1Tag, error) {
	tagBytes := make([]byte, 128)
	r.Seek(-128, 2)
	r.Read(tagBytes)

	header := string(tagBytes[:3])
	if header != "TAG" {
		return Id3v1Tag{}, errors.New("Source does not have tag in standard location")
	} else {
		return createTag(tagBytes), nil
	}
}

func createTag(tagBytes []byte) Id3v1Tag {
	tag := new(Id3v1Tag)
	tag.title = trimmedString(tagBytes[3:33])
	tag.artist = trimmedString(tagBytes[33:63])
	tag.album = trimmedString(tagBytes[63:93])
	tag.year = trimmedString(tagBytes[93:97])
	comment, trackNo := readCommentAndOrTrackNo(tagBytes[97:127])
	tag.comment = comment
	tag.trackNumber = trackNo
	return *tag
}

func trimmedString(b []byte) string {
	return string(bytes.TrimRight(b, "\x00"))
}

func readCommentAndOrTrackNo(b []byte) (string, int) {
	if b[28] == 0 {
		return trimmedString(b[:29]), int(b[29])
	} else {
		return trimmedString(b), -1
	}
}
