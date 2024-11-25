package tests

import (
	"testing"

	"github.com/olachat/gola/v2/golalib/testdata/account"
	"github.com/olachat/gola/v2/golalib/testdata/profile"
)

func TestProfile(t *testing.T) {
	p := profile.NewWithPK(8)
	p.SetNickName("tom")
	err := p.Insert()
	if err != nil {
		t.Error(err)
	}

	p2 := profile.FetchByPK(8)
	if p2.GetNickName() != "tom" {
		t.Error("profile re-fetch error")
	}

	p3 := profile.NewWithPK(8)
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

	p2 = profile.FetchByPK(8)
	if p2.GetNickName() != "tom" {
		t.Error("profile re-fetch error")
	}
}

func TestAccount(t *testing.T) {
	pk := account.PK{
		UserId:      8,
		CountryCode: 4,
	}
	a := account.NewWithPK(pk)
	a.SetMoney(47)
	err := a.Insert()
	if err != nil {
		t.Error(err)
	}

	a2 := account.FetchByPK(pk)
	if a2.GetMoney() != a.GetMoney() || a2.GetType() != a.GetType() || a2.GetUserId() != a.GetUserId() {
		t.Error("account re-fetch error")
	}
	a.SetMoney(84)
	a.SetType(account.AccountTypeVip)
	ok, err := a.Update()
	if err != nil || !ok {
		t.Error("account update error")
	}

	a2 = account.FetchByPK(pk)
	if a2.GetMoney() != a.GetMoney() || a2.GetType() != a.GetType() || a2.GetUserId() != a.GetUserId() {
		t.Error("account re-fetch error")
	}
}
