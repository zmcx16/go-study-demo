package demo

import (
	"time"

	"github.com/gin-gonic/gin"
	mwLimiter "github.com/ulule/limiter/v3"
	mwGin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	mwMemory "github.com/ulule/limiter/v3/drivers/store/memory"
)

func NewRateLimiter(rps int) gin.HandlerFunc {

	store := mwMemory.NewStore()
	rate := mwLimiter.Rate{
		Period: 1 * time.Second,
		Limit:  int64(rps),
	}

	rateLimiter := mwGin.NewMiddleware(mwLimiter.New(store, rate))
	return rateLimiter
}
