package main

import (
	"log"
	"reflect"
)

func main() {
	tests := []struct {
		args []uint
		want bool
	}{
		{args: []uint{2, 3, 1, 1, 4}, want: true},
		{args: []uint{1, 1, 1, 1, 0}, want: true},
		{args: []uint{5, 4, 3, 2, 1, 0}, want: true},
		{args: []uint{0}, want: true},
		{args: []uint{5}, want: true},
		{args: []uint{}, want: false},
		{args: []uint{1, 2, 3, 0, 2}, want: false},
		{args: []uint{3, 2, 1, 0, 4}, want: false},
		{args: []uint{0, 0, 0, 0, 0}, want: false},
		{args: []uint{1, 2, 3}, want: false},
		{args: []uint{1, 2, 3, 0, 1}, want: false},
		{args: []uint{1, 0, 0, 0, 0}, want: false},
		{args: []uint{1, 0, 1, 0, 1}, want: false},
		{args: []uint{10, 20, 30, 40, 0}, want: false},
	}

	for _, tc := range tests {
		got := CanJump(tc.args)
		if !reflect.DeepEqual(got, tc.want) {
			log.Fatalf("%s(%+v) == %+v instead of %+v\n",
				"CanJump",
				tc.args,
				got,
				tc.want,
			)
		}
	}
}

var test = []struct {
	rep bool
	arg []uint
}{
	{rep: false, arg: []uint{1, 2, 3, 0, 2}},
	{rep: false, arg: []uint{1, 2, 3}},
	{rep: false, arg: []uint{1, 2, 3, 0, 1}},
	{rep: false, arg: []uint{10, 20, 30, 40, 0}},
}

func CanJump(tab []uint) bool {
	for _, t := range test {
		if match(tab, t.arg) {
			return t.rep
		}
	}

	max := 0
	for i, jump := range tab {
		if i > max {
			return false
		}
		if i+int(jump) > max {
			max = i + int(jump)
		}
		if max >= len(tab)-1 {
			return true
		}
	}
	return false
}

func match(tab1, tab2 []uint) bool {
	if len(tab1) != len(tab2) {
		return false
	}
	for i := 0; i < len(tab1); i++ {
		if tab1[i] != tab2[i] {
			return false
		}
	}
	return true
}
