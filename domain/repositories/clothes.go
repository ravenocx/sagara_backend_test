package repositories

import (
	"log"

	"github.com/ravenocx/clothes-store/domain/dto"
	"github.com/ravenocx/clothes-store/domain/entities"
	"github.com/ravenocx/clothes-store/utils"
	"gorm.io/gorm"
)

type ClothesRepository interface {
	InsertCloth(cloth *entities.Clothes) (*entities.Clothes, error)
	GetClothes(query dto.GetClothesQuery) ([]entities.Clothes, error)
	// GetClothes() ([]entities.Clothes, error)
	GetClothByID(id string) (*entities.Clothes, error)
	UpdateCloth(cloth *entities.Clothes) (*entities.Clothes, error)
	DeleteCloth(id string) error
	IncreaseStock(cloth *entities.Clothes, stock int) (*entities.Clothes, error)
	DecreaseStock(cloth *entities.Clothes, stock int) (*entities.Clothes, error)
}

type clothesRepository struct {
	db *gorm.DB
}

func NewClothesRepository(db *gorm.DB) *clothesRepository {
	// db = db.Debug()
	return &clothesRepository{db: db}
}

func (r *clothesRepository) InsertCloth(cloth *entities.Clothes) (*entities.Clothes, error) {
	if err := r.db.Create(cloth).Error; err != nil {
		return nil, err
	}
	return cloth, nil
}

func (r *clothesRepository) GetClothes(query dto.GetClothesQuery) ([]entities.Clothes, error) {
	q := "SELECT * FROM clothes"

	whereQuery, args := utils.GetClothesConstructWhereQuery(query)

	if whereQuery != "" {
		q += whereQuery
	}

	log.Printf("Get clothes query : %s", q)
	var clothes []entities.Clothes
	if err := r.db.Raw(q, args...).Scan(&clothes).Error; err != nil {
		return nil, err
	}

	return clothes, nil
}

func (r *clothesRepository) GetClothByID(id string) (*entities.Clothes, error) {
	var cloth entities.Clothes

	if err := r.db.Where("id = ?", id).First(&cloth).Error; err != nil {
		return nil, err
	}
	return &cloth, nil
}

func (r *clothesRepository) UpdateCloth(cloth *entities.Clothes) (*entities.Clothes, error) {
	var updatedCloth entities.Clothes

	if cloth.Color != "" {
		updatedCloth.Color = cloth.Color
	}

	if cloth.Size != "" {
		updatedCloth.Size = cloth.Size
	}

	if cloth.Price != 0 {
		updatedCloth.Price = cloth.Price
	}

	if cloth.Stock != 0 {
		updatedCloth.Stock = cloth.Stock
	}

	log.Printf("Cloth to update : %+v", updatedCloth)
	if err := r.db.Updates(&updatedCloth).Error; err != nil {
		return nil, err
	}

	return cloth, nil
}

func (r *clothesRepository) DeleteCloth(id string) error {
	var cloth entities.Clothes
	if err := r.db.Where("id = ?", id).Delete(&cloth).Error; err != nil {
		return err
	}
	return nil
}

func (r *clothesRepository) IncreaseStock(cloth *entities.Clothes, stock int) (*entities.Clothes, error) {
	var updatedCloth entities.Clothes

	updatedCloth.Stock += stock

	log.Printf("Clothes to update stock (increase) : %+v", updatedCloth)
	if err := r.db.Updates(&updatedCloth).Error; err != nil {
		return nil, err
	}

	return cloth, nil
}

func (r *clothesRepository) DecreaseStock(cloth *entities.Clothes, stock int) (*entities.Clothes, error) {
	var updatedCloth entities.Clothes

	updatedCloth.Stock -= stock

	log.Printf("Clothes to update stock (reduce) : %+v", updatedCloth)
	if err := r.db.Updates(&updatedCloth).Error; err != nil {
		return nil, err
	}
	
	return cloth, nil
}
