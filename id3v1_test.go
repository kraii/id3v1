package id3v1

import (
	. "gopkg.in/check.v1"
	"os"
	"testing"
)

func TestId3s(t *testing.T) { TestingT(t) }

type Id3v1TagSuite struct{}

var _ = Suite(&Id3v1TagSuite{})
var file *os.File

func (s *Id3v1TagSuite) SetUpTest(c *C) {
	var err error
	file, err = os.Open("_testdata/spice.mp3")
	c.Assert(err, Equals, nil)
}

func (s *Id3v1TagSuite) TestReadWellFormedTag(c *C) {
	tag := ReadTag(file)
	c.Check(tag.artist, Equals, "Xander")
	c.Check(tag.title, Equals, "Spice")
	c.Check(tag.album, Equals, "Things")
	c.Check(tag.year, Equals, "2015")
	c.Check(tag.comment, Equals, "say -v Xander")
	c.Check(tag.trackNumber, Equals, 1)
}

func (s *Id3v1TagSuite) TearDownTest(c *C) {
	file.Close()
}
