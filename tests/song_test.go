package tests

import (
	"testing"

	"github.com/olachat/gola/coredb"
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

	f := song_user_favourites.NewSongUserFavourite()
	f.SetUserId(3)
	f.SetSongId(99)
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

	/* TODO: This case is not yet supported, as it actually result in SQL
		update song_user_favourites set user_id=4 where user_id=4

	which is definitely avoid, the correct sql should be:
		update song_user_favourites set user_id=4 where user_id=3

	but it would required specially handling when update non-autoincremented PK
	I guess this is an very odd case in actual applications, may support later
	Just keep test cases here, if they failed, guess it means PK updated is supported
	And these test case should be udpated.
	*/
	f.SetUserId(4)
	updated, err = f.Update()
	if err != coredb.ErrAvoidUpdate {
		t.Error("Avoid update detect failed")
	}
	if updated != false || f.GetUserId() != 4 {
		t.Error("SongUserFavourite void update failed")
	}

	obj = song_user_favourites.FetchSongUserFavouriteByPK(song_user_favourites.PK{
		UserId: 4,
		SongId: 99,
	})
	if obj != nil {
		t.Error("SongUserFavourite update PK failed")
	}
}
