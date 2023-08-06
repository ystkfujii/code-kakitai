package query_service

import (
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/db_test"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
	"testing"

	"gopkg.in/testfixtures.v2"
)

var (
	query    *dbgen.Queries
	fixtures *testfixtures.Context
	err      error
)

func TestMain(m *testing.M) {
	// DBの立ち上げ
	resource, pool := db_test.CreateContainer()
	defer db_test.CloseContainer(resource, pool)

	// DBへ接続する
	db := db_test.ConnectDB(resource, pool)
	defer db.Close()

	// テスト用DBをセットアップ
	db_test.SetupTestDB()

	// テストデータの準備
	fixturePath := "../../fixtures"
	fixtures, err = testfixtures.NewFolder(db, &testfixtures.MySQL{}, fixturePath)
	if err != nil {
		panic(err)
	}

	query = dbgen.New(db)

	// テスト実行
	m.Run()
}

func prepare(t *testing.T) {
	t.Helper()
	if err := fixtures.Load(); err != nil {
		t.Fatal(err)
	}
}
