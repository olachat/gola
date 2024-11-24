package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/olachat/gola/v2/golalib/testdata/gifts_nn_with_default"
)

func TestFetchGiftNotNullWithDefault(t *testing.T) {
	gift := gifts_nn_with_default.FetchByPK(1)
	if gift == nil {
		t.Fatal("gift nn should not be nil")
	}
	if gift.GetId() != 1 {
		t.Error("wrong id")
	}

	assertDefaultGiftNnWithDefault(t, gift, "gift1")
}

func assertDefaultGiftNnWithDefault(t *testing.T, gift *gifts_nn_with_default.GiftsNnWithDefault, desc string) {
	branches := gift.GetBranches()
	if !contains(branches, gifts_nn_with_default.GiftsNnWithDefaultBranchesChangi) {
		t.Errorf("branches should contain changi")
	}
	if !contains(branches, gifts_nn_with_default.GiftsNnWithDefaultBranchesSentosa) {
		t.Errorf("branches should contain sentosa")
	}

	if gift.GetCreateTime() != 999 {
		t.Errorf("create time should be 999")
	}

	if gift.GetDescription() != desc {
		t.Error("wrong description")
	}

	if !isFloatSimilar(gift.GetDiscount(), 0.1) {
		t.Error("wrong discount")
	}

	if gift.GetGiftCount() != 1 {
		t.Errorf("wrong gift count")
	}

	if gift.GetGiftType() != gifts_nn_with_default.GiftsNnWithDefaultGiftTypeMembership {
		t.Error("wrong gift type")
	}

	if gift.GetIsFree() != true {
		t.Error("wrong is free")
	}

	if string(gift.GetManifest()) != "manifest data" {
		t.Error("wrong manifest")
	}

	if gift.GetName() != "gift for you" {
		t.Error("wrong name")
	}

	if gift.GetPrice() != 5.0 {
		t.Error("wrong price")
	}

	if gift.GetRemark() != "hope you like it" {
		t.Error("wrong remark")
	}

	ti, err := time.Parse("2006-01-02 15:04:05.0", "2023-01-19 03:14:07.0")
	if err != nil {
		panic(err)
	}
	if !gift.GetUpdateTime().Equal(ti) {
		t.Error("wrong update time", ti, gift.GetUpdateTime())
	}
	fmt.Println(gift.GetUpdateTime2())
}

func TestGiftNotNullWithDefaultInsertRetrieveUpdate(t *testing.T) {
	gift := gifts_nn_with_default.NewWithPK(3)
	if gift == nil {
		t.Fatalf("gift should not be nil")
	}
	if gift.GetId() != 3 {
		t.Error("wrong id")
	}

	assertDefaultGiftNnWithDefault(t, gift, "")

	err := gift.Insert()
	if err != nil {
		panic(err.Error())
	}

	gift = gifts_nn_with_default.FindOne("where id = 3")
	if gift == nil {
		t.Fatal("gift should not be nil")
	}
	if gift.GetId() != 3 {
		t.Error("wrong id")
	}
	assertDefaultGiftNnWithDefault(t, gift, "")

	gift.SetBranches([]gifts_nn_with_default.GiftsNnWithDefaultBranches{
		gifts_nn_with_default.GiftsNnWithDefaultBranchesVivo,
		gifts_nn_with_default.GiftsNnWithDefaultBranchesSentosa,
	})
	gift.SetCreateTime(111)
	gift.SetDescription("describe 2")
	gift.SetDiscount(4.67)
	gift.SetGiftCount(50)
	gift.SetGiftType(gifts_nn_with_default.GiftsNnWithDefaultGiftTypeSovenir)
	gift.SetIsFree(true)
	gift.SetManifest([]byte("manifest 2"))
	gift.SetName("gift 2")
	gift.SetPrice(65.555)
	gift.SetRemark("remark 2")
	now2 := time.Now().UTC().Truncate(time.Second)
	gift.SetUpdateTime(now2)
	ok, err := gift.Update()
	if err != nil {
		panic(err.Error())
	}
	if !ok {
		t.Fatal("update not done")
	}

	gift = gifts_nn_with_default.FindOne("where id = 3")
	if gift == nil {
		t.Fatal("gift should not be nil")
	}
	if gift.GetId() != 3 {
		t.Error("wrong id")
	}
	branches := gift.GetBranches()
	if !contains(branches, gifts_nn_with_default.GiftsNnWithDefaultBranchesVivo) {
		t.Errorf("branches should contain vivo")
	}
	if !contains(branches, gifts_nn_with_default.GiftsNnWithDefaultBranchesSentosa) {
		t.Errorf("branches should contain sentosa")
	}

	if gift.GetCreateTime() != 111 {
		t.Errorf("create time wrong")
	}

	if gift.GetDescription() != "describe 2" {
		t.Error("wrong description")
	}

	if !isFloatSimilar(gift.GetDiscount(), 4.67) {
		t.Error("wrong discount")
	}

	if gift.GetGiftCount() != 50 {
		t.Errorf("wrong gift count")
	}

	if gift.GetGiftType() != gifts_nn_with_default.GiftsNnWithDefaultGiftTypeSovenir {
		t.Error("wrong gift type")
	}

	if gift.GetIsFree() != true {
		t.Error("wrong is free")
	}

	if string(gift.GetManifest()) != "manifest 2" {
		t.Error("wrong manifest")
	}

	if gift.GetName() != "gift 2" {
		t.Error("wrong name")
	}

	if gift.GetPrice() != 65.555 {
		t.Error("wrong price")
	}

	if gift.GetRemark() != "remark 2" {
		t.Error("wrong remark")
	}

	if !gift.GetUpdateTime().Equal(now2) {
		t.Error("wrong update time")
	}
	fmt.Println(gift.GetUpdateTime2())
}
