package databasetest

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/require"
)

var (
	dbConnSet    = map[string]*sql.DB{}
	unusedDBConn = map[string]struct{}{}
	mutex        sync.Mutex
)

func Initialize(t *testing.T, schema string) (*sql.DB, func()) {
	t.Helper()
	mutex.Lock()
	defer mutex.Unlock()
	for dbname := range unusedDBConn {
		delete(unusedDBConn, dbname)
		db, ok := dbConnSet[dbname]
		if !ok {
			continue
		}
		cleanup(t, db)
		return db, func() {
			unusedDBConn[dbname] = struct{}{}
		}
	}
	dsn := os.Getenv("DATABASE_DSN")
	pgConfig, err := pgx.ParseConfig(dsn)
	require.NoError(t, err)
	dbname := fmt.Sprintf("%s_%03d", pgConfig.Database, len(dbConnSet)+1)

	pgConfig.Database = dbname
	connStr := stdlib.RegisterConnConfig(pgConfig)
	db, err := sql.Open("pgx", connStr)
	require.NoError(t, err)

	if err := db.Ping(); err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) && pgError.Code == "3D000" {
			origDB, err := sql.Open("pgx", dsn)
			require.NoError(t, err)
			func() {
				defer origDB.Close()
				require.NoError(t, origDB.Ping())
				_, err := origDB.ExecContext(t.Context(), fmt.Sprintf("CREATE DATABASE %s", dbname))
				require.NoError(t, err)
			}()
			require.NoError(t, db.Ping())
		} else {
			t.Fatal(err)
		}
	}
	loadScehma(t, db, schema)

	dbConnSet[dbname] = db
	return db, func() {
		unusedDBConn[dbname] = struct{}{}
	}
}

func cleanup(t *testing.T, db *sql.DB) {
	t.Helper()
	rows, err := db.QueryContext(
		t.Context(),
		"SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname = $1",
		"public",
	)
	require.NoError(t, err)
	tableNames := []string{}
	func() {
		defer rows.Close()
		for rows.Next() {
			var tableName string
			require.NoError(t, rows.Scan(&tableName))
			tableNames = append(tableNames, tableName)
		}
	}()
	for _, tableName := range tableNames {
		_, err := db.ExecContext(t.Context(), "TRUNCATE TABLE "+tableName)
		require.NoError(t, err, "table: %s", tableName)
	}
}

func loadScehma(t *testing.T, db *sql.DB, schema string) {
	t.Helper()
	bs, err := os.ReadFile(schema)
	require.NoError(t, err)
	seq := bytes.SplitSeq(bs, []byte(";\n"))
	for stmt := range seq {
		_, err := db.ExecContext(t.Context(), string(stmt))
		require.NoError(t, err)
	}
}
