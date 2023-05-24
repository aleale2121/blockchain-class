package public

import "github.com/ardanlabs/blockchain/foundation/blockchain/database"

type act struct {
	Account database.AccountID `json:"account"`
	Name    string             `json:"name"`
	Balance uint64             `json:"balance"`
	Nonce   uint64             `json:"nonce"`
}
