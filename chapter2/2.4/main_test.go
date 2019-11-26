package main

import (
	"container/list"
	"reflect"
	"testing"
)

func Test_rmElFromMiddle(t *testing.T) {

	tests := []struct {
		name    string
		k       int
		numbers []int
		want    []int
	}{
		{"empty", 1, []int{}, []int{}},
		{"nil", 2, nil, []int{}},
		{"len1", 2, []int{7}, []int{7}},
		{"len7", 4, []int{0, 7, 9, 3, 4, 5, 6}, []int{4, 3, 0, 7, 9, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := list.New()
			for _, v := range tt.numbers {
				l.PushBack(v)
			}
			l = partitionByX(l, tt.k)

			got := []int{}
			for element := l.Front(); element != nil; element = element.Next() {
				v, ok := element.Value.(int)
				if !ok {
					t.Errorf("cannot cast value to int")
				}
				got = append(got, v)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionByX(%v) = %v, want %v", tt.numbers, got, tt.want)
			}
		})
	}
}
