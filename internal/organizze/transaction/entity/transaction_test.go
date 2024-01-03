package entity

import (
	"testing"
	"time"

	"github.com/Solon97/goexpert-api/internal/organizze/account/entity"
)

const (
	rule = "Regra"
)

func TestGetDate(t *testing.T) {
	cases := []struct {
		name        string
		transaction OrganizzeTransaction
		want        time.Time
	}{
		{
			name: "Standard Date",
			transaction: OrganizzeTransaction{
				Date: "2024-01-01",
			},
			want: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "Credit Card Installment Expense Date",
			transaction: OrganizzeTransaction{
				Date:              "2024-01-01",
				CreditCardID:      &[]int{1}[0],
				InvoiceID:         &[]int{1}[0],
				TotalInstallments: 5,
				Invoice: &CreditCardInvoice{
					Date: "2024-02-01",
				},
				IntegrationAccount: entity.OrganizzeAccount{
					InstallmentRule: entity.TransactionTypeRule{
						TagNames: []entity.TextRule{
							{
								Text:     rule,
								Contains: true,
							},
						},
					},
				},
				Tags: []Tag{{Name: rule}},
			},
			want: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "Credit Card Unique Output Date",
			transaction: OrganizzeTransaction{
				Date:         "2024-01-01",
				CreditCardID: &[]int{1}[0],
				InvoiceID:    &[]int{1}[0],
				Invoice: &CreditCardInvoice{
					Date: "2024-02-01",
				},
			},
			want: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.transaction.GetDate()
			if !got.Equal(tc.want) {
				t.Errorf("GetDate() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsInput(t *testing.T) {
	cases := []struct {
		name        string
		transaction OrganizzeTransaction
		want        bool
	}{
		{
			name: "Is Input",
			transaction: OrganizzeTransaction{
				AmountCents: 100,
			},
			want: true,
		},
		{
			name: "Is Output",
			transaction: OrganizzeTransaction{
				AmountCents: -100,
			},
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.transaction.IsInput()
			if got != tc.want {
				t.Errorf("IsInput() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsInstallment(t *testing.T) {
	cases := []struct {
		name        string
		transaction OrganizzeTransaction
		want        bool
	}{
		{
			name: "No Installment",
			transaction: OrganizzeTransaction{
				IntegrationAccount: entity.OrganizzeAccount{
					InstallmentRule: entity.TransactionTypeRule{
						CategoryNames: []entity.TextRule{
							{
								Text: rule,
							},
						},
					},
				},
				Category: Category{Name: "Teste"}},
			want: false,
		},
		{
			name: "Has Installment Tag",
			transaction: OrganizzeTransaction{
				IntegrationAccount: entity.OrganizzeAccount{
					InstallmentRule: entity.TransactionTypeRule{
						TagNames: []entity.TextRule{
							{
								Text: rule,
							},
						},
					},
				},
				Tags: []Tag{{Name: rule}}},
			want: true,
		},
		{
			name: "Has Installment Category",
			transaction: OrganizzeTransaction{
				IntegrationAccount: entity.OrganizzeAccount{
					InstallmentRule: entity.TransactionTypeRule{
						CategoryNames: []entity.TextRule{
							{
								Text: rule,
							},
						},
					},
				},
				Category: Category{Name: rule}},
			want: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.transaction.IsInstallment()
			if got != tc.want {
				t.Errorf("IsInstallment() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsCommitment(t *testing.T) {
	cases := []struct {
		name        string
		transaction OrganizzeTransaction
		want        bool
	}{
		{
			name: "Is Commitment",
			transaction: OrganizzeTransaction{
				IntegrationAccount: entity.OrganizzeAccount{
					CommitmentRule: entity.TransactionTypeRule{
						TagNames: []entity.TextRule{
							{
								Text: rule,
							},
						},
					},
				},
				Tags: []Tag{{Name: rule}},
			},
			want: true,
		},
		{
			name: "No Commitment",
			transaction: OrganizzeTransaction{
				IntegrationAccount: entity.OrganizzeAccount{
					CommitmentRule: entity.TransactionTypeRule{
						TagNames: []entity.TextRule{
							{
								Text:     rule,
								Contains: false,
							},
						},
					},
				},
				Tags: []Tag{{Name: "Teste"}},
			},
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.transaction.IsCommitment()
			if got != tc.want {
				t.Errorf("IsCommitment() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsInvestment(t *testing.T) {
	cases := []struct {
		name        string
		transaction OrganizzeTransaction
		want        bool
	}{
		{
			name: "Is Investment",
			transaction: OrganizzeTransaction{
				IntegrationAccount: entity.OrganizzeAccount{
					InvestmentRule: entity.TransactionTypeRule{
						TagNames: []entity.TextRule{
							{
								Text: rule,
							},
						},
					},
				},
				Tags: []Tag{{Name: rule}},
			},
			want: true,
		},
		{
			name: "No Investment",
			transaction: OrganizzeTransaction{
				IntegrationAccount: entity.OrganizzeAccount{
					InvestmentRule: entity.TransactionTypeRule{
						TagNames: []entity.TextRule{
							{
								Text:     rule,
								Contains: false,
							},
						},
					},
				},
				Tags: []Tag{{Name: "Teste"}},
			},
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.transaction.IsInvestment()
			if got != tc.want {
				t.Errorf("IsInvestment() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsInvestmentRedemption(t *testing.T) {
	cases := []struct {
		name        string
		transaction OrganizzeTransaction
		want        bool
	}{
		{
			name: "Is InvestmentRedemption",
			transaction: OrganizzeTransaction{
				IntegrationAccount: entity.OrganizzeAccount{
					InvestmentRedemptionRule: entity.TransactionTypeRule{
						TagNames: []entity.TextRule{
							{
								Text: rule,
							},
						},
					},
				},
				Tags: []Tag{{Name: rule}},
			},
			want: true,
		},
		{
			name: "No InvestmentRedemption",
			transaction: OrganizzeTransaction{
				IntegrationAccount: entity.OrganizzeAccount{
					InvestmentRedemptionRule: entity.TransactionTypeRule{
						TagNames: []entity.TextRule{
							{
								Text:     rule,
								Contains: false,
							},
						},
					},
				},
				Tags: []Tag{{Name: "Teste"}},
			},
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.transaction.IsInvestmentRedemption()
			if got != tc.want {
				t.Errorf("IsInvestmentRedemption() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsRefund(t *testing.T) {
	cases := []struct {
		name        string
		transaction OrganizzeTransaction
		want        bool
	}{
		{
			name: "Is Refund",
			transaction: OrganizzeTransaction{
				IntegrationAccount: entity.OrganizzeAccount{
					RefundRule: entity.TransactionTypeRule{
						TagNames: []entity.TextRule{
							{
								Text: rule,
							},
						},
					},
				},
				Tags: []Tag{{Name: rule}},
			},
			want: true,
		},
		{
			name: "No Refund",
			transaction: OrganizzeTransaction{
				IntegrationAccount: entity.OrganizzeAccount{
					RefundRule: entity.TransactionTypeRule{
						TagNames: []entity.TextRule{
							{
								Text:     rule,
								Contains: false,
							},
						},
					},
				},
				Tags: []Tag{{Name: "Teste"}},
			},
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.transaction.IsRefund()
			if got != tc.want {
				t.Errorf("IsRefund() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsRefundable(t *testing.T) {
	cases := []struct {
		name        string
		transaction OrganizzeTransaction
		want        bool
	}{
		{
			name: "Is Refundable",
			transaction: OrganizzeTransaction{
				IntegrationAccount: entity.OrganizzeAccount{
					RefundableRule: entity.TransactionTypeRule{
						TagNames: []entity.TextRule{
							{
								Text: rule,
							},
						},
					},
				},
				Tags: []Tag{{Name: rule}},
			},
			want: true,
		},
		{
			name: "No Refundable",
			transaction: OrganizzeTransaction{
				IntegrationAccount: entity.OrganizzeAccount{
					RefundableRule: entity.TransactionTypeRule{
						TagNames: []entity.TextRule{
							{
								Text: rule,
							},
						},
					},
				},
				Tags: []Tag{{Name: "teste"}},
			},
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.transaction.IsRefundable()
			if got != tc.want {
				t.Errorf("IsRefund() = %v, want %v", got, tc.want)
			}
		})
	}
}
func TestToIgnore(t *testing.T) {
	cases := []struct {
		name        string
		transaction OrganizzeTransaction
		want        bool
	}{
		{
			name: "To Ignore",
			transaction: OrganizzeTransaction{
				IntegrationAccount: entity.OrganizzeAccount{
					IgnoreTransactionRule: entity.TransactionTypeRule{
						TagNames: []entity.TextRule{
							{
								Text: rule,
							},
						},
					},
				},
				Tags: []Tag{{Name: rule}},
			},
			want: true,
		},
		{
			name: "No Ignore",
			transaction: OrganizzeTransaction{
				IntegrationAccount: entity.OrganizzeAccount{
					IgnoreTransactionRule: entity.TransactionTypeRule{
						TagNames: []entity.TextRule{
							{
								Text: rule,
							},
						},
					},
				},
				Tags: []Tag{{Name: "teste"}},
			},
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.transaction.ToIgnore()
			if got != tc.want {
				t.Errorf("IsRefund() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsCreditCardExpense(t *testing.T) {
	creditCardID := 1
	invoiceID := 1

	cases := []struct {
		name        string
		transaction OrganizzeTransaction
		want        bool
	}{
		{
			name:        "Credit Card Expense",
			transaction: OrganizzeTransaction{CreditCardID: &creditCardID, InvoiceID: &invoiceID},
			want:        true,
		},
		{
			name:        "No Credit Card Expense",
			transaction: OrganizzeTransaction{},
			want:        false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.transaction.IsCreditCardExpense()
			if got != tc.want {
				t.Errorf("IsCreditCardExpense() = %v, want %v", got, tc.want)
			}
		})
	}
}
