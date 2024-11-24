package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/olachat/gola/v2/golalib/testdata/gifts_nn"
)

func TestFetchGiftNotNull(t *testing.T) {
	gift := gifts_nn.FetchByPK(1)
	if gift == nil {
		t.Fatal("gift nn should not be nil")
	}
	if gift.GetId() != 1 {
		t.Error("wrong id")
	}
	branches := gift.GetBranches()
	if !contains(branches, gifts_nn.GiftsNnBranchesVivo) {
		t.Errorf("branches should contain vivo")
	}
	if !contains(branches, gifts_nn.GiftsNnBranchesSentosa) {
		t.Errorf("branches should contain sentosa")
	}

	if gift.GetCreateTime() != 1678935576 {
		t.Errorf("create time wrong")
	}

	if gift.GetDescription() != "description text" {
		t.Error("wrong description")
	}

	if !isFloatSimilar(gift.GetDiscount(), 7.5) {
		t.Error("wrong discount")
	}

	if gift.GetGiftCount() != 3 {
		t.Errorf("wrong gift count")
	}

	if gift.GetGiftType() != gifts_nn.GiftsNnGiftTypeFreebie {
		t.Error("wrong gift type")
	}

	if gift.GetIsFree() != true {
		t.Error("wrong is free")
	}

	if string(gift.GetManifest()) != "printable manifest" {
		t.Error("wrong manifest")
	}

	if gift.GetName() != "name" {
		t.Error("wrong name")
	}

	if gift.GetPrice() != 255.33 {
		t.Error("wrong price")
	}

	if gift.GetRemark() != "remark is long text" {
		t.Error("wrong remark")
	}

	ti, err := time.Parse("2006-01-02 15:04:05.999999", "2019-01-01 00:00:01.000000")
	if err != nil {
		panic(err)
	}
	if !gift.GetUpdateTime().Equal(ti) {
		t.Error("wrong update time")
	}
}

func TestGiftNotNullInsertRetrieveUpdate(t *testing.T) {
	gift := gifts_nn.NewWithPK(2)
	gift.SetManifest([]byte("hello world"))
	err := gift.Insert()
	if err != nil {
		panic(err.Error())
	}

	gOut := gifts_nn.FetchByPK(2)
	if string(gOut.GetManifest()) != "hello world" {
		t.Error("wrong manifest")
	}

	if len(gOut.GetBranches()) != 0 {
		t.Error("wrong branches")
	}

	if gOut.GetCreateTime() != 0 {
		t.Error("wrong create time")
	}

	if gOut.GetId() != 2 {
		t.Error("wrong id")
	}

	if gOut.GetDescription() != "" {
		t.Error("wrong description")
	}
	if gOut.GetDiscount() != 0.0 {
		t.Error("wrong discount")
	}

	if gOut.GetGiftCount() != 0 {
		t.Error("wrong gift count")
	}

	if gOut.GetGiftType() != gifts_nn.GiftsNnGiftTypeEmpty {
		t.Error("wrong gift type")
	}

	if gOut.GetIsFree() != false {
		t.Error("wrong is free")
	}

	if gOut.GetName() != "" {
		t.Error("wrong name")
	}

	if gOut.GetPrice() != 0.0 {
		t.Error("wrong price")
	}

	if gOut.GetRemark() != "" {
		t.Error("wrong remark")
	}

	fmt.Println(gOut.GetUpdateTime())

	gOut.SetBranches([]gifts_nn.GiftsNnBranches{
		gifts_nn.GiftsNnBranchesVivo,
		gifts_nn.GiftsNnBranchesSentosa,
	})
	gOut.SetCreateTime(111)
	gOut.SetDescription("describe 2")
	gOut.SetDiscount(4.67)
	gOut.SetGiftCount(50)
	gOut.SetGiftType(gifts_nn.GiftsNnGiftTypeSovenir)
	gOut.SetIsFree(true)
	gOut.SetManifest([]byte("manifest 2"))
	gOut.SetName("gift 2")
	gOut.SetPrice(65.555)
	gOut.SetRemark("remark 2")
	now2 := time.Now().UTC().Truncate(time.Second)
	gOut.SetUpdateTime(now2)
	ok, err := gOut.Update()
	if err != nil {
		panic(err.Error())
	}
	if !ok {
		t.Fatal("update not done")
	}

	g22 := gifts_nn.FindOne("where id = ?", gOut.GetId())
	if g22 == nil {
		t.Fatal("g22 should not be nil")
	}
	if g22.GetId() != 2 {
		t.Error("wrong id")
	}
	if !contains(g22.GetBranches(), gifts_nn.GiftsNnBranchesVivo) {
		t.Error("should contain vivo")
	}
	if !contains(g22.GetBranches(), gifts_nn.GiftsNnBranchesVivo) {
		t.Error("should contain vivo")
	}
	if g22.GetCreateTime() != 111 {
		t.Error("wrong create time")
	}
	if g22.GetDescription() != "describe 2" {
		t.Error("wrong desc")
	}
	if !isFloatSimilar(g22.GetDiscount(), 4.67) {
		t.Error("wrong discount")
	}
	if g22.GetGiftCount() != 50 {
		t.Error("wrong gift_count")
	}
	if g22.GetGiftType() != gifts_nn.GiftsNnGiftTypeSovenir {
		t.Error("wrong gift type")
	}
	if g22.GetIsFree() != true {
		t.Error("wrong is free")
	}
	if string(g22.GetManifest()) != "manifest 2" {
		t.Error("wrong manifest")
	}
	if g22.GetName() != "gift 2" {
		t.Error("wrong name")
	}
	if !isFloatSimilar(g22.GetPrice(), 65.555) {
		t.Error("wrong price")
	}
	if g22.GetRemark() != "remark 2" {
		t.Error("wrong remark")
	}
	if !g22.GetUpdateTime().Equal(now2) {
		t.Errorf("wrong update time. expected %s, gotten %s", now2, g22.GetUpdateTime())
	}
}
