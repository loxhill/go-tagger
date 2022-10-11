package tagger

import (
	"fmt"
	"reflect"
)

type Tagger struct {
	RuleGroups []RuleGroup
}

func NewTagger() *Tagger {
	return &Tagger{}
}

// LoadRules will load one JSON rule files to be used for tagging.
func (t *Tagger) LoadRules(path string) {
	var ruleGroup []RuleGroup
	ruleGroup, err := parseRuleFile(path)
	if err != nil {
		panic(err)
	}
	t.RuleGroups = ruleGroup
}

// Tag will begin tagging the provided sample based on the rules loaded with LoadRules.
// https://www.freecodecamp.org/news/iteration-in-golang/
func (t *Tagger) Tag(sample interface{}) (tags []string) {
	s := reflect.ValueOf(sample)
	types := s.Type()
	for i := 0; i < s.NumField(); i++ {
		e := Engine{Rules: t.RuleGroups}
		tags = append(tags, e.parseField(types.Field(i).Name, s.Field(i))...)
	}
	fmt.Println(tags)
	return tags
}
