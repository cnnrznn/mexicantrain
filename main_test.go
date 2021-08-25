package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLongest(t *testing.T) {
	ns := []*Node{
		{
			Domino: Domino{
				n1: 1,
				n2: 2,
			},
		},
		{
			Domino: Domino{
				n1: 4,
				n2: 5,
			},
		},
		{
			Domino: Domino{
				n1: 3,
				n2: 4,
			},
		},
		{
			Domino: Domino{
				n1: 3,
				n2: 2,
			},
		},
	}

	train := findLongest(ns, 5)

	assert.Equal(t, train, []Domino{
		{
			n1: 5,
			n2: 4,
		},
		{
			n1: 4,
			n2: 3,
		},
		{
			n1: 3,
			n2: 2,
		},
		{
			n1: 2,
			n2: 1,
		},
	})
	train = findLongest(ns, 1)

	assert.Equal(t, train, []Domino{
		{
			n1: 1,
			n2: 2,
		},
		{
			n1: 2,
			n2: 3,
		},
		{
			n1: 3,
			n2: 4,
		},
		{
			n1: 4,
			n2: 5,
		},
	})
}
