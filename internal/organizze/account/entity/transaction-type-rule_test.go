package entity

import "testing"

func TestTextRule_IsMatch(t *testing.T) {
	cases := []struct {
		name     string
		rule     TextRule
		text     string
		contains bool
		want     bool
	}{
		{
			name: "text should contain in the rule",
			rule: TextRule{Text: "rule-example", Contains: true},
			text: "example",
			want: true,
		},
		{
			name: "exact text should contain in the rule",
			rule: TextRule{Text: "rule-example", Contains: true},
			text: "rule-example",
			want: true,
		},
		{
			name: "empty text should not contain in the rule",
			rule: TextRule{Text: "rule-example", Contains: true},
			text: "",
			want: false,
		},
		{
			name: "distinct text should not contain in the rule",
			rule: TextRule{Text: "rule-example", Contains: true},
			text: "test",
			want: false,
		},
		{
			name: "text equal to rule with contains false",
			rule: TextRule{Text: "rule-example", Contains: false},
			text: "rule-example",
			want: true,
		},
		{
			name: "different text from the rule with contains false",
			rule: TextRule{Text: "rule-example", Contains: false},
			text: "example",
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.rule.IsMatch(tc.text)
			if got != tc.want {
				t.Errorf("Expected %v, got %v for text: %s", tc.want, got, tc.text)
			}
		})
	}
}
