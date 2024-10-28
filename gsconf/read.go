package gsconf

import (
	"reflect"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/tnnmigga/corev2/conc"
	"github.com/tnnmigga/corev2/log"
	"github.com/tnnmigga/corev2/system"
	"github.com/tnnmigga/corev2/utils"
	"golang.org/x/exp/constraints"
)

var (
	loader  ILoader
	store   = atomic.Value{}
	handles = map[string]func(map[string]any) (any, error){}
	checks  = []func(configs) error{}
	tables  = []string{}
)

type ILoader interface {
	Load() (map[string]any, error)
}

func Init(from ILoader) {
	loader = from
	pull()
	conc.Go(func() {
		ticker := time.NewTicker(time.Minute)
		for {
			select {
			case <-ticker.C:
				utils.ExecAndRecover(pull)
			case <-system.RootCtx().Done():
				return
			}
		}
	})
}

type configs map[string]map[int]any

func (c configs) Get(name string, ID int) TableItem {
	return c[name][ID].(TableItem)
}

func (c configs) Range(name string, h func(item TableItem) bool) {
	table := c[name]
	for _, item := range table {
		if h(item.(TableItem)) {
			return
		}
	}
}

func (c configs) Filter(name string, h func(item TableItem) bool) []TableItem {
	var items []TableItem
	c.Range(name, func(item TableItem) bool {
		if h(item) {
			items = append(items, item)
		}
		return false
	})
	return items
}

func (c configs) All(name string) []TableItem {
	var items []TableItem
	table := c[name]
	for _, item := range table {
		items = append(items, item.(TableItem))
	}
	return items
}

func (c configs) Save(item TableItem) {
	c[item.Table()][item.GetID()] = item
}

func save(m configs) {
	store.Store(m)
}

func read() configs {
	return store.Load().(configs)
}

type TableItem interface {
	Table() string
	GetID() int
}

func RegisterTable[T TableItem](h func(item map[string]any) (*T, error)) {
	var t T
	table := t.Table()
	handles[table] = func(m map[string]any) (any, error) {
		return h(m)
	}
	tables = append(tables, table)
}

func RegisterCheck(h func(newConfs configs) error) {
	checks = append(checks, h)
}

func Get[T TableItem, N constraints.Integer](id N) *T {
	var t T
	table := t.Table()
	m := read()
	confs, ok := m[table]
	if !ok {
		log.Errorf("table not found %#v", new(T))
		return nil
	}
	item, ok := confs[int(id)]
	if !ok {
		// log.Errorf("table item not found %s %d", table, id)
		return nil
	}
	return item.(*T)
}

func Range[T TableItem](h func(*T) (stop bool)) {
	var t T
	table := t.Table()
	m := read()
	for _, v := range m[table] {
		item := v.(*T)
		h(item)
	}
}

func All[T TableItem]() []*T {
	var t T
	table := t.Table()
	m := read()
	items := make([]*T, 0, len(m[table]))
	for _, v := range m[table] {
		items = append(items, v.(*T))
	}
	return items
}

func Filter[T TableItem](h func(*T) bool) []*T {
	var t T
	table := t.Table()
	m := read()
	var items []*T
	for _, v := range m[table] {
		conf := v.(*T)
		if h(conf) {
			items = append(items, conf)
		}
	}
	return items
}

func Find[T TableItem](is func(*T) bool) *T {
	var conf *T
	Range(func(item *T) (stop bool) {
		if is(item) {
			conf = item
			return true
		}
		return false
	})
	return conf
}

func parse[T TableItem](m map[string]any) (*T, error) {
	item := new(T)
	err := mapstructure.WeakDecode(m, item)
	return item, err
}

func pull() {
	confs, err := loader.Load()
	if err != nil {
		log.Panic(err)
	}
	newTables := configs{}
	for _, name := range tables {
		conf := confs[name]
		h := handles[name]
		sub := map[int]any{}
		switch reflect.TypeOf(conf).Kind() {
		case reflect.Slice:
			l := conf.([]any)
			for i, v := range l {
				if v == nil {
					continue
				}
				item, err := h(v.(map[string]any))
				if err != nil {
					log.Panicf("handle table %s %v error", name, v)
				}
				if item == nil { // 无效的数据直接跳过
					continue
				}
				sub[i+1] = item
			}
		case reflect.Map:
			m := conf.(map[string]any)
			for k, v := range m {
				item, err := h(v.(map[string]any))
				if err != nil {
					log.Panicf("handle table %s %v error", name, v)
				}
				if v := reflect.ValueOf(item); v.Kind() == reflect.Ptr && v.IsNil() { // 无效的数据直接跳过
					continue
				}
				id, err := strconv.Atoi(k)
				if err != nil {
					log.Panicf("handler table parse id error %v", id)
				}
				sub[id] = item
			}
		}
		newTables[name] = sub
	}
	for _, check := range checks {
		if err := check(newTables); err != nil {
			log.Panic(err)
		}
	}
	save(newTables)
}
