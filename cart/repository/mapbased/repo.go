package mapbased

import (
	"context"
	"fmt"
	"github.com/kennykarnama/checkout-challenge/cart/entity"
	"sync"
)

type MemoryMappedRepository struct {
	itemByID sync.Map
	mut      *sync.Mutex
}

func NewMemoryMappedRepository() *MemoryMappedRepository {
	m := &MemoryMappedRepository{
		itemByID: sync.Map{},
		mut:      &sync.Mutex{},
	}
	return m
}

func (m *MemoryMappedRepository) Add(ctx context.Context, item *entity.CartItem) error {
	m.mut.Lock()
	defer m.mut.Unlock()

	raw, ok := m.itemByID.Load(item.ID)
	if !ok {
		m.itemByID.Store(item.ID, []*entity.CartItem{
			item,
		})
	} else {
		existing, valid := raw.([]*entity.CartItem)
		if !valid {
			return fmt.Errorf("action=repo.add item=%s err=%v", *item, "not valid instance")
		}
		existing = append(existing, item)
		m.itemByID.Store(item.ID, existing)
	}
	return nil
}

func (m *MemoryMappedRepository) GetCartByID(ctx context.Context, ID string) ([]*entity.CartItem, error) {
	raw, ok := m.itemByID.Load(ID)
	if !ok {
		return nil, nil
	}
	existing, valid := raw.([]*entity.CartItem)
	if !valid {
		return nil, fmt.Errorf("action=repo.getCartByID id=%s err=%v", ID, "not valid instance")
	}
	return existing, nil
}
