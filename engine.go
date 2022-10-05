package tagger

import (
	"reflect"
	"strings"
)

func getTags(ruleGroup RuleGroup, fieldContent reflect.Value) (tags []string) {
	for _, rule := range ruleGroup.Rules {
		tripped := checkSingleRule(rule, fieldContent)
		if tripped {
			tags = append(tags, rule.Tag)
		}
	}
	return tags
}

func checkSingleRule(rule Rule, content reflect.Value) (tripped bool) {
	switch rule.Type {
	case "contains":
		tripped = checkContainsRule(rule.Value, content)
	case "count":
		tripped = checkCountRule(rule, content)
	default:
		tripped = false
	}
	return tripped
}

func checkCountRule(rule Rule, content reflect.Value) (tripped bool) {
	var val float64
	switch content.Kind() {
	case reflect.Slice:
		val = float64(content.Len())
	case reflect.Int:
		val = float64(content.Int())
	}
	switch rule.Op {
	case "eq":
		if val == rule.Value.(float64) {
			tripped = true
		}
	case "gteq":
		if val >= rule.Value.(float64) {
			tripped = true
		}
	}
	return tripped
}

func checkContainsRule(value interface{}, content reflect.Value) (tripped bool) {
	switch content.Kind() {
	case reflect.String:
		if strings.Contains(strings.ToLower(content.String()), strings.ToLower(value.(string))) {
			tripped = true
		}
	case reflect.Slice:
		for i := 0; i < content.Len(); i++ {
			val := content.Index(i)
			tripped = checkContainsRule(value, val)
		}
	case reflect.Struct:
		for i := 0; i < content.NumField(); i++ {
			tripped = checkContainsRule(value, content.Field(0))
		}
	}
	return tripped
}
