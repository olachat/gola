package tests

import (
	"testing"

	"github.com/olachat/gola/golalib/testdata/profile"
)

func TestProfile(t *testing.T) {
	p := profile.NewProfileWithPK(8)
	p.SetNickName("tom")
	err := p.Insert()
	if err != nil {
		t.Error(err)
	}

	p2 := profile.FetchProfileByPK(8)
	if p2.GetNickName() != "tom" {
		t.Error("profile re-fetch error")
	}

	p3 := profile.NewProfileWithPK(8)
	p3.SetNickName("jerry")
	err = p3.Insert()

	/* TODO: profile table doesn't has primay key
	user_id is a unqiue index
	somehow coredb.ErrAvoidInsert is not returned
	should figure out why later
	*/
	if err != nil {
		t.Error(err)
	}

	p2 = profile.FetchProfileByPK(8)
	if p2.GetNickName() != "tom" {
		t.Error("profile re-fetch error")
	}
}
