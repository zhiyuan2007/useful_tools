package main

import (
	//"bytes"
	"encoding/hex"
	"errors"
	"flag"
	"os/user"
	"path/filepath"
	//"errors"
	"fmt"
	"log"
	"sort"

	"github.com/boltdb/bolt"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/visor"
	//"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/cipher/encoder"
	"github.com/skycoin/skycoin/src/coin"
	"github.com/skycoin/skycoin/src/visor/bucket"
	"github.com/skycoin/skycoin/src/visor/historydb"
)

type BlockInfo struct {
	headHash string
	bc       coin.Block
}

//address_in
//address_txns
//blocks_sigs
//block_tree
//blocks
//history_meta
//transactions
//unconfirmed_txns
//unconfirmed_unspents
//unspent_pool
//unspent_meta
//blockchain_meta
//uxouts

func printHelps() {
	fmt.Println("buckt name as follows:")
	fmt.Println("\t address_in")
	fmt.Println("\t address_txns")
	fmt.Println("\t blocks_sigs")
	fmt.Println("\t block_tree")
	fmt.Println("\t blocks")
	fmt.Println("\t history_meta")
	fmt.Println("\t transactions")
	fmt.Println("\t unconfirmed_txns")
	fmt.Println("\t unconfirmed_unspents")
	fmt.Println("\t unspent_pool")
	fmt.Println("\t unspent_meta")
	fmt.Println("\t blockchain_meta")
	fmt.Println("\t uxouts")
}
func main() {
	dbname := flag.String("db", "", "bucket name")
	bkt := flag.String("bkt", "", "bucket name")
	flag.Parse()
	if *dbname == "" {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(usr.HomeDir)
		*dbname = filepath.Join(usr.HomeDir, ".spo/data.db1")
	}
	if *bkt == "" {
		printHelps()
		flag.PrintDefaults()
		log.Fatal("bucket required")
	}
	opt := &bolt.Options{ReadOnly: true}
	db, err := bolt.Open(*dbname, 0400, opt)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	switch *bkt {
	case "uxouts":
		getUxouts(db)
	case "blocks":
		getBlockchain(db)
	case "block_tree":
		getBlocktree(db)
	case "blocks_sigs":
		getBlocksig(db)
	case "transactions":
		getTransction(db)
	case "unspent_pool":
		getUnspentPool(db)
	case "unconfirmed_txns":
		getUnconfirmedTxns(db)
	case "address_in":
		getAddressIn(db)
	case "address_txns":
		getAddressTxns(db)
	case "unconfirmed_unspents":
		getUnconfirmedUnspents(db)
	case "history_meta":
		getHistoryMeta(db)
	case "blockchain_meta":
		getBlockChianMeta(db)
	case "unspent_meta":
		getUnspentMeta(db)
	default:
		printHelps()
	}
}

func getBlockchain(db *bolt.DB) {

	bi := []BlockInfo{}
	bucketname := "blocks"

	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(bucketname))
		fmt.Printf("bucket-name: %s\n", bucketname)

		b.ForEach(func(k, v []byte) error {
			block := coin.Block{}

			if err := encoder.DeserializeRaw(v, &block); err != nil {
				return nil
			}

			bi = append(bi, BlockInfo{hex.EncodeToString(k), block})
			return nil
		})
		return nil
	})

	sort.Slice(bi, func(i, j int) bool { return bi[i].bc.Head.BkSeq < bi[j].bc.Head.BkSeq })
	//dst := "TE4Ti1Drvfm274caXBHAA7yct6Vym9cNdg"
	for k, v := range bi {
		fmt.Println("-------------------------------------------")
		fmt.Println("k:", k)
		fmt.Println("head hash:", v.headHash)
		block := v.bc
		fmt.Printf("block: %+v\n", block)
		ts := block.Body.Transactions

		for idx, tx1 := range ts {

			fmt.Printf("tx index %d\n", idx)
			format_transaction(tx1)
		}
		fmt.Println("==============================================")

	}

}
func format_transaction(tx1 coin.Transaction) {
	fmt.Printf("tx Length: %+v\n", tx1.Length)
	fmt.Printf("tx Type: %+v\n", tx1.Type)
	fmt.Printf("tx InnerHash: %s\n", hex.EncodeToString(tx1.InnerHash[:]))
	for ii, txSig := range tx1.Sigs {
		fmt.Printf("tx Sigs idx %d,  %s\n", ii, hex.EncodeToString(txSig[:]))
	}
	for ii, txIn := range tx1.In {
		fmt.Printf("tx In idx %d, %+v\n", ii, hex.EncodeToString(txIn[:]))
	}
	for ii, txOut := range tx1.Out {
		fmt.Printf("tx Out idx %d, %+v\n", ii, txOut)
	}
}

func getBlocktree(db *bolt.DB) {
	bucketname := "block_tree"

	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(bucketname))
		fmt.Printf("bucket-name: %s\n", bucketname)

		b.ForEach(func(k, v []byte) error {
			key := bucket.Btoi(k)

			hps := []coin.HashPair{}
			if err := encoder.DeserializeRaw(v, &hps); err != nil {
				return nil
			}
			for i, hp := range hps {
				_hash := hex.EncodeToString(hp.Hash[:])
				_prehash := hex.EncodeToString(hp.PreHash[:])
				fmt.Printf("keyindex=%d,idx=%d,hash=%s,prevhash=%s\n", key, i, _hash, _prehash)
			}

			return nil
		})
		return nil
	})

}

func getTransction(db *bolt.DB) {
	bucketname := "transactions"

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketname))
		fmt.Printf("bucket-name: %s\n", bucketname)

		b.ForEach(func(k, v []byte) error {
			tx, _ := coin.TransactionDeserialize(v)
			fmt.Printf("k=%+v\n", hex.EncodeToString(k))
			format_transaction(tx)

			return nil
		})
		return nil
	})

}

func getBlocksig(db *bolt.DB) {
	bucketname := "block_sigs"

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketname))
		fmt.Printf("bucket-name: %s\n", bucketname)

		b.ForEach(func(k, v []byte) error {
			fmt.Printf("k=%+v\n", hex.EncodeToString(k))
			fmt.Printf("v=%+v\n", hex.EncodeToString(v))

			return nil
		})
		return nil
	})

}

//"unspent_pool"
//"unconfirmed_txns"

func getUxouts(db *bolt.DB) {
	bucketname := "uxouts"

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketname))
		fmt.Printf("bucket-name: %s\n", bucketname)

		b.ForEach(func(k, v []byte) error {
			//fmt.Printf("k=%+v, v=%+v\n", k, v)
			var out historydb.UxOut
			if err := encoder.DeserializeRaw(v, &out); err != nil {
				return err
			}

			fmt.Printf("k=%+v, v=%+v\n", hex.EncodeToString(k[:]), &out)
			fmt.Printf("k=%+v, v=%+v, spend block seq=%+v\n", out.Out.Body.Address, out.Out.Body.Coins, out.SpentBlockSeq)

			return nil
		})
		return nil
	})

}

func getUnspentPool(db *bolt.DB) {
	bucketname := "unspent_pool"

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketname))
		fmt.Printf("bucket-name: %s\n", bucketname)

		b.ForEach(func(k, v []byte) error {
			//fmt.Printf("k=%+v, v=%+v\n", k, v)
			var out coin.UxOut
			if err := encoder.DeserializeRaw(v, &out); err != nil {
				return err
			}

			fmt.Printf("k=%+v, v=%+v\n", hex.EncodeToString(k[:]), &out)
			fmt.Printf("k=%+v, v=%+v\n", hex.EncodeToString(k[:]), out.Body.SrcTransaction.Hex())

			return nil
		})
		return nil
	})

}

func getUnconfirmedTxns(db *bolt.DB) {
	bucketname := "unconfirmed_txns"

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketname))
		fmt.Printf("bucket-name: %s\n", bucketname)

		b.ForEach(func(k, v []byte) error {
			var tx visor.UnconfirmedTxn
			if err := encoder.DeserializeRaw(v, &tx); err != nil {
				return err
			}
			//fmt.Printf("k=%+v\n", cipher.MustSHA256FromHex(string(k)).Hex())
			//fmt.Printf("k=%+v\n", string(k))
			fmt.Printf("k=%+v, v=%+v\n", tx.Hash().Hex(), tx)
			fmt.Printf("txid %+v\n", tx.Txn.TxIDHex())
			fmt.Printf("innerhash %+v\n", tx.Txn.InnerHash.Hex())
			for i, _ := range tx.Txn.In {
				fmt.Printf("in-uxid %d: %+v\n", i, tx.Txn.In[i].Hex())
			}
			for i, _ := range tx.Txn.Out {
				fmt.Printf("out-uxinfo %d: %+v\n", i, tx.Txn.Out[i])
			}

			return nil
		})
		return nil
	})

}

func getAddressIn(db *bolt.DB) {
	//k=2myvxs6HMEndBctWk1XxTU48YGULYk9DhUu, v=59d2c7a06f5e4adac983073eca32721e0c94ea63623e4c97b80f96ff97df57bb
	//k=2myvxs6HMEndBctWk1XxTU48YGULYk9DhUu, v=0a9ea68e562e0d3fcc205ef295b90f46c0c4083727ff231c24b90b3629e8c0a8
	bucketname := "address_in"

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketname))
		fmt.Printf("bucket-name: %s\n", bucketname)

		b.ForEach(func(k, v []byte) error {
			var hashes []cipher.SHA256
			if err := encoder.DeserializeRaw(v, &hashes); err != nil {
				return err
			}
			addr, err := AddressFromBytes(k)
			if err != nil {
				fmt.Printf("----err--%+v\n", err)
				return err
			}
			for _, hs := range hashes {
				fmt.Printf("k=%+v, v=%+v\n", addr.String(), hs.Hex())
			}

			return nil
		})
		return nil
	})

}

func getAddressTxns(db *bolt.DB) {
	//k=2myvxs6HMEndBctWk1XxTU48YGULYk9DhUu, v=05000000295781453cfec3201370dd17e075797c2d2e969963155baa18d49cef
	//k=2mzf6XfsJaLG5FMq6Y9ELx6yzi99YTDKaFw, v=010000003dc4e4f25ee91c7bfce524e6086c03d50ba1d6214649cb3fe15c6382
	bucketname := "address_txns"

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketname))
		fmt.Printf("bucket-name: %s\n", bucketname)

		b.ForEach(func(k, v []byte) error {
			var hashes cipher.SHA256
			if err := encoder.DeserializeRaw(v, &hashes); err != nil {
				return err
			}
			addr, err := AddressFromBytes(k)
			if err != nil {
				fmt.Printf("----err--%+v\n", err)
				return err
			}
			fmt.Printf("k=%+v, v=%+v\n", addr.String(), hashes.Hex())

			return nil
		})
		return nil
	})

}

func getBlockChianMeta(db *bolt.DB) {
	bucketname := "blockchain_meta"

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketname))
		fmt.Printf("bucket-name: %s\n", bucketname)

		b.ForEach(func(k, v []byte) error {
			fmt.Printf("k=%+v, v=%+v\n", string(k), bucket.Btoi(v))

			return nil
		})
		return nil
	})

}

func getUnspentMeta(db *bolt.DB) {
	bucketname := "unspent_meta"

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketname))
		fmt.Printf("bucket-name: %s\n", bucketname)

		b.ForEach(func(k, v []byte) error {
			fmt.Printf("k=%+v, v=%+v\n", string(k), bucket.Btoi(v))

			return nil
		})
		return nil
	})
}

func getHistoryMeta(db *bolt.DB) {
	bucketname := "history_meta"

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketname))
		fmt.Printf("bucket-name: %s\n", bucketname)

		b.ForEach(func(k, v []byte) error {
			fmt.Printf("k=%+v, v=%+v\n", string(k), bucket.Btoi(v))

			return nil
		})
		return nil
	})

}

func getUnconfirmedUnspents(db *bolt.DB) {
	bucketname := "unconfirmed_unspents"

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketname))
		fmt.Printf("bucket-name: %s\n", bucketname)

		b.ForEach(func(k, v []byte) error {
			var uxArray coin.UxArray
			if err := encoder.DeserializeRaw(v, &uxArray); err != nil {
				fmt.Printf("eer %+v\n", err)
				return err
			}
			fmt.Printf("k=%+v, v=%+v\n", string(k), uxArray)
			fmt.Printf("k=%+v, src_txid=%+v\n", string(k), uxArray[0].Body.SrcTransaction.Hex())

			return nil
		})
		return nil
	})

}

// Returns an address given an Address.Bytes()
func AddressFromBytes(b []byte) (addr cipher.Address, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	if len(b) != 20+1+4 {
		return cipher.Address{}, errors.New("Invalid address length")
	}
	a := cipher.Address{}
	copy(a.Key[0:20], b[0:20])
	a.Version = b[20]
	if a.Version != 0 {
		return cipher.Address{}, errors.New("Invalid version")
	}

	chksum := a.Checksum()
	var checksum [4]byte
	copy(checksum[0:4], b[21:25])

	if checksum != chksum {
		return cipher.Address{}, errors.New("Invalid checksum")
	}

	return a, nil
}
