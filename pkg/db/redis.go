package db

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v7"
	"github.com/jianggushi/proxypool/pkg/model"
)

var (
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	key = "proxy_pool"
)

// Get read proxy info from db by proxystr[Host:Port]
func Get(proxystr string) (*model.Proxy, error) {
	return get(proxystr)
}

func get(proxystr string) (*model.Proxy, error) {
	res, err := rdb.HGet(key, proxystr).Bytes()
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("not found proxy %s", proxystr)
	}
	proxy := &model.Proxy{}
	err = json.Unmarshal(res, proxy)
	if err != nil {
		return nil, err
	}
	return proxy, nil
}

// Put save proxy into db
func Put(proxy *model.Proxy) error {
	return put(proxy)
}

func put(proxy *model.Proxy) error {
	res, err := json.Marshal(proxy)
	if err != nil {
		return err
	}
	proxystr := fmt.Sprintf("%s:%s", proxy.Host, proxy.Port)
	return rdb.HSet(key, proxystr, string(res)).Err()
}

// Delete remove proxy from db
func Delete(proxy *model.Proxy) error {
	proxystr := fmt.Sprintf("%s:%s", proxy.Host, proxy.Port)
	return delete(proxystr)
}

// Delete2 remove proxy from db
func Delete2(proxystr string) error {
	return delete(proxystr)
}

func delete(proxystr string) error {
	return rdb.HDel(key, proxystr).Err()
}

// Update update proxy
func Update(proxy *model.Proxy) error {
	return put(proxy)
}

// GetAll get all proxy in db
func GetAll() ([]*model.Proxy, error) {
	return getAll()
}

func getAll() ([]*model.Proxy, error) {
	res, err := rdb.HGetAll(key).Result()
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, errors.New("no proxy")
	}
	proxies := make([]*model.Proxy, 0, len(res))
	for _, item := range res {
		proxy := &model.Proxy{}
		err = json.Unmarshal([]byte(item), proxy)
		if err != nil {
			return nil, err
		}
		proxies = append(proxies, proxy)
	}
	return proxies, nil
}

func Clear() error {
	return clear()
}

func clear() error {
	return rdb.Del(key).Err()
}
