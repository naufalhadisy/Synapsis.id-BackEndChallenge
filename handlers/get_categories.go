package handlers

import (
	"e-commerce/models"
	"e-commerce/restapi/swagger"

	"github.com/go-openapi/runtime/middleware"
)

type itemListImpl struct{}

func NewItemCategoryListHandler() swagger.CategoryListHandler {
	return &itemListImpl{}
}

func (impl *itemListImpl) Handle(params swagger.CategoryListParams) middleware.Responder {
	responseVal := &models.Categories{
		&models.Category{
			prID:       0011,
			prName:     "Vegetables",
			prIsActive: true,
			prImageURL: "",
			SubCategories: []*models.SubCategory{
				{
					prScID:       1100,
					prScName:     "Carrot",
					prScImageURL: "",
					prScIsActive: true,
				},
			},
		},
	}
	return swagger.NewCategoryListOK().WithPayload(*responseVal)
}
