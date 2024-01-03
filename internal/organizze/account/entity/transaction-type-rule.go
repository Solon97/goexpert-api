package entity

import "strings"

type TransactionTypeRule struct {
	TagNames      []TextRule `json:"tags"`
	CategoryNames []TextRule `json:"categories"`
}

type TextRule struct {
	Text     string `json:"text"`
	Contains bool   `json:"contains"`
}

func (tr TextRule) IsMatch(text string) bool {
	if text == "" {
		return false
	}
	if tr.Contains {
		return strings.Contains(tr.Text, text)
	}

	return strings.EqualFold(tr.Text, text)
}
