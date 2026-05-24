package services

import (
	"echo-rest-api-mysql/domain/item/helpers"
	"echo-rest-api-mysql/domain/item/models"
	"echo-rest-api-mysql/domain/item/repositories"

	"gorm.io/gorm"
)

type itemService struct {
	itemRepo repositories.ItemRepository
}

// Create implements [ItemService].
func (service *itemService) Create(item models.Item) helpers.Response {
	var response helpers.Response

	if err := service.itemRepo.Create(item); err != nil {
		response.Status = "error"
		response.Message = "Failed to create item: " + err.Error()
		return response
	}

	response.Status = "success"
	response.Message = "Item created successfully"
	response.Data = item
	return response
}

// Delete implements [ItemService].
func (service *itemService) Delete(idItem int) helpers.Response {
	var response helpers.Response

	if err := service.itemRepo.Delete(idItem); err != nil {
		response.Status = "error"
		response.Message = "Failed to delete item: " + err.Error()
		return response
	}

	response.Status = "success"
	response.Message = "Item deleted successfully"
	return response
}

// GetAll implements [ItemService].
func (service *itemService) GetAll() helpers.Response {
	var response helpers.Response

	items, err := service.itemRepo.GetAll()
	if err != nil {
		response.Status = "error"
		response.Message = "Failed to retrieve items: " + err.Error()
		return response
	}

	response.Status = "success"
	response.Message = "Items retrieved successfully"
	response.Data = items
	return response
}

// GetByID implements [ItemService].
func (service *itemService) GetByID(idItem int) helpers.Response {
	var response helpers.Response

	item, err := service.itemRepo.GetByID(idItem)
	if err != nil {
		response.Status = "error"
		response.Message = "Failed to retrieve item: " + err.Error()
		return response
	}

	response.Status = "success"
	response.Message = "Item retrieved successfully"
	response.Data = item
	return response
}

// Update implements [ItemService].
func (service *itemService) Update(idItem int, item models.Item) helpers.Response {
	var response helpers.Response

	if err := service.itemRepo.Update(idItem, item); err != nil {
		response.Status = "error"
		response.Message = "Failed to update item: " + err.Error()
		return response
	}

	response.Status = "success"
	response.Message = "Item updated successfully"
	return response
}

type ItemService interface {
	Create(item models.Item) helpers.Response
	Update(idItem int, item models.Item) helpers.Response
	Delete(idItem int) helpers.Response
	GetByID(idItem int) helpers.Response
	GetAll() helpers.Response
}

func NewItemService(db *gorm.DB) ItemService {
	return &itemService{
		itemRepo: repositories.NewItemRepository(db),
	}
}
