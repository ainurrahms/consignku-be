package discounts

import (
	discountsUsecase "consignku/bussiness/discounts"
	"time"

	"gorm.io/gorm"
)

type Discounts struct {
	ID            int
	Code          string
	DiscountValue int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

func fromDomain(domain discountsUsecase.Domain) *Discounts {
	return &Discounts{
		ID:            domain.ID,
		Code:          domain.Code,
		DiscountValue: domain.DiscountValue,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
		DeletedAt:     domain.DeletedAt,
	}
}

func (rec *Discounts) toDomain() discountsUsecase.Domain {
	return discountsUsecase.Domain{
		ID:            rec.ID,
		Code:          rec.Code,
		DiscountValue: rec.DiscountValue,
		CreatedAt:     rec.CreatedAt,
		UpdatedAt:     rec.UpdatedAt,
		DeletedAt:     rec.DeletedAt,
	}
}
