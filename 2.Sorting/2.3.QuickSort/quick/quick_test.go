package quick

import (
	"math/rand"
	"testing"
	"time"
)

type quick []int

func (q quick) Len() int {
	return len(q)
}

func (q quick) Less(i, j int) bool {
	return q[i] < q[j]
}

func (q quick) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q quick) Divide(i, j int) Interface {
	return q[i:j]
}

func (q quick) IsSorted() bool {
	for i := 1; i < q.Len(); i++ {
		if q.Less(i, i-1) {
			return false
		}
	}
	return true
}

func Test_Sort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	q := make(quick, 10)
	for i := 0; i < len(q); i++ {
		q[i] = len(q) - i - 1
	}
	t.Log(q)
	Sort(q)
	t.Log(q)
	if !q.IsSorted() {
		t.Error("没能排序好")
	}
}

/*
func Test_Sort_1000000(t *testing.T) {
	q := make(quick, 1000000)
	for i := 0; i < q.Len(); i++ {
		q[i] = q.Len() - i - 1
	}
	Sort(q)
	if !q.IsSorted() {
		t.Error("没能排序好")
	}
}
*/
