package models

import "time"

//registered user
type User struct {
	UserID        string    `json:"user_id"`
	Username      string    `json:"username"`
	WalletAddress string    `json:"wallet_address"`
	Email         string    `json:"email"` //Thinking if this is neccesary
	CreatedAt     time.Time `json:"created_at"`
}

// This will be the proposal
type Proposal struct {
	ProposalID  string    `json:"proposal_id"`
	AuthorID    string    `json:"author_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Status      string    `json:"status"`
}

// Vote a proposal
type Vote struct {
	VoteID     string    `json:"vote_id"`
	ProposalID string    `json:"proposal_id"`
	VoterID    string    `json:"voter_id"`
	VoteType   string    `json:"vote_type"`
	Timestamp  time.Time `json:"timestamp"`
}
