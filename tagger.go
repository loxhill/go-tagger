package tagger

import (
	"reflect"
	"strings"
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
func (t *Tagger) Tag(sample interface{}) (tags []string) {
	for _, group := range t.RuleGroups {
		s := reflect.Indirect(reflect.ValueOf(sample))
		fieldKeys := strings.Split(group.Field, ".")
		field := s.FieldByName(fieldKeys[0])
		if len(fieldKeys) < 1 {
			for i := 1; 0 <= len(fieldKeys); i++ {
				field = field.FieldByName(fieldKeys[i])
			}
		}
		tags = append(tags, getTags(group, field)...)
	}
	return tags
}
