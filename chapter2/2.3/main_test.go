package main

import (
	"container/list"
	"reflect"
	"testing"
)

func Test_rmElFromMiddle(t *testing.T) {

	tests := []struct {
		name    string
		numbers []int
		want    []int
	}{
		{"empty", []int{}, []int{}},
		{"nil", nil, []int{}},
		{"len1", []int{7}, []int{7}},
		{"len2", []int{7, 9}, []int{7, 9}},
		{"len3", []int{7, 9, 4}, []int{7, 4}},
		{"len6", []int{7, 9, 3, 4, 5, 6}, []int{7, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := list.New()
			for _, v := range tt.numbers {
				l.PushBack(v)
			}
			if tt.numbers != nil {
				rmElFromMiddle(l)
			} else {
				rmElFromMiddle(nil)
			}

			got := []int{}
			for element := l.Front(); element != nil; element = element.Next() {
				v, ok := element.Value.(int)
				if !ok {
					t.Errorf("cannot cast value to int")
				}
				got = append(got, v)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kToLastElement(%v) = %v, want %v", tt.numbers, got, tt.want)
			}
		})
	}
}
