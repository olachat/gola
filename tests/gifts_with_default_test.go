package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/jordan-bonecutter/goption"
	"github.com/olachat/gola/golalib/testdata/gifts_with_default"
)

func TestFetchGiftWithDefault(t *testing.T) {
	gift := gifts_with_default.FetchByPK(1)
	if gift == nil {
		t.Fatalf("gift should not be nil")
	}
	if gift.GetId() != 1 {
		t.Error("wrong id")
	}
	assertDefaultGiftWithDefault(t, gift)

}

func assertDefaultGiftWithDefault(t *testing.T, gift *gifts_with_default.GiftsWithDefault) {
	branches := gift.GetBranches().Unwrap()
	if !contains(branches, gifts_with_default.GiftsWithDefaultBranchesChangi) {
		t.Errorf("branches should contain changi")
	}
	if !contains(branches, gifts_with_default.GiftsWithDefaultBranchesSentosa) {
		t.Errorf("branches should contain sentosa")
	}

	if gift.GetCreateTime().Unwrap() != 999 {
		t.Errorf("create time should be 999")
	}

	if gift.GetDescription().Unwrap() != "default gift" {
		t.Error("wrong description")
	}

	if !isFloatSimilar(gift.GetDiscount().Unwrap(), 0.1) {
		t.Error("wrong discount")
	}

	if gift.GetGiftCount().Unwrap() != 1 {
		t.Errorf("wrong gift count")
	}

	if gift.GetGiftType().Unwrap() != gifts_with_default.GiftsWithDefaultGiftTypeMembership {
		t.Error("wrong gift type")
	}

	if gift.GetIsFree().Unwrap() != true {
		t.Error("wrong is free")
	}

	if string(gift.GetManifest().Unwrap()) != "manifest data" {
		t.Error("wrong manifest")
	}

	if gift.GetName().Unwrap() != "gift for you" {
		t.Error("wrong name")
	}

	if gift.GetPrice().Unwrap() != 5.0 {
		t.Error("wrong price")
	}

	if gift.GetRemark().Unwrap() != "hope you like it" {
		t.Error("wrong remark")
	}

	ti, err := time.Parse("2006-01-02 15:04:05.999999", "2023-01-19 03:14:07.999999")
	if err != nil {
		panic(err)
	}
	if !gift.GetUpdateTime().Unwrap().Equal(ti) {
		t.Error("wrong update time")
	}
	fmt.Println(gift.GetUpdateTime2().Unwrap())
}

func TestGiftWithDefaultRetrieve(t *testing.T) {
	gift := gifts_with_default.FetchByPK(2)
	if gift == nil {
		t.Fatalf("gift should not be nil")
	}
	if gift.GetId() != 2 {
		t.Error("wrong id")
	}
	branches := gift.GetBranches().Unwrap()
	if !contains(branches, gifts_with_default.GiftsWithDefaultBranchesVivo) {
		t.Errorf("branches should contain vivo")
	}
	if !contains(branches, gifts_with_default.GiftsWithDefaultBranchesSentosa) {
		t.Errorf("branches should contain sentosa")
	}

	if gift.GetCreateTime().Unwrap() != 1678935576 {
		t.Errorf("create time wrong")
	}

	if gift.GetDescription().Unwrap() != "description text" {
		t.Error("wrong description")
	}

	if !isFloatSimilar(gift.GetDiscount().Unwrap(), 7.5) {
		t.Error("wrong discount")
	}

	if gift.GetGiftCount().Unwrap() != 3 {
		t.Errorf("wrong gift count")
	}

	if gift.GetGiftType().Unwrap() != gifts_with_default.GiftsWithDefaultGiftTypeFreebie {
		t.Error("wrong gift type")
	}

	if gift.GetIsFree().Unwrap() != true {
		t.Error("wrong is free")
	}

	if string(gift.GetManifest().Unwrap()) != "printable manifest" {
		t.Error("wrong manifest")
	}

	if gift.GetName().Unwrap() != "name" {
		t.Error("wrong name")
	}

	if gift.GetPrice().Unwrap() != 255.33 {
		t.Error("wrong price")
	}

	if gift.GetRemark().Unwrap() != "remark is long text" {
		t.Error("wrong remark")
	}

	ti, err := time.Parse("2006-01-02 15:04:05.999999", "2019-01-01 00:00:01.000000")
	if err != nil {
		panic(err)
	}
	if !gift.GetUpdateTime().Unwrap().Equal(ti) {
		t.Error("wrong update time")
	}
	fmt.Println(gift.GetUpdateTime2().Unwrap())
}

func TestGiftWithDefaultInsertRetrieveUpdate(t *testing.T) {
	g1 := gifts_with_default.NewWithPK(3)
	err := g1.Insert()
	if err != nil {
		t.Fatalf("fail to insert. %v", err)
	}
	gOut := gifts_with_default.FetchByPK(3)
	if gOut == nil {
		t.Fatal("should not be nil")
	}
	if gOut.GetId() != 3 {
		t.Error("wrong id")
	}
	assertDefaultGiftWithDefault(t, gOut)

	gOut.SetBranches(goption.Some([]gifts_with_default.GiftsWithDefaultBranches{
		gifts_with_default.GiftsWithDefaultBranchesVivo,
		gifts_with_default.GiftsWithDefaultBranchesSentosa,
	}))
	gOut.SetCreateTime(goption.Some[int64](111))
	gOut.SetDescription(goption.Some("describe 2"))
	gOut.SetDiscount(goption.Some(4.67))
	gOut.SetGiftCount(goption.Some[int16](50))
	gOut.SetGiftType(goption.Some(gifts_with_default.GiftsWithDefaultGiftTypeSovenir))
	gOut.SetIsFree(goption.Some(true))
	gOut.SetManifest(goption.Some([]byte("manifest 2")))
	gOut.SetName(goption.Some("gift 2"))
	gOut.SetPrice(goption.Some(65.555))
	gOut.SetRemark(goption.Some("remark 2"))
	now2 := time.Now().UTC().Truncate(time.Microsecond)
	gOut.SetUpdateTime(goption.Some(now2))
	ok, err := gOut.Update()
	if err != nil {
		panic(err.Error())
	}
	if !ok {
		t.Fatal("update not done")
	}

	g22 := gifts_with_default.FindOne("where id = ?", gOut.GetId())
	if g22 == nil {
		t.Fatal("g22 should not be nil")
	}
	if g22.GetId() != 3 {
		t.Error("wrong id")
	}
	if !contains(g22.GetBranches().Unwrap(), gifts_with_default.GiftsWithDefaultBranchesVivo) {
		t.Error("should contain vivo")
	}
	if !contains(g22.GetBranches().Unwrap(), gifts_with_default.GiftsWithDefaultBranchesVivo) {
		t.Error("should contain vivo")
	}
	if g22.GetCreateTime().Unwrap() != 111 {
		t.Error("wrong create time")
	}
	if g22.GetDescription().Unwrap() != "describe 2" {
		t.Error("wrong desc")
	}
	if !isFloatSimilar(g22.GetDiscount().Unwrap(), 4.67) {
		t.Error("wrong discount")
	}
	if g22.GetGiftCount().Unwrap() != 50 {
		t.Error("wrong gift_count")
	}
	if g22.GetGiftType().Unwrap() != gifts_with_default.GiftsWithDefaultGiftTypeSovenir {
		t.Error("wrong gift type")
	}
	if g22.GetIsFree().Unwrap() != true {
		t.Error("wrong is free")
	}
	if string(g22.GetManifest().Unwrap()) != "manifest 2" {
		t.Error("wrong manifest")
	}
	if g22.GetName().Unwrap() != "gift 2" {
		t.Error("wrong name")
	}
	if !isFloatSimilar(g22.GetPrice().Unwrap(), 65.555) {
		t.Error("wrong price")
	}
	if g22.GetRemark().Unwrap() != "remark 2" {
		t.Error("wrong remark")
	}
	if g22.GetUpdateTime().Unwrap() != now2 {
		t.Error("wrong update time")
	}
}
