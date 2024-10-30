package repositories_users

import (
	"fmt"
	"github.com/karlseguin/ccache"
	"time"
	dao "users/dao_users"
)

type CacheConfig struct {
	TTL          time.Duration
	MaxSize      int64
	ItemsToPrune uint32
}

const (
	keyByID    = "user_id:%d"
	keyByEmail = "user_email:%s"
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

func (repository Cache) GetUserById(id int64) (dao.User, error) {
	key := fmt.Sprintf(keyByID, id)
	item := repository.client.Get(key)
	fmt.Println(key)
	if item == nil {
		return dao.User{}, fmt.Errorf("not found item with key %s", key)
	}
	if item.Expired() {
		return dao.User{}, fmt.Errorf("item with key %s is expired", key)
	}
	userDAO, ok := item.Value().(dao.User)
	if !ok {
		return dao.User{}, fmt.Errorf("error converting item with key %s", key)
	}
	return userDAO, nil
}

func (repository Cache) GetUserByEmail(email string) (dao.User, error) {
	fmt.Println("Buscando", email)

	// Use username as cache key
	userKey := fmt.Sprintf(keyByEmail, email)

	// Try to get from cache
	item := repository.client.Get(userKey)
	if item != nil && !item.Expired() {
		// Return cached value
		user, ok := item.Value().(dao.User)
		if !ok {
			return dao.User{}, fmt.Errorf("failed to cast cached value to user")
		}
		return user, nil
	}

	// If not found, return cache miss error
	return dao.User{}, fmt.Errorf("cache miss for email %s", email)
}

func (repository Cache) CreateUser(user dao.User) (int64, error) {
	keyByID := fmt.Sprintf(keyByID, user.User_id)
	keyByEmail := fmt.Sprintf(keyByEmail, user.Email)
	fmt.Println("Guardando", keyByID)
	fmt.Println("Guardando", keyByEmail)
	fmt.Println("saving with duration", repository.ttl)
	repository.client.Set(keyByID, user, repository.ttl)
	repository.client.Set(keyByEmail, user, repository.ttl)
	return user.User_id, nil
}

//agregar login
