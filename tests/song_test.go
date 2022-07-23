package tests

import (
	"fmt"
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
	obj := song_user_favourites.FetchSongUserFavouriteByPK(pk)
	if obj == nil || obj.GetSongId() != 99 {
		t.Error("SongUserFavourite insert failed")
	}

	obj.SetRemark("bingo")
	ok, err := obj.Update()
	if !ok || err != nil {
		t.Error("SongUserFavourite update failed")
	}

	obj = song_user_favourites.FetchSongUserFavouriteByPK(pk)
	if obj.GetRemark() != "bingo" {
		t.Error("SongUserFavourite update failed")
	}
}
func TestSongDelte(t *testing.T) {
	pk := song_user_favourites.PK{
		UserId: 3,
		SongId: 99,
	}
	obj := song_user_favourites.FetchSongUserFavouriteByPK(pk)
	err := obj.Delete()
	if err != nil {
		t.Error("SongUserFavourite delete fail")
	}
	obj = song_user_favourites.FetchSongUserFavouriteByPK(pk)
	if obj != nil {
		t.Error("SongUserFavourite delete failed")
	}

	obj = song_user_favourites.NewSongUserFavouriteWithPK(pk)
	obj.SetRemark("remark")
	err = obj.Insert()
	if err != nil {
		t.Error(err)
	}
	obj = song_user_favourites.FetchSongUserFavouriteByPK(pk)
	if obj.GetRemark() != "remark" {
		t.Error("SongUserFavourite reinsert failed")
	}

	partialObj := song_user_favourites.FetchByPK[struct {
		song_user_favourites.SongId
		song_user_favourites.UserId
		song_user_favourites.Remark
	}](pk)
	if partialObj.GetRemark() != "remark" {
		t.Error("SongUserFavourite reinsert failed")
	}

	err = song_user_favourites.Delete(partialObj)
	if err != nil {
		t.Error(err)
	}
	obj = song_user_favourites.FetchSongUserFavouriteByPK(pk)
	if obj != nil {
		t.Error("SongUserFavourite partial delete failed")
	}
}

func TestSongUpdate(t *testing.T) {
	s := songs.NewSong()
	s.SetHash("hashhash")
	err := s.Insert()
	if err != nil {
		t.Error(err)
	}

	id := s.GetId()

	obj := songs.FetchByPK[struct {
		songs.Id
		songs.Hash
	}](id)
	obj.SetHash("bingo")
	ok, err := songs.Update(obj)
	if !ok || err != nil {
		t.Error("Partail update failed")
	}

	s = songs.FetchSongByPK(id)
	if s.GetHash() != "bingo" {
		t.Error("Partail update failed")
	}

	ok, err = songs.Update(obj)
	if ok || err != nil {
		fmt.Printf("ok: %v\n", ok)
		fmt.Printf("err: %v\n", err)
		t.Error("Partail void update failed")
	}
}
