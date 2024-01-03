package entity

import (
	"time"

	"github.com/Solon97/goexpert-api/internal/organizze/account/entity"
)

type IOrganizzeTransaction struct {
	OrganizzeTransaction
}

type OrganizzeTransaction struct {
	ID                 int64  `json:"id"`
	Description        string `json:"description"`
	Date               string `json:"date"`
	Paid               bool   `json:"paid"`
	AmountCents        int64  `json:"amount_cents"`
	TotalInstallments  int    `json:"total_installments"`
	Installment        int    `json:"installment"`
	Recurring          bool   `json:"recurring"`
	CreditCardID       *int   `json:"credit_card_id"`
	InvoiceID          *int   `json:"credit_card_invoice_id"`
	Tags               []Tag  `json:"tags"`
	CategoryID         int    `json:"category_id"`
	Category           Category
	Invoice            *CreditCardInvoice
	IntegrationAccount entity.OrganizzeAccount
}

func (transaction OrganizzeTransaction) GetDate() time.Time {
	if transaction.IsCreditCardExpense() && transaction.IsInstallment() {
		date, _ := time.Parse("2006-01-02", transaction.Invoice.Date)
		return date
	}

	date, _ := time.Parse("2006-01-02", transaction.Date)
	return date
}

func (transaction OrganizzeTransaction) IsInput() bool {
	return transaction.AmountCents > 0
}

func (transaction OrganizzeTransaction) IsRefund() bool {
	return transaction.validateTransactionRule(transaction.IntegrationAccount.RefundRule)
}
func (transaction OrganizzeTransaction) IsRefundable() bool {
	return transaction.validateTransactionRule(transaction.IntegrationAccount.RefundableRule)
}

func (transaction OrganizzeTransaction) IsInstallment() bool {
	return transaction.validateTransactionRule(transaction.IntegrationAccount.InstallmentRule)
}

func (transaction OrganizzeTransaction) IsCommitment() bool {
	return transaction.validateTransactionRule(transaction.IntegrationAccount.CommitmentRule)
}

func (transaction OrganizzeTransaction) IsInvestment() bool {
	return transaction.validateTransactionRule(transaction.IntegrationAccount.InvestmentRule)
}

func (transaction OrganizzeTransaction) IsInvestmentRedemption() bool {
	return transaction.validateTransactionRule(transaction.IntegrationAccount.InvestmentRedemptionRule)
}

func (transaction OrganizzeTransaction) ToIgnore() bool {
	return transaction.validateTransactionRule(transaction.IntegrationAccount.IgnoreTransactionRule)
}

func (transaction OrganizzeTransaction) IsCreditCardExpense() bool {
	return transaction.CreditCardID != nil && transaction.InvoiceID != nil
}

func (transaction OrganizzeTransaction) validateTransactionRule(rule entity.TransactionTypeRule) bool {
	//? Validar Categorias
	for _, categoryRule := range rule.CategoryNames {
		if categoryRule.IsMatch(transaction.Category.Name) {
			return true
		}
	}
	//? Validar Tags
	for _, tagRule := range rule.TagNames {
		for _, transactionTag := range transaction.Tags {
			if tagRule.IsMatch(transactionTag.Name) {
				return true
			}
		}
	}

	return false
}
