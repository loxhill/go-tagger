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
	Title   string
	Price   float32
	Options ItemOptions
}
type ItemOptions struct {
	Colour string
	Weight float32
}
type OrderAddress struct {
	Residential bool
	Postcode    string
}

func TestContains(t *testing.T) {
	order := Order{
		Email: "hello@loxhill.com",
		Items: []Item{
			{
				Title: "Mens Jumper Classic Sweater with V-Neck and Long Sleeve",
				Options: ItemOptions{
					Colour: "Red",
					Weight: 0.1,
				},
			},
		},
	}
	tags := timeAndStart(order)
	assert.Contains(t, tags, "loxhill")
	assert.NotContains(t, tags, "multiple-items")
}

func TestCount(t *testing.T) {
	order := Order{
		Email: "hello@loxhill.com",
		Items: []Item{
			{
				Title: "Colored Refrigerator Magnets (10-pack)",
				Options: ItemOptions{
					Colour: "Multi",
					Weight: 0.2,
				},
			},
			{
				Title: "Mens Jumper Classic Sweater with V-Neck and Long Sleeve",
				Options: ItemOptions{
					Colour: "Red",
					Weight: 0.4,
				},
			},
		},
	}
	tags := timeAndStart(order)
	assert.Contains(t, tags, "multiple-items")
	assert.Contains(t, tags, "heavy")
}

func TestSlice(t *testing.T) {
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
	assert.Contains(t, tags, "clothing")
}

func TestBool(t *testing.T) {
	order := Order{
		Email: "user1@gmail.com",
		Address: OrderAddress{
			Postcode:    "AA1 1AA",
			Residential: true,
		},
	}
	tags := timeAndStart(order)
	assert.Contains(t, tags, "residential")
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
