package gormrepo

import (
	"github.com/google/uuid"
	"github.com/mmaxim2710/orders-service/internal/entity"
	"github.com/mmaxim2710/orders-service/pkg/logger"
	"gorm.io/gorm"
)

type TokenRepository struct {
	db *gorm.DB
	l  logger.Interface
}

func NewTokenRepository(db *gorm.DB, l logger.Interface) *TokenRepository {
	return &TokenRepository{
		db: db,
		l:  l,
	}
}

func (t *TokenRepository) Create(userID uuid.UUID, token string) error {
	t.l.Info("tokenRepo - Create: Creating token for user %s", userID.String())
	return t.db.Create(&entity.Token{
		UserID: userID,
		Token:  token,
	}).Error
}

func (t *TokenRepository) GetActiveToken(userID uuid.UUID) (entity.Token, error) {
	t.l.Info("tokenRepo - GetActiveToken: Getting token for user %s", userID.String())
	var token entity.Token
	err := t.db.Where("user_id = ? AND revoked = ?", userID, false).First(&token).Error
	return token, err
}

func (t *TokenRepository) Revoke(userID uuid.UUID) error {
	t.l.Info("tokenRepo - Revoke: Revoking token for user %s", userID.String())
	return t.db.Model(&entity.Token{}).
		Where("user_id = ?", userID).
		Update("revoked", true).Error
}

func (t *TokenRepository) DeleteByUserID(userID uuid.UUID) error {
	t.l.Info("tokenRepo - DeleteByUserID: Deleting tokens for user %s", userID.String())
	return t.db.Model(&entity.Token{}).
		Where("user_id = ?", userID).
		Delete(&entity.Token{}).Error
}
