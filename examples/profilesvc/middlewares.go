package profilesvc

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Service) Service

type Emid endpoint.Middleware

func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			defer func(begin time.Time) {
				logger.Log("method", reflect.TypeOf(request), "took", time.Since(begin))
			}(time.Now())
			return next(ctx, request)
		}
	}
}

type IsAdminMiddleware struct {
	Service
}

func NewIsAdminMiddleware(s Service) Service {
	return &IsAdminMiddleware{s}
}

func (mw IsAdminMiddleware) DeleteProfile(ctx context.Context, id string) (err error) {
	return errors.New("No, Bob, you're not an admin, I will not shred this file.")
}
