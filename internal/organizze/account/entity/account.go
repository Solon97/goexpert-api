package entity

import (
	"fmt"

	"github.com/Solon97/goexpert-api/pkg/common"
)

type OrganizzeAccountValidator interface {
	IsValidAccount(username, email, token string) (bool, error)
}

type OrganizzeAccount struct {
	ID                       common.ID           `json:"id"`
	Username                 string              `json:"username"`
	Email                    string              `json:"email"`
	BasicToken               string              `json:"basic_token"`
	InitialDate              common.ISODate      `json:"initial_date"`
	InstallmentRule          TransactionTypeRule `json:"installment_rule"`
	CommitmentRule           TransactionTypeRule `json:"commitment_rule"`
	RefundRule               TransactionTypeRule `json:"refund_rule"`
	RefundableRule           TransactionTypeRule `json:"refundable_rule"`
	InvestmentRule           TransactionTypeRule `json:"investment_rule"`
	InvestmentRedemptionRule TransactionTypeRule `json:"investment_redemption_rule"`
	IgnoreTransactionRule    TransactionTypeRule `json:"ignore_transaction_rule"`
}

func NewOrganizzeAccount(username, email, basicToken string, initialDate common.ISODate, accountValidator OrganizzeAccountValidator) (*OrganizzeAccount, error) {
	// validar conta
	isValidAccount, err := accountValidator.IsValidAccount(username, email, basicToken)
	if err != nil {
		return nil, fmt.Errorf("account validation error: %v", err)
	}

	if !isValidAccount {
		return nil, fmt.Errorf("invalid credentials")
	}

	return &OrganizzeAccount{
		ID:          common.NewID(),
		Username:    username,
		Email:       email,
		BasicToken:  basicToken,
		InitialDate: initialDate,
	}, nil
}

// func isValidAccount(username, email, token string) (bool, error) {
// 	url := "https://api.organizze.com.br/rest/v2/transactions"

// 	req, err := http.NewRequest(http.MethodGet, url, nil)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	req.Header.Add("User-Agent", fmt.Sprintf("%s (%s)", username, email))
// 	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", token))

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return false, err
// 	}

// 	return res.StatusCode == http.StatusOK, nil

// }
