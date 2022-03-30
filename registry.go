package pac

func NewRegistry[T any]() *Registry[T] {
	reg := Registry[T]{
		items: make(map[string]T),
	}

	return &reg
}

type Registry[T any] struct {
	items map[string]T
}

func (r *Registry[T]) Get(key string) *T {
	item, ok := r.items[key]

	if !ok {
		return nil
	}

	return &item
}

func (r *Registry[T]) Add(key string, value T) {
	r.items[key] = value
}

func (r *Registry[T]) Delete(key string) {
	delete(r.items, key)
}
