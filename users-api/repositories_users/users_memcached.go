package repositories_users

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"users/dao_users"
)

type MemcachedConfig struct {
	Host string
	Port string
}

type Memcached struct {
	client *memcache.Client
}

func idKey(id int64) string {
	return fmt.Sprintf("id:%d", id)
}

func emailKey(email string) string {
	return fmt.Sprintf("email:%s", email)
}

func NewMemcached(config MemcachedConfig) Memcached {
	// Connect to Memcached
	address := fmt.Sprintf("%s:%s", config.Host, config.Port)
	client := memcache.New(address)

	return Memcached{client: client}
}

func (repository Memcached) GetUserById(id int64) (users.User, error) {
	// Retrieve the user from Memcached
	key := idKey(id)
	item, err := repository.client.Get(key)
	if err != nil {
		if errors.Is(err, memcache.ErrCacheMiss) {
			return users.User{}, fmt.Errorf("user not found")
		}
		return users.User{}, fmt.Errorf("error fetching user from memcached: %w", err)
	}

	// Deserialize the data
	var user users.User
	if err := json.Unmarshal(item.Value, &user); err != nil {
		return users.User{}, fmt.Errorf("error unmarshaling user: %w", err)
	}
	return user, nil
}

func (repository Memcached) GetByEmail(email string) (users.User, error) {
	// Assume we store users with "email:<email>" as key
	key := emailKey(email)
	item, err := repository.client.Get(key)
	if err != nil {
		if errors.Is(err, memcache.ErrCacheMiss) {
			return users.User{}, fmt.Errorf("user not found")
		}
		return users.User{}, fmt.Errorf("error fetching user by email from memcached: %w", err)
	}

	// Deserialize the data
	var user users.User
	if err := json.Unmarshal(item.Value, &user); err != nil {
		return users.User{}, fmt.Errorf("error unmarshaling user: %w", err)
	}

	return user, nil
}

func (repository Memcached) CreateUser(user users.User) (int64, error) {
	// Serialize user data
	data, err := json.Marshal(user)
	if err != nil {
		return 0, fmt.Errorf("error marshaling user: %w", err)
	}

	// Store user with ID as key and email as an alternate key
	idKey := idKey(user.User_id)
	if err := repository.client.Set(&memcache.Item{Key: idKey, Value: data}); err != nil {
		return 0, fmt.Errorf("error storing user in memcached: %w", err)
	}

	// Set key for email as well for easier lookup by email
	emailKey := emailKey(user.Email)
	if err := repository.client.Set(&memcache.Item{Key: emailKey, Value: data}); err != nil {
		return 0, fmt.Errorf("error storing email in memcached: %w", err)
	}

	return user.User_id, nil
}

//agregar login
