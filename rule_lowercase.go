package main

import "sync"

type RuleLowercase struct {
	Name string
	*sync.RWMutex
}

func (rule *RuleLowercase) Init() Rule {
	rule.RWMutex = new(sync.RWMutex)

	return rule
}

func (rule *RuleLowercase) GetName() string {
	rule.Lock()
	defer rule.Unlock()

	return rule.Name
}

func (rule *RuleLowercase) Validate(value string) bool {
	return false
}