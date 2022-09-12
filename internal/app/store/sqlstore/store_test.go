package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseUrl string
)

func TestMain(m *testing.M) {
	databaseUrl = os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		databaseUrl = "host=localhost port=6543 password=1138486 dbname=restapi_test sslmode=disable"
	}

	os.Exit(m.Run())
}
