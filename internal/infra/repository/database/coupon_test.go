package database

import (
	"github.com/ecommerce-study/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CouponRepositoryTestSuite struct {
	suite.Suite
	tearDownConnection func() error
	tearDownTestDB     func(t *testing.T)
	repository         *CouponRepository
}

func (c *CouponRepositoryTestSuite) BeforeTest(_, _ string) {
	t := c.T()
	testDB := database.NewTestDatabase(t)
	c.tearDownTestDB = testDB.Close
	database.SetupTestDatabase(t, testDB)
	conn, err := database.NewConnection("postgres", testDB.ConnectionString(t))
	c.tearDownConnection = conn.Close
	require.NoError(t, err)
	c.repository = NewCouponRepository(conn)
}

func (c *CouponRepositoryTestSuite) AfterTest(_, _ string) {
	c.tearDownConnection()
	c.tearDownTestDB(c.T())
}

func (c *CouponRepositoryTestSuite) TestGetByCode() {
	c.Run("should return an existing coupon", func() {
		t := c.T()
		// Given
		code := "VALE20"
		// When
		coupon, err := c.repository.GetByCode(code)
		require.NoError(t, err)
		assert.NotNil(t, coupon)
	})

	c.Run("should return an error when trying get an non existing coupon", func() {
		t := c.T()
		// Given
		code := "CAFE"
		// When
		coupon, err := c.repository.GetByCode(code)
		require.Nil(t, coupon)
		assert.Error(t, err)
	})
}

func TestCouponRepository(t *testing.T) {
	suite.Run(t, new(CouponRepositoryTestSuite))
}
