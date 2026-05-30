package repository

import (
	"github.com/menli02/advertisement/server/services/ad/rpc/internal/repository/cache"
	"github.com/menli02/advertisement/server/services/ad/rpc/internal/repository/db"
)

type Repository struct {
	Ad       *db.AdRepository
	Category *db.CategoryRepository
	Image    *db.ImageRepository
	AdCache  *cache.AdCache
}
