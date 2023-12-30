package entity

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Solon97/goexpert-api/pkg/entity"
)

type OrganizzeAccount struct {
	ID          entity.ID      `json:"id"`
	Username    string         `json:"username"`
	Email       string         `json:"email"`
	BasicToken  string         `json:"basic_token"`
	InitialDate entity.ISODate `json:"initial_date"`
}

func NewOrganizzeAccount(username, email, basicToken string, initialDate entity.ISODate) (*OrganizzeAccount, error) {
	// validar conta
	isValidAccount, err := isValidAccount(username, email, basicToken)
	if err != nil {
		return nil, fmt.Errorf("account validation error: %v", err)
	}

	if !isValidAccount {
		return nil, fmt.Errorf("invalid credentials")
	}

	return &OrganizzeAccount{
		ID:          entity.NewID(),
		Username:    username,
		Email:       email,
		BasicToken:  basicToken,
		InitialDate: initialDate,
	}, nil
}
func isValidAccount(username, email, token string) (bool, error) {
	url := "https://api.organizze.com.br/rest/v2/transactions"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	req.Header.Add("User-Agent", fmt.Sprintf("%s (%s)", username, email))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}

	return res.StatusCode == http.StatusOK, nil

}
