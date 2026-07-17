package main

import (
  "os"
  "fmt"
  "time"
  PSL "github.com/abstractpotato/potato-serialization-lib/psl"
  Ledger "github.com/abstractpotato/starch-pay-ledger/ledger"
)

func loadPrivateKey() ([]byte, error) {
  privateKey, err := os.ReadFile(".env/skey")
  if err != nil { return nil, err }
  return privateKey[:96], nil
}

func main() {
  privateKey, err := loadPrivateKey()
  if err != nil { panic(err) }

  block := generate_genesis_block(privateKey)
  blockJSON, err := block.ToJSON()
  fmt.Printf("%s\n\n", blockJSON)

  blockCBOR, err := block.ToCBOR()
  if err != nil { panic(err) }

  disk := Ledger.NewDisk()
  disk.CreatedDirs()
  err = disk.SaveBlockCBOR(0, blockCBOR)
  if err != nil { panic(err) }

  ledger := Ledger.NewLedger()
  ledger.Genesis = block.Body.Genesis
  ledger.InitTime = block.Body.Timestamp
  ledger.Params = &block.Body.Genesis.Params
  ledger.AddCertificate(block.Body.Genesis.Certificate)

  ledgerJSON, err := ledger.ToJSON()
  if err != nil { panic(err) }

  fmt.Printf("ledger: %s\n", ledgerJSON)
}

func generate_genesis_block(privateKey []byte) PSL.Block {

  params := PSL.NewParams()
  params.Network = 0
  params.MaxBlockHeaderSize = 1100 // 128 bytes
  params.MaxBlockBodySize = 4000000 // 4 MB or ~15k simple transactions
  params.MaxTxSize = 4000 // 4 KB
  params.TxFeePerByte = 430
  params.MinTxFee = params.TxFeePerByte * 175 // signature size
  params.SlotsPerEpoch = 432000
  params.SlotTimeInMs = 1000
  params.ProtocolVersion = 0

  cert := PSL.NewCertificate()
  cert.RequestTx = "genesis"
  cert.RewardAddr = "genesis"
  cert.AddRelay("0.0.0.0:5001")
  cert.AddRelay("0.0.0.0:5002")
  cert.Status = 1

  genesis := PSL.Genesis{}
  genesis.Seed = []byte("bonepool")
  genesis.Certificate = cert
  genesis.Params = params

  block := PSL.NewBlock()
  block.Body.Genesis = &genesis
  block.Body.Timestamp = uint(time.Now().UnixMilli())
  block.Hash()

  err := block.Sign(privateKey)
  if err != nil { panic(err) }

  return block
}
