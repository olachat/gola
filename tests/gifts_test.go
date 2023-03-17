package tests

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/jordan-bonecutter/goption"
	"github.com/olachat/gola/golalib/testdata/gifts"
)

func TestFetchGiftNull(t *testing.T) {
	gift := gifts.FetchByPK(1)
	if gift == nil {
		t.Errorf("gift should not be nil")
	}
	if gift.GetBranches().Ok() {
		t.Errorf("branches should be null")
	}
	if gift.GetCreateTime().Ok() {
		t.Errorf("create time should be null")
	}
	if gift.GetDescription().Ok() {
		t.Errorf("description should be null")
	}
	if gift.GetDiscount().Ok() {
		t.Errorf("discount should be null")
	}
	if gift.GetGiftCount().Ok() {
		t.Errorf("gift count should be null")
	}
	if gift.GetIsFree().Ok() {
		t.Errorf("is free should be null")
	}
	if gift.GetManifest().Ok() {
		t.Errorf("manifest should be null")
	}
	if gift.GetName().Ok() {
		t.Errorf("name should be null")
	}
	if gift.GetPrice().Ok() {
		t.Errorf("price should be null")
	}
	if gift.GetRemark().Ok() {
		t.Errorf("remark should be null")
	}
	if gift.GetGiftType().Ok() {
		t.Errorf("gift type should be null")
	}
	if gift.GetUpdateTime().Ok() {
		t.Errorf("update time should be null")
	}
}

func TestFetchGiftWithValue(t *testing.T) {
	gift := gifts.FetchByPK(2)
	if gift == nil {
		t.Errorf("gift should not be nil")
	}

	if !gift.GetBranches().Ok() {
		t.Errorf("branches should not be null")
	}
	branches := gift.GetBranches().Unwrap()
	if branches[0] != gifts.GiftBranchesVivo {
		t.Errorf("branch 0 should be vivo")
	}
	if branches[1] != gifts.GiftBranchesSentosa {
		t.Errorf("branch 1 should be sentosa")
	}

	if !gift.GetCreateTime().Ok() {
		t.Errorf("create time should not be null")
	}
	if gift.GetCreateTime().Unwrap() != 1678935576 {
		t.Errorf("create time should be 1678935576")
	}

	if !gift.GetDescription().Ok() {
		t.Errorf("description should not be null")
	}
	if gift.GetDescription().Unwrap() != "description text" {
		t.Error("description text should be 'description text'")
	}

	if !gift.GetDiscount().Ok() {
		t.Errorf("discount should not be null")
	}
	if gift.GetDiscount().Unwrap() != 7.5 {
		t.Errorf("discount should be 7.5")
	}

	if !gift.GetGiftCount().Ok() {
		t.Errorf("gift count should not be null")
	}
	if gift.GetGiftCount().Unwrap() != 3 {
		t.Errorf("gift count should be 3")
	}

	if !gift.GetIsFree().Ok() {
		t.Errorf("is free should not be null")
	}
	if !gift.GetIsFree().Unwrap() {
		t.Errorf("is free should be true")
	}

	if !gift.GetManifest().Ok() {
		t.Errorf("manifest should not be null")
	}
	if string(gift.GetManifest().Unwrap()) != "printable manifest" {
		t.Error("manifest should be 'printable manifest'")
	}

	if !gift.GetName().Ok() {
		t.Errorf("name should not be null")
	}
	if gift.GetName().Unwrap() != "name" {
		t.Errorf("name should be 'name'")
	}

	if !gift.GetPrice().Ok() {
		t.Errorf("price should not be null")
	}
	if gift.GetPrice().Unwrap() != 255.33 {
		t.Errorf("price should be 255.33")
	}

	if !gift.GetRemark().Ok() {
		t.Errorf("remark should not be null")
	}
	if gift.GetRemark().Unwrap() != "remark is long text" {
		t.Errorf("remark should be 'remark is long text'")
	}

	if !gift.GetGiftType().Ok() {
		t.Errorf("gift type should not be null")
	}
	if gift.GetGiftType().Unwrap() != gifts.GiftGiftTypeFreebie {
		t.Errorf("gift type should be freebie")
	}

	if !gift.GetUpdateTime().Ok() {
		t.Errorf("update time should not be null")
	}
	expectedTime, err := time.Parse("2006-01-02T15:04:05Z", "2019-01-01T00:00:01Z")
	if err != nil {
		panic(err)
	}
	if gift.GetUpdateTime().Unwrap() != expectedTime {
		t.Errorf("wrong time")
	}
}

// TODO: insert and retrieve and compare, update and retrieve and compare
func TestInsertRetrieveUpdate(t *testing.T) {
	g1 := gifts.NewWithPK(21)
	g1.SetBranches(goption.Some([]gifts.GiftBranches{
		gifts.GiftBranchesOrchard,
		gifts.GiftBranchesChangi,
	}))
	g1.SetCreateTime(goption.Some[int64](9999999))
	g1.SetDescription(goption.Some("describe what this gift is about"))
	g1.SetDiscount(goption.Some(7.5))
	g1.SetGiftCount(goption.Some[int16](5))
	g1.SetGiftType(goption.Some(gifts.GiftGiftTypeEmpty))
	g1.SetIsFree(goption.Some(false))
	g1.SetManifest(goption.Some([]byte("manifest string")))
	g1.SetName(goption.Some("xmas gift"))
	g1.SetPrice(goption.Some(15.5))
	g1.SetRemark(goption.Some("selling out soon"))
	now := time.Now().UTC().Truncate(time.Microsecond)
	g1.SetUpdateTime(goption.Some(now))
	err := g1.Insert()
	if err != nil {
		panic(err.Error())
	}
	g1out := gifts.FindOne("where gift_type = ?", gifts.GiftGiftTypeEmpty)
	j1, _ := json.Marshal(g1)
	j1Out, _ := json.Marshal(g1out)
	if string(j1) != string(j1Out) {
		t.Fatalf("gift fetched is not as expected")
	}
}
