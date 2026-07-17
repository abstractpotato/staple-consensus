package consensus

import (
  PSL "github.com/abstractpotato/potato-serialization-lib"
  Ledger "github.com/abstractpotato/starch-pay-ledger"
)

type Consensus struct {
  Ledger Ledger.Ledger
  Disk   Ledger.Disk
}

func NewConsensus() Consensus {
  return {
    Ledger: Ledger.NewLedger(),
    Disk: Ledger.NewDisk(),
  }
}

func (consensus *Consensus) AddBlock(block PSL.Block) {
  // check if disk has block
  // if not check if block is valid
    // calc the epoch
    // verify the leader has signed
    //
}

func (consensus *Consensus) UpdateWitnesses(blockID uint, witness PSL.Witness) {
  // verify cert
  // add witness if valid
}

// func (consensus *Consensus) ValidateBlock(block PSL.Block) bool {}

// func (consensus *Consensus) ValidateTx(tx PSL.Transaction) bool {}

// func (consensus *Consensus) DetermineLeader(id uint) {}
