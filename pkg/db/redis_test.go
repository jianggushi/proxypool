package db

import (
	"testing"

	"github.com/jianggushi/proxypool/pkg/model"
)

func init() {
	key = "proxy_pool_test"
}

func Test_Connect(t *testing.T) {
	pong, err := rdb.Ping().Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(pong)
}

func Test_Get1(t *testing.T) {
	err := Clear()
	if err != nil {
		t.Error(err)
	}
	proxystr := "1.1.1.1:8080"
	_, err = Get(proxystr)
	if err == nil {
		t.Log(err)
	}
}

func Test_Get2(t *testing.T) {
	err := Clear()
	if err != nil {
		t.Error(err)
	}
	err = Put(&model.Proxy{
		Host: "1.1.1.1",
		Port: "8080",
	})
	if err != nil {
		t.Error(err)
	}
	proxystr := "1.1.1.1:8080"
	_, err = Get(proxystr)
	if err != nil {
		t.Error(err)
	}
}
