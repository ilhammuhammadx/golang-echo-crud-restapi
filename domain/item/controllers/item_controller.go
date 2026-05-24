package controllers

import (
	"net/http"
	"strconv"

	"echo-rest-api-mysql/domain/item/models"
	"echo-rest-api-mysql/domain/item/services"
	"echo-rest-api-mysql/validators"

	"github.com/labstack/echo/v4"
)

type ItemController struct {
	itemService services.ItemService
	validate    validators.CustomValidator
}

func (controller *ItemController) Create(c echo.Context) error {
	type payload struct {
		NamaItem    string  `json:"nama_item" validate:"required"`
		Unit        string  `json:"unit" validate:"required"`
		Stok        int     `json:"stok" validate:"required"`
		HargaSatuan float64 `json:"harga_satuan" validate:"required"`
	}

	payloadValidator := new(payload)
	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload: " + err.Error()})
	}

	if err := controller.validate.Validate(payloadValidator); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Validation failed: " + err.Error()})
	}

	result := controller.itemService.Create(models.Item{
		NamaItem:    payloadValidator.NamaItem,
		Unit:        payloadValidator.Unit,
		Stok:        payloadValidator.Stok,
		HargaSatuan: payloadValidator.HargaSatuan,
	})

	return c.JSON(http.StatusOK, result)
}

func (controller *ItemController) GetAll(c echo.Context) error {
	result := controller.itemService.GetAll()
	return c.JSON(http.StatusOK, result)
}

func (controller *ItemController) GetByID(c echo.Context) error {
	idItem, err := getIDParam(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID parameter: " + err.Error()})
	}

	result := controller.itemService.GetByID(idItem)
	return c.JSON(http.StatusOK, result)
}

func (controller *ItemController) Update(c echo.Context) error {
	idItem, err := getIDParam(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID parameter: " + err.Error()})
	}

	type payload struct {
		NamaItem    string  `json:"nama_item" validate:"required"`
		Unit        string  `json:"unit" validate:"required"`
		Stok        int     `json:"stok" validate:"required"`
		HargaSatuan float64 `json:"harga_satuan" validate:"required"`
	}

	payloadValidator := new(payload)
	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload: " + err.Error()})
	}

	if err := controller.validate.Validate(payloadValidator); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Validation failed: " + err.Error()})
	}

	result := controller.itemService.Update(idItem, models.Item{
		NamaItem:    payloadValidator.NamaItem,
		Unit:        payloadValidator.Unit,
		Stok:        payloadValidator.Stok,
		HargaSatuan: payloadValidator.HargaSatuan,
	})
	return c.JSON(http.StatusOK, result)
}

func (controller *ItemController) Delete(c echo.Context) error {
	idItem, err := getIDParam(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID parameter: " + err.Error()})
	}

	result := controller.itemService.Delete(idItem)
	return c.JSON(http.StatusOK, result)
}

func NewItemController(
	itemService services.ItemService,
	validate validators.CustomValidator,
) *ItemController {
	return &ItemController{
		itemService: itemService,
		validate:    validate,
	}
}

func getIDParam(c echo.Context) (int, error) {
	return strconv.Atoi(c.Param("id_item"))
}
