package cache

type Item struct {
	Object interface{}
}

type Cache struct {
	items map[string]Item
}

func New() *Cache {
	items := make(map[string]Item)
	return &Cache{items}
}

func (c *Cache) Set(k string, x interface{}) {

	c.items[k] = Item{
		Object: x,
	}

}

func (c *Cache) Get(k string) interface{} {
	item, found := c.items[k]
	if !found {
		return nil
	} else {
		return item.Object
	}
}

func (c *Cache) Delete(k string) {
	delete(c.items, k)
}
