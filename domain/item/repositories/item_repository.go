package repositories

import (
	"echo-rest-api-mysql/domain/item/models"

	"gorm.io/gorm"
)

type dbItem struct {
	Conn *gorm.DB
}

// Create implements [ItemRepository].
func (db *dbItem) Create(item models.Item) error {
	return db.Conn.Create(&item).Error
}

// Delete implements [ItemRepository].
func (db *dbItem) Delete(idItem int) error {
	result := db.Conn.Where("id_item = ?", idItem).Delete(&models.Item{IdItem: idItem})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// GetAll implements [ItemRepository].
func (db *dbItem) GetAll() ([]models.Item, error) {
	var items []models.Item
	if err := db.Conn.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

// GetByID implements [ItemRepository].
func (db *dbItem) GetByID(idItem int) (models.Item, error) {
	var item models.Item
	if err := db.Conn.Where("id_item", idItem).First(&item).Error; err != nil {
		return models.Item{}, err
	}

	return item, nil
}

// Update implements [ItemRepository].
func (db *dbItem) Update(idItem int, item models.Item) error {
	result := db.Conn.Model(&models.Item{}).Where("id_item = ?", idItem).Updates(item)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

type ItemRepository interface {
	Create(item models.Item) error
	Update(idItem int, item models.Item) error
	Delete(idItem int) error
	GetByID(idItem int) (models.Item, error)
	GetAll() ([]models.Item, error)
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &dbItem{Conn: db}
}
