package usecase

import "github.com/josecarlosmoura/gointensivo/internal/order/interfaces"

type GetTotalOutputDTO struct {
	Total int
}

type GetTotalUseCase struct {
	OrderRepository interfaces.OerderRepositoryInterface
}

func NewGetTotalUseCase(orderRepository interfaces.OerderRepositoryInterface) *GetTotalUseCase {
	return &GetTotalUseCase{
		OrderRepository: orderRepository,
	}
}

func (c *GetTotalUseCase) Execute() (*GetTotalOutputDTO, error) {
	total, err := c.OrderRepository.GetTotal()
	if err != nil {
		return nil, err
	}

	return &GetTotalOutputDTO{
		Total: total,
	}, nil
}
