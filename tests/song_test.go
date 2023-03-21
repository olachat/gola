package tests

import (
	"fmt"
	"testing"

	"encoding/json"

	"github.com/jordan-bonecutter/goption"
	"github.com/olachat/gola/golalib/testdata/song_user_favourites"
	"github.com/olachat/gola/golalib/testdata/songs"
)

func TestSong(t *testing.T) {
	s := songs.New()
	s.SetRank(5)
	s.SetType(goption.Some(songs.SongType1x2B9))
	s.SetHash("hash")
	s.SetRemark(goption.Some("hello remark"))
	s.SetManifest([]byte("manifest"))
	err := s.Insert()
	if err != nil {
		t.Error(err)
	}

	id := s.GetId()
	if id != 1000 {
		t.Error("Insert song error")
	}

	s = songs.FetchByPK(id)
	if s.GetRank() != 5 || s.GetHash() != "hash" || s.GetType().Unwrap() != songs.SongType1x2B9 {
		t.Error("Re-fetch song error")
	}

	if !s.GetRemark().Ok() {
		t.Error("song remarks should be present")
	}
	if s.GetRemark().String() != "hello remark" {
		t.Error("song remarks should be 'hello remark'")
	}

	s.SetType(goption.Some(songs.SongTypeEmpty))
	s.SetRemark(goption.None[string]())
	s.Update()
	s = songs.FetchByPK(id)
	if s.GetType().Unwrap() != songs.SongTypeEmpty {
		t.Error("Song update error")
	}
	if s.GetRemark().Ok() {
		t.Error("remark should be not present")
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
	if err == nil {
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

func TestFetchSong(t *testing.T) {
	s := songs.FetchByPK(999)
	if s == nil {
		t.Error("song should exist")
	}
	if s.GetId() != 999 {
		t.Error("song id should be 999")
	}
	if s.GetTitle() != "song1 2 3" {
		t.Error("song title wrong")
	}
	if string(s.GetManifest()) != "a" {
		t.Error("song manifest wrong")
	}
	if s.GetType().Unwrap() != songs.SongType101 {
		t.Error("song wrong type")
	}
	if s.GetRemark().Ok() {
		t.Error("song remark should not be present")
	}
	ss := songs.FetchFieldsByPK[struct {
		songs.Id
		songs.Remark
	}](999)
	if ss.GetId() != 999 {
		t.Errorf("song id should be 999")
	}
	if ss.GetRemark().Ok() {
		t.Error("song remark should not be present")
	}
	ss.SetRemark(goption.Some("remarks 123"))
	updated, err := songs.Update(ss)
	if !updated {
		t.Error("song should be updated")
	}
	if err != nil {
		t.Error("update song should succeed")
	}
	sss := songs.FetchFieldsByPK[struct {
		songs.Remark
	}](999)
	if !sss.GetRemark().Ok() {
		t.Error("song remark should be present")
	}
	if sss.GetRemark().String() != "remarks 123" {
		t.Error("song remark should 'remarks 123'")
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

	objs := song_user_favourites.Select().WhereUserIdEQ(pk.UserId).AndIsFavouriteEQ(false).All()
	if len(objs) != 1 && objs[0].GetSongId() != pk.SongId {
		t.Error("Song select failed")
	}

	objs = song_user_favourites.Select().WhereUserIdEQ(pk.UserId).AndIsFavouriteEQ(false).Limit(0, 1)
	if len(objs) != 1 && objs[0].GetSongId() != pk.SongId {
		t.Error("Song select with limit failed")
	}
}

func TestSongDelete(t *testing.T) {
	pk := song_user_favourites.PK{
		UserId: 3,
		SongId: 99,
	}

	err := song_user_favourites.DeleteByPK(pk)
	if err != nil {
		t.Error("SongUserFavourite delete fail")
	}
	obj := song_user_favourites.FetchByPK(pk)
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
	s.SetManifest([]byte(""))
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
	id := uint(1100)
	s := songs.NewWithPK(id)
	s.SetHash("100")
	s.SetManifest([]byte(""))
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
	s.SetManifest([]byte(""))
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
	s.SetManifest([]byte("abc"))
	err = s.Insert()
	if err != nil {
		t.Error(err)
	}
	s = songs.FetchByPK(id)
	if string(s.GetManifest()) != "abc" {
		t.Error("Insert with []byte error")
	}

	if s == nil {
		t.Error("Reinsert with pk failed")
	}
	s = songs.New()
	s.SetHash("hashhash")
	s.SetManifest([]byte(""))
	err = s.Insert()
	if err != nil {
		t.Error(err)
	}
	if s.GetId() != 1100+1+1 {
		t.Error("Insert without pk after pk failed")
	}
}

func TestSongJSONEncode(t *testing.T) {
	s := songs.New()
	s.SetRank(5)
	s.SetType(goption.Some(songs.SongType1x2B9))
	s.SetHash("hash")
	s.SetManifest([]byte("abc"))

	jsondata, err := json.Marshal(&s)
	if err != nil {
		t.Error(err)
	}

	str := string(jsondata)
	if str != `{"id":0,"title":"","rank":5,"type":"1+9","hash":"hash","remark":null,"manifest":"YWJj"}` {
		t.Error("Song json encode err: " + str)
	}

	var s2 *songs.Song
	json.Unmarshal(jsondata, &s2)
	jsondata, err = json.Marshal(&s2)
	if err != nil {
		t.Error(err)
	}

	if str != string(jsondata) {
		t.Error("Song json re-encode err: " + string(jsondata))
	}
}
