package tests

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/jordan-bonecutter/goption"
	"github.com/olachat/gola/v2/golalib/testdata/wallet"
)

func TestWallet(t *testing.T) {
	ctx := context.Background()
	ctxCancel, cancel := context.WithCancelCause(ctx)
	wallet1 := wallet.NewWithPK(wallet.PK{
		UserId:     2,
		WalletType: 5,
	})

	var errCancelled = errors.New("context cancelled")
	mustNil(wallet1.InsertCtx(ctxCancel))
	cancel(errCancelled)
	_, err := wallet.FindCtx(ctxCancel, "where user_id = ?", 2)
	mustBe(err, errCancelled)

	_, err = wallet.FindOneCtx(ctxCancel, "where user_id = ?", 2)
	mustBe(err, errCancelled)

	wallet1.SetWalletName(goption.Some("goldbean"))
	errDup := wallet1.InsertCtx(ctx)
	if errDup == nil {
		t.Fatal("expect duplicated error")
	}
	mustBeSqlErrorWithNumber(errDup, 1062)
	ok, err := wallet1.UpdateCtx(ctx)
	mustNil(err)
	if !ok {
		t.Fatal("must be ok")
	}

	wallet1.SetWalletName(goption.Some("diamond"))
	mustBe(wallet1.InsertCtx(ctxCancel), errCancelled)

	wallet1.SetMoney(3000)
	ok, err = wallet1.UpdateCtx(ctx)
	mustNil(err)
	if !ok {
		t.Fatal("update should be ok")
	}

	wallet1.SetWalletName(goption.Some("seashell"))
	ok, err = wallet1.UpdateCtx(ctxCancel)
	mustBe(err, errCancelled)
	if ok {
		t.Fatalf("update should fail")
	}

	wallet2 := wallet.NewWithPK(wallet.PK{
		UserId:     2,
		WalletType: 4,
	})
	wallet2.SetMoney(5000)
	wallet2.SetWalletName(goption.Some("seashell"))
	mustBe(wallet2.InsertCtx(ctxCancel), errCancelled)
	mustNil(wallet2.InsertCtx(ctx))

	wallets, err := wallet.FindFieldsCtx[struct {
		wallet.Money      `json:"money"`
		wallet.WalletName `json:"wallet_name"`
	}](ctx, "where user_id = 2")
	mustNil(err)
	if len(wallets) != 2 {
		t.Fatal("expected 2 wallets")
	}
	b, _ := json.Marshal(wallets)
	fmt.Println(string(b))
}

func mustNil(err error) {
	if err != nil {
		panic("error: " + err.Error())
	}
}
func mustBe(err error, errExpected error) {
	if err == nil && !errors.Is(err, errExpected) {
		panic(fmt.Sprintf("expected %v, gotten: %v", errExpected, err))
	}
}

func mustBeSqlErrorWithNumber(err error, num int) {
	var errMySQL *mysql.MySQLError
	if err == nil || !errors.As(err, &errMySQL) {
		panic("is not mysql error")
	}
	if errMySQL.Number != uint16(num) {
		panic(fmt.Sprintf("expected %d, gotten: %d", num, errMySQL.Number))
	}
}
