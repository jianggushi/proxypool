package filter

import (
	"errors"

	"github.com/jianggushi/proxypool/pkg/model"
)

// see: https://github.com/sparrc/go-ping
func ping(proxy *model.Proxy) (int, error) {
	return 0, errors.New("Unsupport")
}
