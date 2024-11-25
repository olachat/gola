package tests

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/jordan-bonecutter/goption"
	"github.com/olachat/gola/v2/golalib/testdata/gifts"
)

func TestFetchGiftNull(t *testing.T) {
	gift := gifts.FetchByPK(1)
	assertNullGift(t, gift, 1)
}

func assertNullGift(t *testing.T, gift *gifts.Gift, pk uint) {
	if gift == nil {
		t.Fatalf("gift should not be nil")
	}
	if gift.GetId() != pk {
		t.Errorf("wrong pk")
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
		t.Fatal("gift should not be nil")
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

func TestInsertRetrieveUpdate(t *testing.T) {
	g1 := gifts.NewWithPK(21)
	err := g1.Insert()
	if err != nil {
		t.Fatalf("fail to insert. %v", err)
	}
	gOut := gifts.FetchByPK(21)
	assertNullGift(t, gOut, 21)

	g2 := gifts.New()
	g2.SetBranches(goption.Some([]gifts.GiftBranches{
		gifts.GiftBranchesOrchard,
		gifts.GiftBranchesChangi,
	}))
	g2.SetCreateTime(goption.Some[int64](9999999))
	g2.SetDescription(goption.Some("describe what this gift is about"))
	g2.SetDiscount(goption.Some(7.5))
	g2.SetGiftCount(goption.Some[int16](5))
	g2.SetGiftType(goption.Some(gifts.GiftGiftTypeEmpty))
	g2.SetIsFree(goption.Some(false))
	g2.SetManifest(goption.Some([]byte("manifest string")))
	g2.SetName(goption.Some("xmas gift"))
	g2.SetPrice(goption.Some(15.5))
	g2.SetRemark(goption.Some("selling out soon"))
	now := time.Now().UTC().Truncate(time.Second)
	g2.SetUpdateTime(goption.Some(now))
	err = g2.Insert()
	if err != nil {
		panic(err.Error())
	}
	g2out := gifts.FindOne("where id = ?", g2.GetId())
	if g2out == nil {
		t.Fatal("g2out should not be nil")
	}
	j2, _ := json.Marshal(g2)
	j2Out, _ := json.Marshal(g2out)
	if string(j2) != string(j2Out) {
		t.Fatalf("gift fetched is not as expected %s %s", string(j2), string(j2Out))
	}

	g2out.SetBranches(goption.Some([]gifts.GiftBranches{
		gifts.GiftBranchesVivo,
		gifts.GiftBranchesSentosa,
	}))
	g2out.SetCreateTime(goption.Some[int64](111))
	g2out.SetDescription(goption.Some("describe 2"))
	g2out.SetDiscount(goption.Some(4.67))
	g2out.SetGiftCount(goption.Some[int16](50))
	g2out.SetGiftType(goption.Some(gifts.GiftGiftTypeSovenir))
	g2out.SetIsFree(goption.Some(true))
	g2out.SetManifest(goption.Some([]byte("manifest 2")))
	g2out.SetName(goption.Some("gift 2"))
	g2out.SetPrice(goption.Some(65.555))
	g2out.SetRemark(goption.Some("remark 2"))
	now2 := time.Now().UTC().Truncate(time.Second)
	g2out.SetUpdateTime(goption.Some(now2))
	ok, err := g2out.Update()
	if err != nil {
		panic(err.Error())
	}
	if !ok {
		t.Fatal("update not done")
	}

	glist := gifts.Select().WherePriceEQ(65.555).All()
	if len(glist) != 1 {
		t.Errorf("should retrieve 1 record")
	}
	if glist[0].GetId() != g2.GetId() {
		t.Error("wrong id")
	}
	glist2 := gifts.Select().WherePriceEQ(65.55).All()
	if len(glist2) != 0 {
		t.Errorf("should retrieve 1 record")
	}
	glist3 := gifts.Select().WherePriceEQ(65.555).AndRemarkEQ("remark 2").All()
	if len(glist3) != 1 {
		t.Errorf("should retrieve 1 record")
	}
	if glist3[0].GetId() != g2.GetId() {
		t.Error("wrong id")
	}
	glist4 := gifts.SelectFields[struct {
		gifts.Id
		gifts.Discount
		gifts.CreateTime
	}]().OrderBy(gifts.CreateTimeAsc).All()
	if len(glist4) != 4 {
		t.Error("wrong count")
	}

	g22 := gifts.FindOne("where id = ?", g2out.GetId())
	if g22 == nil {
		t.Fatal("g22 should not be nil")
	}
	if !contains(g22.GetBranches().Unwrap(), gifts.GiftBranchesVivo) {
		t.Error("should contain vivo")
	}
	if !contains(g22.GetBranches().Unwrap(), gifts.GiftBranchesVivo) {
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
	if g22.GetGiftType().Unwrap() != gifts.GiftGiftTypeSovenir {
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
