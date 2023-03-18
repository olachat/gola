package tests

import (
	"testing"

	"github.com/olachat/gola/golalib/testdata/gifts_with_default"
)

func TestFetchGiftWithDefault(t *testing.T) {
	gift := gifts_with_default.FetchByPK(1)
	if gift == nil {
		t.Errorf("gift should not be nil")
	}
	gift2 := gifts_with_default.NewWithPK(2)
	err := gift2.Insert()
	if err != nil {
		t.Fatalf("insert failed: %s", err.Error())
	}
}
