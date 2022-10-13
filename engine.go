package tagger

import (
	"reflect"
	"strings"
)

type Engine struct {
	Rules []RuleGroup
	Tags  []string
}

func (e *Engine) parseField(field string, data reflect.Value) (tags []string) {
	switch data.Kind() {
	case reflect.Slice:
		tags = append(tags, e.scanElement(field, data)...)
		for i := 0; i < data.Len(); i++ {
			tags = e.parseField(field, data.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < data.NumField(); i++ {
			title := field + "." + data.Type().Field(i).Name
			tags = e.parseField(title, data.Field(i))
		}
	case reflect.String, reflect.Int, reflect.Float64, reflect.Float32, reflect.Bool:
		tags = append(tags, e.scanElement(field, data)...)
	}
	return e.Tags
}

// scanElement is used to scan data once we've got past the layers within the struct/slice.
func (e *Engine) scanElement(name string, data reflect.Value) (tags []string) {
	ruleGroup := e.findRuleGroupForField(name)
	for _, rule := range ruleGroup.Rules {
		processedTags := e.processRule(rule, data)
		if processedTags != "" {
			e.Tags = append(e.Tags, processedTags)
		}
	}
	return tags
}

func (e *Engine) findRuleGroupForField(name string) RuleGroup {
	for _, r := range e.Rules {
		if r.Field == name {
			return r
		}
	}
	return RuleGroup{}
}

func (e *Engine) processRule(rule Rule, data reflect.Value) string {
	tripped := false
	switch rule.Type {
	case "contains":
		tripped = e.checkContainsRule(rule.Value, data)
	case "count":
		tripped = e.checkCountRule(rule, data)
	case "bool":
		tripped = e.checkBoolRule(rule.Value, data)
	default:
		tripped = false
	}
	if tripped {
		return rule.Tag
	}
	return ""
}

func (e *Engine) checkContainsRule(value interface{}, data reflect.Value) (tripped bool) {
	switch data.Kind() {
	case reflect.String:
		if strings.Contains(strings.ToLower(data.String()), strings.ToLower(value.(string))) {
			tripped = true
		}
	}
	return tripped
}

func (e *Engine) checkCountRule(rule Rule, data reflect.Value) (tripped bool) {
	var val float64
	switch data.Kind() {
	case reflect.Slice:
		val = float64(data.Len())
	case reflect.Int:
		val = float64(data.Int())
	case reflect.Float64, reflect.Float32:
		val = data.Float()
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
	case "gt":
		if val > rule.Value.(float64) {
			tripped = true
		}
	case "lteq":
		if val <= rule.Value.(float64) {
			tripped = true
		}
	case "lt":
		if val < rule.Value.(float64) {
			tripped = true
		}
	}
	return tripped
}

func (e *Engine) checkBoolRule(value interface{}, data reflect.Value) (tripped bool) {
	switch data.Kind() {
	case reflect.Bool:
		if value.(bool) == data.Bool() {
			tripped = true
		}
	}
	return tripped
}
