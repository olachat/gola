package coredb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParams(t *testing.T) {
	as := assert.New(t)
	p := NewParams(1, "1", 1.0)
	out := p.Get()
	as.Equal(1, out[0])
	as.Equal("1", out[1])
	as.Equal(1.0, out[2])

	p.Add([]string{"2", "two"}, "king", -66)
	out = p.Get()
	as.Equal(1, out[0])
	as.Equal("1", out[1])
	as.Equal(1.0, out[2])
	as.Equal("2", out[3])
	as.Equal("two", out[4])
	as.Equal("king", out[5])
	as.Equal(-66, out[6])

	p.Add(0.6, 77, []int{88, 99, 100})
	out = p.Get()
	as.Equal(1, out[0])
	as.Equal("1", out[1])
	as.Equal(1.0, out[2])
	as.Equal("2", out[3])
	as.Equal("two", out[4])
	as.Equal("king", out[5])
	as.Equal(-66, out[6])
	as.Equal(0.6, out[7])
	as.Equal(77, out[8])
	as.Equal(88, out[9])
	as.Equal(99, out[10])
	as.Equal(100, out[11])

	p2 := NewParams(1, 2, []int{3, 4, 5}, "Mary", []string{"Wilson", "Mandy"})
	p2.Get()
	as.Equal([]any{1, 2, 3, 4, 5, "Mary", "Wilson", "Mandy"}, p2.Get())
}
