package database

import (
	"context"
	"embed"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"
	"time"
)

// TestDatabase provides a Docker container to run tests
type TestDatabase struct {
	instance testcontainers.Container
}

func NewTestDatabase(t *testing.T) *TestDatabase {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15.1-alpine3.16",
		ExposedPorts: []string{"5432/tcp"},
		AutoRemove:   true,
		Env: map[string]string{
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_DB":       "postgres",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}
	postgres, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)
	return &TestDatabase{
		instance: postgres,
	}
}

func (db *TestDatabase) Port(t *testing.T) int {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	p, err := db.instance.MappedPort(ctx, "5432")
	require.NoError(t, err)
	return p.Int()
}

func (db *TestDatabase) Host(t *testing.T) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	h, err := db.instance.Host(ctx)
	require.NoError(t, err)
	return h
}

func (db *TestDatabase) DataSource(t *testing.T) string {
	return fmt.Sprintf("postgres:postgres@%s:%d/postgres?sslmode=disable", db.Host(t), db.Port(t))
}

func (db *TestDatabase) ConnectionString(t *testing.T) string {
	return fmt.Sprintf("postgres://%s", db.DataSource(t))
}

func (db *TestDatabase) Close(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	require.NoError(t, db.instance.Terminate(ctx))
}

//go:embed migrations
var migrations embed.FS

func SetupTestDatabase(t *testing.T, testDatabase *TestDatabase) {
	dataSourceName := testDatabase.ConnectionString(t)
	conn, err := NewConnection("postgres", dataSourceName)
	require.NoError(t, err)
	defer conn.Close()
	migrateDB(t, dataSourceName)
	insertTestData(t, conn)
}

func migrateDB(t *testing.T, dataSourceName string) {
	t.Helper()
	source, err := iofs.New(migrations, "migrations")
	require.NoError(t, err)
	defer source.Close()
	m, err := migrate.NewWithSourceInstance("iofs", source, dataSourceName)
	require.NoError(t, err)
	defer m.Close()
	require.NoError(t, m.Up())
}

func insertTestData(t *testing.T, connection Connection) {
	t.Helper()
	queries := []map[string][]any{
		{`insert into ccca.item(category, description, price, width, height, length, weight) values ($1, $2, $3, $4, $5, $6, $7);`: {"Instrumentos Musicais", "Guitarra", 1000, 100, 50, 15, 3}},
		{`insert into ccca.item(category, description, price, width, height, length, weight) values ($1, $2, $3, $4, $5, $6, $7);`: {"Instrumentos Musicais", "Amplificador", 5000, 50, 50, 50, 22}},
		{`insert into ccca.item(category, description, price, width, height, length, weight) values ($1, $2, $3, $4, $5, $6, $7);`: {"Acessórios", "Cabo", 30, 10, 10, 10, 1}},
		{`insert into ccca.coupon(code, percentage, expire_date) values ($1, $2, $3);`: {"VALE20", 20, "2022-10-10T10:00:00"}},
		{`insert into ccca.coupon(code, percentage, expire_date) values ($1, $2, $3);`: {"VALE20_EXPIRED", 20, "2020-10-10T10:00:00"}},
	}
	for _, query := range queries {
		for stmt, values := range query {
			require.NoErrorf(
				t, connection.Exec(stmt, values...),
				"Error on execute statement: %s with values %v", stmt, values,
			)
		}
	}
}
