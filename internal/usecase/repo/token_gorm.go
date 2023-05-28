package repo

import (
	"github.com/google/uuid"
	"github.com/mmaxim2710/orders-service/internal/entity"
	"gorm.io/gorm"
)

type TokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) *TokenRepository {
	return &TokenRepository{
		db: db,
	}
}

func (t *TokenRepository) Create(userID uuid.UUID, token string) error {
	return t.db.Create(&entity.Token{
		UserID: userID,
		Token:  token,
	}).Error
}

func (t *TokenRepository) GetActiveToken(userID uuid.UUID) (entity.Token, error) {
	var token entity.Token
	err := t.db.Where("user_id = ? AND revoked = ?", userID, false).First(&token).Error
	return token, err
}

func (t *TokenRepository) Revoke(userID uuid.UUID) error {
	return t.db.Model(&entity.Token{}).
		Where("user_id = ?", userID).
		Update("revoked", true).Error
}
