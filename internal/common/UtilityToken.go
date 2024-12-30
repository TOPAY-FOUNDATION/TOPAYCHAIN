package common

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/google/uuid"
)

type UtilityToken struct {
	Name         string
	Symbol       string
	TotalSupply  *big.Int
	Decimals     uint
	Balances     map[string]*big.Int
	VotingPower  map[string]*big.Int
	Proposals    []*Proposal
	Address      string
}

type Proposal struct {
	ID          string
	Title       string
	Description string
	Votes       map[string]*big.Int
	YesVotes    *big.Int
	NoVotes     *big.Int
	Status      string
}

// NewUtilityToken initializes a new token with a unique address.
func NewUtilityToken(name, symbol string, supply uint64, decimals uint) *UtilityToken {
	totalSupply := new(big.Int).Mul(big.NewInt(int64(supply)), new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
	address := generateUniqueAddress()

	return &UtilityToken{
		Name:        name,
		Symbol:      symbol,
		TotalSupply: totalSupply,
		Decimals:    decimals,
		Balances:    make(map[string]*big.Int),
		VotingPower: make(map[string]*big.Int),
		Proposals:   []*Proposal{},
		Address:     address,
	}
}

// Transfer transfers tokens from the sender to the receiver.
func (token *UtilityToken) Transfer(sender, receiver string, amount *big.Int) error {
	if senderBalance, exists := token.Balances[sender]; !exists || senderBalance.Cmp(amount) < 0 {
		return fmt.Errorf("insufficient balance")
	}

	token.Balances[sender].Sub(token.Balances[sender], amount)

	if _, exists := token.Balances[receiver]; !exists {
		token.Balances[receiver] = big.NewInt(0)
	}
	token.Balances[receiver].Add(token.Balances[receiver], amount)

	return nil
}

// AddProposal creates a new governance proposal.
func (token *UtilityToken) AddProposal(title, description string) *Proposal {
	proposal := &Proposal{
		ID:          generateProposalID(),
		Title:       title,
		Description: description,
		Votes:       make(map[string]*big.Int),
		YesVotes:    big.NewInt(0),
		NoVotes:     big.NewInt(0),
		Status:      "Active",
	}
	token.Proposals = append(token.Proposals, proposal)
	return proposal
}

// Vote allows a user to vote on a proposal.
func (token *UtilityToken) Vote(proposalID, voterAddress string, voteYes bool) error {
	var proposal *Proposal
	for _, p := range token.Proposals {
		if p.ID == proposalID {
			proposal = p
			break
		}
	}
	if proposal == nil {
		return fmt.Errorf("proposal not found")
	}
	if proposal.Status != "Active" {
		return fmt.Errorf("proposal is not active")
	}

	votingPower := token.VotingPower[voterAddress]
	if votingPower == nil || votingPower.Sign() == 0 {
		return fmt.Errorf("no voting power")
	}

	if voteYes {
		proposal.YesVotes.Add(proposal.YesVotes, votingPower)
	} else {
		proposal.NoVotes.Add(proposal.NoVotes, votingPower)
	}

	proposal.Votes[voterAddress] = votingPower
	return nil
}

// CloseProposal finalizes a proposal.
func (token *UtilityToken) CloseProposal(proposalID string) error {
	var proposal *Proposal
	for _, p := range token.Proposals {
		if p.ID == proposalID {
			proposal = p
			break
		}
	}
	if proposal == nil {
		return fmt.Errorf("proposal not found")
	}

	if proposal.YesVotes.Cmp(proposal.NoVotes) > 0 {
		proposal.Status = "Passed"
	} else {
		proposal.Status = "Rejected"
	}
	return nil
}

func generateUniqueAddress() string {
	id := uuid.New().String()
	hash := sha256.Sum256([]byte(id))
	return hex.EncodeToString(hash[:])[:40]
}

func generateProposalID() string {
	return uuid.New().String()
}
