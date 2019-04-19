package collection

import (
	"github.com/shopspring/decimal"
)

type MapArrayCollection struct {
	value []map[string]interface{}
	BaseCollection
}

func (c MapArrayCollection) Value() interface{} {
	return c.value
}

func (c MapArrayCollection) Sum(key ...string) decimal.Decimal {
	var sum = decimal.New(0, 0)

	for i := 0; i < len(c.value); i++ {
		sum = sum.Add(newDecimalFromInterface(c.value[i][key[0]]))
	}

	return sum
}

func (c MapArrayCollection) Min(key ...string) decimal.Decimal {

	var (
		smallest = decimal.New(0, 0)
		number   decimal.Decimal
	)

	for i := 0; i < len(c.value); i++ {
		number = newDecimalFromInterface(c.value[i][key[0]])
		if i == 0 {
			smallest = number
			continue
		}
		if smallest.GreaterThan(number) {
			smallest = number
		}
	}

	return smallest
}

func (c MapArrayCollection) Max(key ...string) decimal.Decimal {

	var (
		biggest = decimal.New(0, 0)
		number  decimal.Decimal
	)

	for i := 0; i < len(c.value); i++ {
		number = newDecimalFromInterface(c.value[i][key[0]])
		if i == 0 {
			biggest = number
			continue
		}
		if biggest.LessThan(number) {
			biggest = number
		}
	}

	return biggest
}

func (c MapArrayCollection) Pluck(key string) Collection {
	var s = make([]interface{}, 0)
	for i := 0; i < len(c.value); i++ {
		s = append(s, c.value[i][key])
	}
	return Collect(s)
}

func (c MapArrayCollection) Prepend(values ...interface{}) Collection {

	var d MapArrayCollection

	var n = make([]map[string]interface{}, len(c.value))
	copy(n, c.value)

	d.value = append([]map[string]interface{}{values[0].(map[string]interface{})}, n...)
	d.length = len(d.value)

	return d
}

func (c MapArrayCollection) Only(keys []string) Collection {
	var d MapArrayCollection

	var ma = make([]map[string]interface{}, 0)
	for _, k := range keys {
		m := make(map[string]interface{}, 0)
		for _, v := range c.value {
			m[k] = v[k]
		}
		ma = append(ma, m)
	}
	d.value = ma
	d.length = len(ma)

	return d
}

func (c MapArrayCollection) Splice(index ...int) Collection {

	if len(index) == 1 {
		var n = make([]map[string]interface{}, len(c.value))
		copy(n, c.value)
		n = n[index[0]:]

		return MapArrayCollection{n, BaseCollection{length: len(n)}}
	} else if len(index) > 1 {
		var n = make([]map[string]interface{}, len(c.value))
		copy(n, c.value)
		n = n[index[0] : index[0]+index[1]]

		return MapArrayCollection{n, BaseCollection{length: len(n)}}
	} else {
		panic("invalid argument")
	}
}

func (c MapArrayCollection) Take(num int) Collection {
	var d MapArrayCollection
	if num > c.length {
		panic("not enough elements to take")
	}

	if num >= 0 {
		d.value = c.value[:num]
		d.length = num
	} else {
		d.value = c.value[len(c.value)+num:]
		d.length = 0 - num
	}

	return d
}

func (c MapArrayCollection) All() []interface{} {
	s := make([]interface{}, len(c.value))
	for i := 0; i < len(c.value); i++ {
		s[i] = c.value[i]
	}

	return s
}

func (c MapArrayCollection) Mode(key ...string) []interface{} {
	valueCount := make(map[interface{}]int)
	for i := 0; i < c.length; i++ {
		if v, ok := c.value[i][key[0]]; ok {
			valueCount[v]++
		}
	}

	maxCount := 0
	maxValue := make([]interface{}, len(valueCount))
	for v, c := range valueCount {
		switch {
		case c < maxCount:
			continue
		case c == maxCount:
			maxValue = append(maxValue, v)
		case c > maxCount:
			maxValue = append([]interface{}{}, v)
			maxCount = c
		}
	}
	return maxValue
}

func (c MapArrayCollection) ToMapArray() []map[string]interface{} {
	return c.value
}
func (c MapArrayCollection) SortBy(key string) Collection {

	var (
		d = make([]map[string]interface{}, c.length)
		//[]map[string]interface{}
		m    = make([]map[string]interface{}, 0)
		keys = make(map[decimal.Decimal]int, 0)
	)

	m = c.value

	for i, v := range m {
		keys[newDecimalFromInterface(v[key])] = i
	}

	sortKeys := make(map[int]decimal.Decimal, 0)
	index := 0
	for i := range keys {
		sortKeys[index] = i
		index = index + 1
	}
	var temp decimal.Decimal
	for i := 0; i < c.length-1; i++ {
		for j := 0; j < c.length-1-i; j++ {
			if sortKeys[j].GreaterThanOrEqual(sortKeys[j+1]) {
				temp = sortKeys[j]
				sortKeys[j] = sortKeys[j+1]
				sortKeys[j+1] = temp
			}
		}
	}

	index = 0

	for _, v := range sortKeys {
		for _, v1 := range m {
			if v.Equal(newDecimalFromInterface(v1[key])) {
				d[index] = v1
				index = index + 1
			}
		}
	}
	return MapArrayCollection{d, BaseCollection{length: len(d)}}

	//return MapArrayCollection(d,d.)
}
