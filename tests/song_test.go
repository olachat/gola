package tests

import (
	"fmt"
	"testing"

	"encoding/json"

	"github.com/olachat/gola/coredb"
	"github.com/olachat/gola/golalib/testdata/song_user_favourites"
	"github.com/olachat/gola/golalib/testdata/songs"
)

func TestSong(t *testing.T) {
	s := songs.New()
	s.SetRank(5)
	s.SetType(songs.SongType1x2B9)
	s.SetHash("hash")
	err := s.Insert()
	if err != nil {
		t.Error(err)
	}

	id := s.GetId()
	if id != 1 {
		t.Error("Insert song error")
	}

	s = songs.FetchByPK(id)
	if s.GetRank() != 5 || s.GetHash() != "hash" || s.GetType() != songs.SongType1x2B9 {
		t.Error("Re-fetch song error")
	}

	s.SetType(songs.SongTypeEmpty)
	s.Update()
	s = songs.FetchByPK(id)
	if s.GetType() != songs.SongTypeEmpty {
		t.Error("Song update error")
	}

	pk := song_user_favourites.PK{
		UserId: 3,
		SongId: 99,
	}
	f := song_user_favourites.NewWithPK(pk)
	err = f.Insert()
	if err != nil {
		t.Error(err)
	}

	if f.GetUserId() != 3 {
		t.Error("Insert non auto-increment PK failed")
	}

	f2 := song_user_favourites.NewWithPK(pk)
	err = f2.Insert()
	if err != coredb.ErrAvoidInsert {
		t.Error("Repeat insert must fail")
	}

	updated, err := f.Update()
	if err != nil {
		t.Error(err)
	}
	if updated != false {
		t.Error("Avoid update failed")
	}
	obj := song_user_favourites.FetchByPK(pk)
	if obj == nil || obj.GetSongId() != 99 {
		t.Error("SongUserFavourite insert failed")
	}

	obj.SetRemark("bingo")
	ok, err := obj.Update()
	if !ok || err != nil {
		t.Error("SongUserFavourite update failed")
	}
	obj = song_user_favourites.FetchByPK(pk)
	if obj.GetRemark() != "bingo" {
		t.Error("SongUserFavourite update failed")
	}
}

func TestSongBoolUpdate(t *testing.T) {
	pk := song_user_favourites.PK{
		UserId: 3,
		SongId: 99,
	}
	obj := song_user_favourites.FetchByPK(pk)
	if obj.GetRemark() != "bingo" {
		t.Error("SongUserFavourite update failed")
	}
	if obj.GetIsFavourite() != true {
		t.Error("SongUserFavourite GetIsFavourite failed")
	}
	obj.SetIsFavourite(false)
	obj.Update()
	obj = song_user_favourites.FetchByPK(pk)
	if obj.GetIsFavourite() != false {
		t.Error("SongUserFavourite GetIsFavourite update false failed")
	}

	obj.SetIsFavourite(true)
	obj.Update()
	obj = song_user_favourites.FetchByPK(pk)
	if obj.GetIsFavourite() != true {
		t.Error("SongUserFavourite GetIsFavourite update true failed")
	}

	pb := song_user_favourites.FetchFieldsByPK[struct {
		song_user_favourites.SongId
		song_user_favourites.UserId
		song_user_favourites.IsFavourite
	}](pk)
	if obj.GetIsFavourite() != true {
		t.Error("SongUserFavourite GetIsFavourite failed")
	}
	ok := pb.SetIsFavourite(false)
	if !ok {
		t.Error("SetIsFavourite failed")
	}
	ok, err := song_user_favourites.Update(pb)
	if !ok || err != nil {
		t.Error("pb update failed")
	}
	obj = song_user_favourites.FetchByPK(pk)
	if obj.GetIsFavourite() != false {
		t.Error("SongUserFavourite GetIsFavourite update failed")
	}
}

func TestSongDelete(t *testing.T) {
	pk := song_user_favourites.PK{
		UserId: 3,
		SongId: 99,
	}
	obj := song_user_favourites.FetchByPK(pk)
	err := song_user_favourites.DeleteByPK(pk)
	if err != nil {
		t.Error("SongUserFavourite delete fail")
	}
	obj = song_user_favourites.FetchByPK(pk)
	if obj != nil {
		t.Error("SongUserFavourite delete failed")
	}

	obj = song_user_favourites.NewWithPK(pk)
	obj.SetRemark("remark")
	err = obj.Insert()
	if err != nil {
		t.Error(err)
	}
	obj = song_user_favourites.FetchByPK(pk)
	if obj.GetRemark() != "remark" {
		t.Error("SongUserFavourite reinsert failed")
	}

	partialObj := song_user_favourites.FetchFieldsByPK[struct {
		song_user_favourites.SongId
		song_user_favourites.UserId
		song_user_favourites.Remark
	}](pk)
	if partialObj.GetRemark() != "remark" {
		t.Error("SongUserFavourite reinsert failed")
	}

	err = song_user_favourites.DeleteByPK(pk)
	if err != nil {
		t.Error(err)
	}
	obj = song_user_favourites.FetchByPK(pk)
	if obj != nil {
		t.Error("SongUserFavourite partial delete failed")
	}
}

func TestSongUpdate(t *testing.T) {
	s := songs.New()
	s.SetHash("hashhash")
	err := s.Insert()
	if err != nil {
		t.Error(err)
	}

	id := s.GetId()

	obj := songs.FetchFieldsByPK[struct {
		songs.Id
		songs.Hash
	}](id)
	obj.SetHash("bingo")
	ok, err := songs.Update(obj)
	if !ok || err != nil {
		t.Error("Partail update failed")
	}

	s = songs.FetchByPK(id)
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
func TestSongInsertWithPK(t *testing.T) {
	id := uint(100)
	s := songs.NewWithPK(id)
	s.SetHash("100")
	err := s.Insert()
	if err != nil {
		t.Error(err)
	}
	s = songs.FetchByPK(id)

	if s == nil {
		t.Error("Insert with pk failed")
	}
	s = songs.New()
	s.SetHash("hashhash")
	err = s.Insert()
	if err != nil {
		t.Error(err)
	}
	if s.GetId() != id+1 {
		t.Error("Insert without pk after pk failed")
	}

	id = uint(99)
	s = songs.NewWithPK(id)
	s.SetHash("99")
	err = s.Insert()
	if err != nil {
		t.Error(err)
	}
	s = songs.FetchByPK(id)

	if s == nil {
		t.Error("Reinsert with pk failed")
	}
	s = songs.New()
	s.SetHash("hashhash")
	err = s.Insert()
	if err != nil {
		t.Error(err)
	}
	if s.GetId() != 102 {
		t.Error("Insert without pk after pk failed")
	}
}

func TestSongJSONEncode(t *testing.T) {
	s := songs.New()
	s.SetRank(5)
	s.SetType(songs.SongType1x2B9)
	s.SetHash("hash")

	jsondata, err := json.Marshal(&s)
	if err != nil {
		t.Error(err)
	}

	str := string(jsondata)
	if str != `{"id":0,"title":"","rank":5,"type":"1+9","hash":"hash"}` {
		t.Error("Song json encode err: " + str)
	}
}
