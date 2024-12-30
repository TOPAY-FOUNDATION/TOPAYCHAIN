package common

const (
	// Tokenomics
	InitialTokenSupply uint64 = 120000000 // Total supply of tokens
	TokenDecimals      uint   = 18        // Number of decimal places for tokens

	// Blockchain Parameters
	BlockReward = 10                // Reward for mining a block
	Difficulty  = 4                 // Proof-of-Work difficulty level

	// Governance
	ProposalQuorum = 1000           // Minimum votes required for a proposal to pass
)

var (
	// Application-wide error messages
	ErrInsufficientFunds = "insufficient funds for the transaction"
	ErrInvalidProposal   = "proposal ID not found or inactive"
)
