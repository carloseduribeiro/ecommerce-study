package database

import (
	"github.com/ecommerce-study/internal/infra/database"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ItemRepositoryTestSuite struct {
	suite.Suite
	conn         database.Connection
	testDatabase *database.TestDatabase
}

func (i *ItemRepositoryTestSuite) BeforeTest(_, _ string) {
	t := i.T()
	i.testDatabase = database.NewTestDatabase(t)
	database.SetupTestDatabase(t, i.testDatabase)
	conn, err := database.NewConnection("postgres", i.testDatabase.ConnectionString(t))
	require.NoError(t, err)
	i.conn = conn
}

func (i *ItemRepositoryTestSuite) AfterTest(_, _ string) {
	i.conn.Close()
	i.testDatabase.Close(i.T())
}

func (i *ItemRepositoryTestSuite) TestGetById() {
	i.Run("should return an existing item", func() {
		t := i.T()
		itemRepository := NewItemRepository(i.conn)
		item, err := itemRepository.GetById(1)
		require.NoError(t, err)
		assert.Equal(t, item.Description(), "Guitarra")
	})

	i.Run("should return an error when trying get an non existing item", func() {
		t := i.T()
		itemRepository := NewItemRepository(i.conn)
		_, err := itemRepository.GetById(999)
		require.Error(t, err)
	})

	// TODO - desenvolver o resto dos testes de integração e de unidade
}

func TestItemRepository(t *testing.T) {
	suite.Run(t, new(ItemRepositoryTestSuite))
}
