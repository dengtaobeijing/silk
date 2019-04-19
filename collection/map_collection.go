package collection

type MapCollection struct {
	value map[string]interface{}
	BaseCollection
}

func (c MapCollection) Only(keys []string) Collection {
	var (
		d MapCollection
		m = make(map[string]interface{}, 0)
	)

	for _, k := range keys {
		m[k] = c.value[k]
	}
	d.value = m
	d.length = len(m)

	return d
}

func (c MapCollection) Prepend(values ...interface{}) Collection {

	var m = copyMap(c.value)
	m[values[0].(string)] = values[1]

	return MapCollection{m, BaseCollection{length: len(m)}}
}

func (c MapCollection) ToMap() map[string]interface{} {
	return c.value
}

func (c MapCollection) Pull(key interface{}) Collection {

	var m = copyMap(c.value)

	delete(m, key.(string))

	return MapCollection{m, BaseCollection{length: len(m)}}

}

func (c MapCollection) Put(key string, value interface{}) Collection {

	var (
		d MapCollection
		m = make(map[string]interface{}, 0)
	)

	m = copyMap(c.value)
	m[key] = value

	d.value = m
	d.length = len(m)

	return d
}
