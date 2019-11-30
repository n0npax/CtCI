package main

import (
	"container/list"
	"reflect"
	"testing"
)

func Test_findIntersection(t *testing.T) {

	tests := []struct {
		name     string
		numbers1 []int
		numbers2 []int
		want     []int
	}{
		{"len 0 intersection", []int{}, []int{}, []int{}},
		{"len 1 intersection", []int{1}, []int{1}, []int{1}},
		{"len 4 & 4 intersection", []int{1, 2, 2, 1}, []int{3, 2, 2, 1}, []int{2, 2, 1}},
		{"len 4 & 4 no intersection", []int{1, 2, 3, 4}, []int{8, 7, 6, 5}, []int{}},
		{"len 4 & 4 no intersection (tail)", []int{1, 2, 2, 1}, []int{3, 2, 2, 0}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := list.New()
			for _, v := range tt.numbers1 {
				l1.PushBack(v)
			}
			l2 := list.New()
			for _, v := range tt.numbers2 {
				l2.PushBack(v)
			}
			want := list.New()
			for _, v := range tt.want {
				want.PushBack(v)
			}

			got := findIntersection(l1, l2)

			gotS := []int{}
			for el := got.Front(); el != nil; el = el.Next() {
				v, _ := el.Value.(int)
				gotS = append(gotS, v)
			}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("findIntersection(%v, %v) = %v, want %v", tt.numbers1, tt.numbers2, gotS, tt.want)
			}
		})
	}
}
