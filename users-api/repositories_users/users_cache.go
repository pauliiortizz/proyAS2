package repositories_users

import (
	"context"
	"fmt"
	"github.com/karlseguin/ccache"
	"time"
	usersDAO "users/dao_users"
)

const (
	keyFormat = "user:%s"
)

type CacheConfig struct {
	MaxSize      int64
	ItemsToPrune uint32
	Duration     time.Duration
}

type Cache struct {
	client   *ccache.Cache
	duration time.Duration
}

func NewCache(config CacheConfig) Cache {
	client := ccache.New(ccache.Configure().
		MaxSize(config.MaxSize).
		ItemsToPrune(config.ItemsToPrune))
	return Cache{
		client:   client,
		duration: config.Duration,
	}
}

func (repository Cache) GetUserByID(ctx context.Context, id string) (usersDAO.User, error) {
	key := fmt.Sprintf(keyFormat, id)
	item := repository.client.Get(key)
	fmt.Println(key)
	if item == nil {
		return usersDAO.User{}, fmt.Errorf("not found item with key %s", key)
	}
	if item.Expired() {
		return usersDAO.User{}, fmt.Errorf("item with key %s is expired", key)
	}
	userDAO, ok := item.Value().(usersDAO.User)
	if !ok {
		return usersDAO.User{}, fmt.Errorf("error converting item with key %s", key)
	}
	return userDAO, nil
}

func (repository Cache) Create(ctx context.Context, user usersDAO.User) (string, error) {
	key := fmt.Sprintf(keyFormat, user.User_id)
	fmt.Println("saving with duration", repository.duration)
	repository.client.Set(key, user, repository.duration)
	return user.User_id, nil
}
