package services

import (
	"net/http"

	"github.com/ravenocx/clothes-store/db"
	"github.com/ravenocx/clothes-store/domain/dto"
	"github.com/ravenocx/clothes-store/domain/entities"
	"github.com/ravenocx/clothes-store/domain/repositories"
	"github.com/ravenocx/clothes-store/utils"
)

type ClothesService interface {
	InsertCloth(cloth *entities.Clothes) (*entities.Clothes, error)
	GetClothes(query dto.GetClothesQuery) ([]entities.Clothes, error)
	// GetClothByID() (*entities.Clothes, error)
	UpdateCloth(cloth *entities.Clothes) (*entities.Clothes, error)
	DeleteCloth(id string) error
	// IncreaseStock(cloth *entities.Clothes, stock int) (*entities.Clothes, error)
	// DecreaseStock(cloth *entities.Clothes, stock int) (*entities.Clothes, error)
}

type clothesService struct {
	clothesRepository repositories.ClothesRepository
}

func NewClothesService(clothesRepository repositories.ClothesRepository) *clothesService {
	return &clothesService{clothesRepository: clothesRepository}
}

func (s *clothesService) InsertCloth(cloth *entities.Clothes) (*entities.Clothes, error) {
	db, err := db.OpenConnection()
	if err != nil {
		return nil, &utils.ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	r := repositories.NewClothesRepository(db)

	cloth, err = r.InsertCloth(cloth)

	if err != nil {
		return nil, &utils.ErrorMessage{
			Message: "Failed to add new cloth",
			Code:    http.StatusInternalServerError,
		}
	}

	return cloth, nil
}

func (s *clothesService) GetClothes(query dto.GetClothesQuery) ([]entities.Clothes, error) {
	db, err := db.OpenConnection()
	if err != nil {
		return nil, &utils.ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	r := repositories.NewClothesRepository(db)

	cloth, err := r.GetClothes(query)

	if err != nil {
		return nil, &utils.ErrorMessage{
			Message: "Failed to get clothes data",
			Code:    http.StatusInternalServerError,
		}
	}

	return cloth, nil
}

func (s *clothesService) UpdateCloth(cloth *entities.Clothes) (*entities.Clothes, error) {
	db, err := db.OpenConnection()
	if err != nil {
		return nil, &utils.ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	r := repositories.NewClothesRepository(db)

	_, err = r.GetClothByID(cloth.ID)

	if err != nil {
		return nil, &utils.ErrorMessage{
			Message: "Cloth not found",
			Code:    http.StatusNotFound,
		}
	}

	cloth, err = r.UpdateCloth(cloth)

	if err != nil {
		return nil, &utils.ErrorMessage{
			Message: "Failed to update cloth to database",
			Code : http.StatusInternalServerError,
		}
	}

	return cloth, nil
}

func (s *clothesService) DeleteCloth(id string) error {
	db, err := db.OpenConnection()

	if err != nil {
		return &utils.ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	r := repositories.NewClothesRepository(db)

	_, err = r.GetClothByID(id)

	if err != nil {
		return &utils.ErrorMessage{
			Message: "Failed to get cloth data",
			Code:    http.StatusInternalServerError,
		}
	}

	err = r.DeleteCloth(id)

	if err != nil {
		return &utils.ErrorMessage{
			Message: "Failed to delete cloth data",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
} 
