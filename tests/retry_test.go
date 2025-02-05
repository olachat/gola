package tests

import (
	"context"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/olachat/gola/v2/coredb"
)

func TestExecWithRetry_Success(t *testing.T) {
	ctx := context.Background()

	_, err := coredb.ExecWithRetry(ctx, testDBName, "INSERT INTO test_table (name, email) VALUES (?, ?)", coredb.DefaultRetryConfig, "test", "test@example.com")
	if err != nil {
		t.Fatalf("Expected success, but got error: %v", err)
	}
}

func TestExecWithRetry_Fail(t *testing.T) {
	ctx := context.Background()
	now := time.Now()
	_, err := coredb.ExecWithRetry(ctx, testDBName, "INSERT INTO no_such_table (name, email) VALUES (?, ?)", coredb.DefaultRetryConfig, "test", "test@example.com")
	if err == nil {
		t.Fatalf("Expected error, but got success")
	}
	elapsed := time.Since(now)
	if elapsed < (200+400+800+1600+3200)*time.Millisecond {
		t.Fatalf("Expected retry to take at least 100ms, but took: %v", elapsed)
	}
	t.Logf("retry took: %v", elapsed)
}
