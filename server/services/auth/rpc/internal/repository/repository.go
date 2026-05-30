package repository

import (
	"github.com/menli02/advertisement/server/services/auth/rpc/internal/repository/cache"
	"github.com/menli02/advertisement/server/services/auth/rpc/internal/repository/db"
)

type Repository struct {
	User *db.UserRepository
	OTP  *cache.OTPCache
}
