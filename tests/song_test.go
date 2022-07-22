package tests

import (
	"testing"

	"github.com/olachat/gola/testdata/song_user_favourites"
	"github.com/olachat/gola/testdata/songs"
)

func TestSong(t *testing.T) {
	songs.FetchSongByPK(1)

	s := songs.NewSong()
	s.SetHash("hash")
	err := s.Insert()
	if err != nil {
		t.Error(err)
	}

	pk := song_user_favourites.PK{
		UserId: 3,
		SongId: 99,
	}
	f := song_user_favourites.NewSongUserFavouriteWithPK(pk)
	err = f.Insert()
	if err != nil {
		t.Error(err)
	}

	if f.GetUserId() != 3 {
		t.Error("Insert non auto-increment PK failed")
	}

	updated, err := f.Update()
	if err != nil {
		t.Error(err)
	}
	if updated != false {
		t.Error("Avoid update failed")
	}
	obj := song_user_favourites.FetchSongUserFavouriteByPK(song_user_favourites.PK{
		UserId: 3,
		SongId: 99,
	})
	if obj == nil || obj.GetSongId() != 99 {
		t.Error("SongUserFavourite insert failed")
	}

	obj.SetRemark("bingo")
	ok, err := obj.Update()
	if !ok || err != nil {
		t.Error("SongUserFavourite update failed")
	}

	obj = song_user_favourites.FetchSongUserFavouriteByPK(song_user_favourites.PK{
		UserId: 3,
		SongId: 99,
	})
	if obj.GetRemark() != "bingo" {
		t.Error("SongUserFavourite update failed")
	}

	// test delete
	obj = song_user_favourites.FetchSongUserFavouriteByPK(song_user_favourites.PK{
		UserId: 3,
		SongId: 99,
	})
	err = obj.Delete()
	if err != nil {
		t.Error("SongUserFavourite delete fail")
	}
	obj = song_user_favourites.FetchSongUserFavouriteByPK(song_user_favourites.PK{
		UserId: 3,
		SongId: 99,
	})
	if obj != nil {
		t.Error("SongUserFavourite delete failed")
	}
}
