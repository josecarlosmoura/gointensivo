package usecase

import (
	"github.com/josecarlosmoura/gointensivo/internal/order/entity"
	"github.com/josecarlosmoura/gointensivo/internal/order/infra/database"
	"github.com/josecarlosmoura/gointensivo/internal/order/interfaces"
)

type OrderInputDTO struct {
	ID    string
	Price float64
	Tax   float64
}

type OrderOutputDTO struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPriceUseCase struct {
	OerderRepository interfaces.OerderRepositoryInterface
}

func NewCalculateFinalPriceUseCase(orderRepository database.OrderRepository) *CalculateFinalPriceUseCase {
	return &CalculateFinalPriceUseCase{
		OerderRepository: &orderRepository,
	}
}

func (c *CalculateFinalPriceUseCase) Execute(input OrderInputDTO) (*OrderOutputDTO, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}

	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}

	err = c.OerderRepository.Save(order)
	if err != nil {
		return nil, err
	}

	return &OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
