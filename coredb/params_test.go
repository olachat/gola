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

	p.Add([]string{"2", "two"})
	out = p.Get()
	as.Equal(1, out[0])
	as.Equal("1", out[1])
	as.Equal(1.0, out[2])
	as.Equal("2", out[3])
	as.Equal("two", out[4])

	p.Add(0.6, 77)
	out = p.Get()
	as.Equal(1, out[0])
	as.Equal("1", out[1])
	as.Equal(1.0, out[2])
	as.Equal("2", out[3])
	as.Equal("two", out[4])
	as.Equal(0.6, out[5])
	as.Equal(77, out[6])
}
