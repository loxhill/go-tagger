package tagger

import (
	"encoding/json"
	"io/ioutil"
)

type RuleGroup struct {
	Field string
	Rules []Rule
}

type Rule struct {
	Field string
	Type  string
	Op    string
	Value interface{}
	Tag   string
}

func parseRuleFile(path string) (ruleGroup []RuleGroup, err error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &ruleGroup)
	if err != nil {
		return nil, err
	}
	return ruleGroup, nil
}
