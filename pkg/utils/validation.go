package utils

import (
	"errors"
	"regexp"
)

func ValidateAddress(address string) error {
	if len(address) != 40 {
		return errors.New("invalid address length")
	}
	match, _ := regexp.MatchString("^[a-fA-F0-9]{40}$", address)
	if !match {
		return errors.New("invalid address format")
	}
	return nil
}

func ValidateAmount(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	return nil
}
