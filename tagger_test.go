package tagger

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type Order struct {
	Email   string
	Items   []Item
	Value   float32
	Address OrderAddress
}
type Item struct {
	Title string
}
type OrderAddress struct {
	Postcode string
}

func TestContains(t *testing.T) {
	order := Order{
		Email: "hello@loxhill.com",
		Items: []Item{
			{
				Title: "16 Pack Rechargeable AA Batteries",
			},
		},
	}
	tags := timeAndStart(order)
	assert.Contains(t, tags, "loxhill")
	assert.NotContains(t, tags, "multiple-items")
}

func TestCount(t *testing.T) {
	order := Order{
		Email: "user1@gmail.com",
		Items: []Item{
			{
				Title: "10 Colors Refrigerator Magnets",
			},
			{
				Title: "Mens Jumper Classic Sweater with V-Neck and Long Sleeve",
			},
		},
	}
	tags := timeAndStart(order)
	assert.Contains(t, tags, "multiple-items")
	assert.Contains(t, tags, "clothing")
}

func timeAndStart(order Order) (tags []string) {
	start := time.Now()
	tagger := NewTagger()
	tagger.LoadRules("resources/ruleset_test.json")
	tags = tagger.Tag(order)
	elapsed := time.Since(start)
	fmt.Printf("Got %v in %s\n", tags, elapsed)
	return tags
}
