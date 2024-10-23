package repositories_users

import (
	"fmt"
	"github.com/karlseguin/ccache"
	"time"
	"users/dao_users"
)

type CacheConfig struct {
	TTL          time.Duration
	MaxSize      int64
	ItemsToPrune uint32
}

const (
	keyFormat = "user:%s"
)

type Cache struct {
	client *ccache.Cache
	ttl    time.Duration
}

func NewCache(config CacheConfig) Cache {
	client := ccache.New(ccache.Configure().
		MaxSize(config.MaxSize).
		ItemsToPrune(config.ItemsToPrune))
	return Cache{
		client: client,
		ttl:    config.TTL,
	}
}

func (repository Cache) GetUserById(id string) (users.User, error) {
	key := fmt.Sprintf(keyFormat, id)
	item := repository.client.Get(key)
	fmt.Println(key)
	if item == nil {
		return users.User{}, fmt.Errorf("not found item with key %s", key)
	}
	if item.Expired() {
		return users.User{}, fmt.Errorf("item with key %s is expired", key)
	}
	userDAO, ok := item.Value().(users.User)
	if !ok {
		return users.User{}, fmt.Errorf("error converting item with key %s", key)
	}
	return userDAO, nil
}

func (repository Cache) GetUserByEmail(email string) (users.User, error) {
	// Use username as cache key
	userKey := fmt.Sprintf("user:username:%s", email)

	// Try to get from cache
	item := repository.client.Get(userKey)
	if item != nil && !item.Expired() {
		// Return cached value
		user, ok := item.Value().(users.User)
		if !ok {
			return users.User{}, fmt.Errorf("failed to cast cached value to user")
		}
		return user, nil
	}

	// If not found, return cache miss error
	return users.User{}, fmt.Errorf("cache miss for username %s", email)
}

func (repository Cache) CreateUser(user users.User) (int64, error) {
	key := fmt.Sprintf(keyFormat, user.User_id)
	fmt.Println("saving with duration", repository.ttl)
	repository.client.Set(key, user, repository.ttl)
	return user.User_id, nil
}

//agregar login
