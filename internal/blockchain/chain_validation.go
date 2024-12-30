package blockchain

func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		current := bc.Blocks[i]
		previous := bc.Blocks[i-1]

		if current.PreviousHash != previous.Hash || current.Hash != CalculateHash(current) {
			return false
		}
	}
	return true
}
