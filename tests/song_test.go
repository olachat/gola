package tests

import (
	"testing"

	"github.com/olachat/gola/testdata/songs"
)

func TestSong(t *testing.T) {
	songs.FetchSongByPK(1)
}
