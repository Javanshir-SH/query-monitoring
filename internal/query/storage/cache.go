package storage

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Javanshir-SH/query-monitoring/internal/pkg/config"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"time"
)

type QueryMonitoringCache struct {
	rdb             *redis.Client
	CacheExpiration int
}

func NewQueryMonitoringCache(rc *redis.Client, cnf *config.Config) QueryMonitoringCache {
	return QueryMonitoringCache{
		rdb:             rc,
		CacheExpiration: cnf.Redis.CacheExpiration,
	}
}

func (ch QueryMonitoringCache) Set(ctx context.Context, key string, sts ListOfQuery) error {
	value, err := json.Marshal(sts)
	if err != nil {
		return err
	}

	expTime := time.Duration(rand.Int31n(int32(ch.CacheExpiration))) * time.Second
	err = ch.rdb.Set(ctx, key, value, expTime).Err()
	if err != nil {
		return err
	}

	return nil
}

func (ch QueryMonitoringCache) Get(ctx context.Context, key string) (ListOfQuery, error) {
	var res ListOfQuery

	val, err := ch.rdb.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return res, errors.New("item does not exist")
	}

	if err != nil {
		return res, err
	}

	if err := json.Unmarshal(val, &res); err != nil {
		return res, err
	}
	return res, nil
}
