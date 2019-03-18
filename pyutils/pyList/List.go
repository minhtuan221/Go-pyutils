package pylist

import (
	"encoding/json"

	"github.com/minhtuan221/go-pyutils/pyutils/tryexcept"
)

type iterable interface {
	Len() int
	Contain(key interface{}) bool
}

func Len(list iterable) int {
	return list.Len()
}

func IfKeyIn(key interface{}, list iterable) bool {
	return list.Contain(key)
}

func NewList(x ...interface{}) *List {
	return &List{x}
}

type List struct {
	Values []interface{}
}

// Append one item to list
func (list *List) Append(x interface{}) {
	list.Values = append(list.Values, x)
}

// Extend list by another list
func (list *List) Extend(x []interface{}) {
	list.Values = append(list.Values, x...)
}

// Insert an item in to a list with specified index position
func (list *List) Insert(i int, x interface{}) {
	if i > list.Len() {
		panic("Insert to out of range position in list")
	} else if i < 0 {
		i = list.Len() + i
	}

	// Make space in the array for a new element. You can assign it any value.
	list.Values = append(list.Values, 0)

	// Copy over elements sourced from index 2, into elements starting at index 3.
	copy(list.Values[i+1:], list.Values[i:])

	// assign value to index
	list.Values[i] = x
}

// Remove an specified item in list
func (list *List) Remove(x interface{}) {
	// find value of x
	for i, value := range list.Values {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			list.Values = append(list.Values[:i], list.Values[i+1:]...)
			break
		}
	}
}

// Pop an item in list with specified index
func (list *List) Pop(x ...int) interface{} {
	k := len(list.Values) - 1
	if len(x) == 0 {
		// make a copy of last item
		res := list.Values[k]
		// remove the last item in list
		list.Values = append(list.Values[:k], list.Values[k+1:]...)
		return res
	}
	i := x[0]
	res := list.Values[i]
	// remove item in index i
	list.Values = append(list.Values[:i], list.Values[i+1:]...)
	return res
}

// Popleft the first item add to list
func (list *List) Popleft() {
	list.Pop(0)
}

func (list *List) Clear() {
	list.Values = nil
}

// Index = Delete the item with the specified item
func (list *List) Index(x interface{}) []int {
	// find value of x
	var res []int
	for i, value := range list.Values {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			res = append(res, i)
		}
	}
	return res
}

// Contain find item by item
func (list List) Contain(x interface{}) bool {
	// find value of x
	for _, value := range list.Values {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			return true
		}
	}
	return false
}

// Count how many item in a list
func (list *List) Count(x interface{}) int {
	// find value of x
	res := 0
	for _, value := range list.Values {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			res++
		}
	}
	return res
}

func (list *List) Copy() List {
	res := List{}
	// process a deepcopy
	res.Values = make([]interface{}, len(list.Values))
	copy(res.Values, list.Values)
	return res
}

// FindOneBy => find the first item in a list and return its index. return -1 if not found
func (list *List) FindOneBy(oneItem interface{}, f func(one interface{}, item interface{}) bool) int {
	// find value of x
	for i, value := range list.Values {
		if f(oneItem, value) == true {
			// Where a is the slice, and i is the index of the element you want to delete:
			return i
		}
	}
	return -1
}
func (list *List) ToJson() string {
	data, err := json.Marshal(list.Values)
	if err != nil {
		tryexcept.Throw("Error when converting to json")
		return "{}"
	}
	return string(data)
}

func (list List) Len() int {
	return len(list.Values)
}

type test struct {
	value int
}
