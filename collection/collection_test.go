package collection

import (
	"github.com/magiconair/properties/assert"
	"github.com/shopspring/decimal"
	"reflect"
	"testing"
)

func TestStringArrayCollection_Join(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, Collect(a).Join(""), "hello")
}

var (
	numbers = []int{1, 2, 3, 4, 5, 6, 6, 7, 8, 8}
	foo     = []map[string]interface{}{
		{
			"foo": 10,
		}, {
			"foo": 30,
		}, {
			"foo": 20,
		}, {
			"foo": 40,
		},
	}
)

func TestNumberArrayCollection_Sum(t *testing.T) {
	assert.Equal(t, Collect(numbers).Sum().IntPart(), int64(50))

	var floatTest = []float64{143.66, -14.55}
	//c := floatTest[0] + floatTest[1]
	//fmt.Println(c)

	assert.Equal(t, Collect(floatTest).Sum().String(), "129.11")
}

func TestBaseCollection_Splice(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	c := Collect(a)
	assert.Equal(t, c.Splice(1, 3).ToStringArray(), []string{"e", "l", "l"})
	assert.Equal(t, c.Splice(1).ToStringArray(), []string{"e", "l", "l", "o"})

	assert.Equal(t, Collect(numbers).Splice(2, 1).ToNumberArray(),
		[]decimal.Decimal{nd(3)})

	assert.Equal(t, Collect(foo).Splice(1, 2).ToMapArray(), []map[string]interface{}{
		{
			"foo": 30,
		}, {
			"foo": 20,
		},
	})
}

func TestCollection_Take(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, Collect(a).Take(-2).ToStringArray(), []string{"l", "o"})

	assert.Equal(t, Collect(numbers).Take(4).ToNumberArray(),
		[]decimal.Decimal{nd(1), nd(2), nd(3), nd(4)})

	assert.Equal(t, Collect(foo).Take(2).ToMapArray(), []map[string]interface{}{
		{
			"foo": 10,
		}, {
			"foo": 30,
		},
	})
}

func TestBaseCollection_All(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, Collect(a).All(), []interface{}{"h", "e", "l", "l", "o"})
	assert.Equal(t, len(Collect(numbers).All()), 10)
	assert.Equal(t, Collect(foo).All()[1], map[string]interface{}{"foo": 30})
}

func TestBaseCollection_Mode(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o", "w", "o", "l", "d"}
	foo2 := []map[string]interface{}{
		{
			"foo": 10,
		}, {
			"foo": 30,
		}, {
			"foo": 20,
		}, {
			"foo": 40,
		}, {
			"foo": 40,
		},
	}

	m := Collect(numbers).Mode()
	assert.Equal(t, m[0].(decimal.Decimal).IntPart() == int64(8) ||
		m[0].(decimal.Decimal).IntPart() == int64(6), true)
	assert.Equal(t, m[1].(decimal.Decimal).IntPart() == int64(8) ||
		m[1].(decimal.Decimal).IntPart() == int64(6), true)

	assert.Equal(t, Collect(a).Mode(), []interface{}{"l"})
	assert.Equal(t, Collect(foo2).Mode("foo"), []interface{}{40})
}

func TestBaseCollection_Prepend(t *testing.T) {
	m := map[string]interface{}{
		"foo": 10,
	}
	assert.Equal(t, Collect(m).Prepend("bar", 20).ToMap(), map[string]interface{}{
		"foo": 10,
		"bar": 20,
	})
}

func TestBaseCollection_Pull(t *testing.T) {

	a := map[string]interface{}{
		"product_id": 1,
		"name":       "Desk",
		"price":      100,
		"discount":   false,
	}

	assert.Equal(t, reflect.DeepEqual((Collect(a).Pull("name")).ToMap(), map[string]interface{}{
		"product_id": 1,
		"price":      100,
		"discount":   false,
	}), true)
}
func TestBaseCollection_Put(t *testing.T) {
	a := map[string]interface{}{
		"product_id": 1,
		"name":       "Desk",
		"price":      100,
		"discount":   false,
	}

	assert.Equal(t, reflect.DeepEqual((Collect(a).Put("name1", 111)).ToMap(), map[string]interface{}{
		"product_id": 1,
		"price":      100,
		"discount":   false,
		"name":       "Desk",
		"name1":      111,
	}), true)

}

func TestBaseCollection_Sort(t *testing.T) {

	m := make([]map[string]interface{}, 3)
	c := map[string]interface{}{
		"product_id": 122,
		"name":       "Desk",
		"price":      100,
		"discount":   false,
	}

	m[0] = c
	d := map[string]interface{}{
		"product_id": 1.2,
		"name":       "Desk",
		"price":      1010,
		"discount":   false,
	}
	m[1] = d
	e := map[string]interface{}{
		"product_id": 1,
		"name":       "Desk",
		"price":      1001,
		"discount":   false,
	}
	m[2] = e

	m1 := make([]map[string]interface{}, 3)
	c1 := map[string]interface{}{
		"product_id": 122,
		"name":       "Desk",
		"price":      100,
		"discount":   false,
	}
	e1 := map[string]interface{}{
		"product_id": 1,
		"name":       "Desk",
		"price":      1001,
		"discount":   false,
	}
	d1 := map[string]interface{}{
		"product_id": 1.2,
		"name":       "Desk",
		"price":      1010,
		"discount":   false,
	}

	m1[0] = c1

	m1[1] = e1

	m1[2] = d1
	assert.Equal(t, reflect.DeepEqual((Collect(m).SortBy("price")).Value(), m1), true)
}

func TestBaseCollection_Average(t *testing.T) {

	m := make([]map[string]interface{}, 3)
	c := map[string]interface{}{
		"product_id": 122,
		"name":       "Desk",
		"price":      1,
		"discount":   false,
	}

	m[0] = c
	d := map[string]interface{}{
		"product_id": 1.2,
		"name":       "Desk",
		"price":      2,
		"discount":   false,
	}
	m[1] = d
	e := map[string]interface{}{
		"product_id": 1,
		"name":       "Desk",
		"price":      3,
		"discount":   false,
	}
	m[2] = e

	m1 := make([]map[string]interface{}, 3)
	c1 := map[string]interface{}{
		"product_id": 122,
		"name":       "Desk",
		"price":      4,
		"discount":   false,
	}
	e1 := map[string]interface{}{
		"product_id": 1,
		"name":       "Desk",
		"price":      5,
		"discount":   false,
	}
	d1 := map[string]interface{}{
		"product_id": 1.2,
		"name":       "Desk",
		"price":      6,
		"discount":   false,
	}

	m1[0] = c1

	m1[1] = e1

	m1[2] = d1

	assert.Equal(t, Collect(m).Average("price").IntPart(), int64(2))
}
