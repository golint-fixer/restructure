package main

import (
	"reflect"
	"testing"
)

func TestRestructure(t *testing.T) {
	golden := []struct {
		path string
		want []*Primitive
	}{
		{
			path: "testdata/foo.dot",
			want: []*Primitive{
				{
					Prim:  "list",
					Node:  "list0",
					Nodes: map[string]string{"A": "F", "B": "G"},
				},
				{
					Prim:  "if",
					Node:  "if0",
					Nodes: map[string]string{"A": "E", "B": "list0", "C": "H"},
				},
			},
		},
	}

	for i, g := range golden {
		got, err := restructure(g.path)
		if err != nil {
			t.Errorf("i=%d: error; %v", i, err)
			continue
		}
		if !reflect.DeepEqual(got, g.want) {
			t.Errorf("i=%d: primitive mismatch; expected %v, got %v", i, g.want, got)
		}
	}
}