package main

// TODO ; 3awedha

// ProofOfWork :
func (bc *Blockchain) ProofOfWork() int {
	// hash(p * p' ) = ef7....00

	lastProof := bc.GetLastBlock().Proof
	return lastProof + 1
}

// 3awed kifach tzid block , w kifash t affecte proof jdid
