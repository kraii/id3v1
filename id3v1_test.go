package id3v1

import (
	. "gopkg.in/check.v1"
	"os"
	"testing"
)

func TestId3s(t *testing.T) { TestingT(t) }

type Id3v1TagSuite struct{}

var _ = Suite(&Id3v1TagSuite{})

func readFile(fn string, c *C) *os.File {
	file, err := os.Open(fn)
	c.Assert(err, Equals, nil)
	return file
}

func (s *Id3v1TagSuite) TestReadWellFormedTag(c *C) {
	file := readFile("_testdata/spice.mp3", c)
	defer file.Close()

	tag, err := ReadTag(file)

	c.Assert(err, Equals, nil)
	c.Check(tag.Artist(), Equals, "Xander")
	c.Check(tag.Title(), Equals, "Spice")
	c.Check(tag.Album(), Equals, "Things")
	c.Check(tag.Year(), Equals, "2015")
	c.Check(tag.Comment(), Equals, "say -v Xander")
	c.Check(tag.TrackNumber(), Equals, 1)
}

func (s *Id3v1TagSuite) TestReadNoTag(c *C) {
	file := readFile("_testdata/tagless-batman.mp3", c)
	defer file.Close()

	tag, err := ReadTag(file)
	
	c.Check(tag, Equals, Id3v1Tag{})
	c.Check(err.Error(), Equals, "Source does not have tag in standard location")
}
