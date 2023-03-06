package tests

import (
	"testing"

	"github.com/olachat/gola/golalib/testdata/gifts"
)

func TestGift(t *testing.T) {
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
