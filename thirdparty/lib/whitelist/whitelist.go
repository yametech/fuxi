package whitelist

import (
	"strings"
	"sync"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/reader"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/util/log"
)

type Whitelist struct {
	conf config.Config
	data map[string]struct{}

	rwlock sync.RWMutex
}

func (srv *Whitelist) update(value reader.Value) error {
	srv.clean()

	srv.rwlock.Lock()
	defer srv.rwlock.Unlock()
	var whiteList []string
	err := value.Scan(&whiteList)
	if err != nil {
		return err
	}
	log.Info("load whitelist: [", strings.Join(whiteList, ",")+"]")
	for _, v := range whiteList {
		srv.data[v] = struct{}{}
	}

	return nil
}

func (srv *Whitelist) In(url string) bool {
	srv.rwlock.Lock()
	defer srv.rwlock.Unlock()
	if _, exist := srv.data[url]; !exist {
		return false
	}
	return true
}

func (srv *Whitelist) clean() {
	srv.rwlock.Lock()
	defer srv.rwlock.Unlock()
	for k := range srv.data {
		delete(srv.data, k)
	}
}

// InitConfig
func InitConfig(source source.Source, path ...string) (*Whitelist, error) {
	whileList := &Whitelist{}
	whileList.conf = config.NewConfig()
	err := whileList.conf.Load(source)
	if err != nil {
		return nil, err
	}
	whileList.data = make(map[string]struct{})
	value := whileList.conf.Get(path...)
	if err := whileList.update(value); err != nil {
		return nil, err
	}
	whileList.enableAutoUpdate(path...)

	return whileList, nil
}

func (srv *Whitelist) enableAutoUpdate(path ...string) {
	go func() {
		for {
			w, err := srv.conf.Watch(path...)
			if err != nil {
				log.Error(err)
			}
			v, err := w.Next()
			if err != nil {
				log.Error(err)
			}
			if err := srv.update(v); err != nil {
				log.Error(err)
			}
		}
	}()
}
