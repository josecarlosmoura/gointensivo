package interfaces

import "github.com/josecarlosmoura/gointensivo/internal/order/entity"

type OerderRepositoryInterface interface {
	Save(order *entity.Order) error
	GetTotal() (int, error)
}
