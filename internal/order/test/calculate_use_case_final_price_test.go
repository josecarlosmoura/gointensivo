package test

import (
	"database/sql"
	"testing"

	"github.com/josecarlosmoura/gointensivo/internal/order/entity"
	"github.com/josecarlosmoura/gointensivo/internal/order/infra/database"
	"github.com/josecarlosmoura/gointensivo/internal/order/usecase"
	"github.com/stretchr/testify/suite"
)

type CalculateUseCaseFinalPriceTestSuite struct {
	suite.Suite
	OderRepository database.OrderRepository
	Db             *sql.DB
}

// Always run before rotating testicles
func (suite *CalculateUseCaseFinalPriceTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	// crete table orders
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
	suite.OderRepository = *database.NewOrderRepository(db)
}

// Runs every time after running the tests
func (suite *CalculateUseCaseFinalPriceTestSuite) TearDownTest() {
	suite.Db.Exec("DELETE FROM orders")
	suite.Db.Close()
}

// Esta função com o nome "TestSuite" sem o "s" no final está dando problema de duplicidade.
// Estudar o motivo e como posso melhorar para uma abistração
func TestSuites(t *testing.T) {
	suite.Run(t, new(CalculateUseCaseFinalPriceTestSuite))
}

func (suite *CalculateUseCaseFinalPriceTestSuite) TestCalculateFinalPrice() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	order.CalculateFinalPrice()

	calculateFinalPriceInput := usecase.OrderInputDTO{
		ID:    order.ID,
		Price: order.Price,
		Tax:   order.Tax,
	}

	calculateFinalPriceUseCase := usecase.NewCalculateFinalPriceUseCase(suite.OderRepository)
	output, err := calculateFinalPriceUseCase.Execute(calculateFinalPriceInput)
	suite.NoError(err)

	suite.Equal(order.ID, output.ID)
	suite.Equal(order.Price, output.Price)
	suite.Equal(order.Tax, output.Tax)
	suite.Equal(order.FinalPrice, output.FinalPrice)
}
