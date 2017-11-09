package main

import (
	//"bytes"
	"encoding/hex"
	//"errors"
	"fmt"
	"log"
	"sort"

	"github.com/boltdb/bolt"
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

func main() {
	get_blockchain()
	//get_blocktree()
	//get_transction()
	//get_blogsig()
	//get_uxouts()
	//get_unspent_pool()
	//get_unconfirmed_txns()
}

func get_blockchain() {
	opt := &bolt.Options{ReadOnly: true}
	db, err := bolt.Open("/Users/liuguirong/.skycoin/data.db", 0400, opt)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	bi := []BlockInfo{}

	//block_tree := "block_tree"
	//blocks := "blocks"
	//blockchain_meta := "blockchain_meta"
	//block_sigs := "block_sigs"

	//transactions := "transactions"
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

func get_blocktree() {
	db, err := bolt.Open("/Users/liuguirong/.skycoin/data.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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

func get_transction() {
	db, err := bolt.Open("/Users/liuguirong/.skycoin/data.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bucketname := "transactions"

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketname))
		fmt.Printf("bucket-name: %s\n", bucketname)

		b.ForEach(func(k, v []byte) error {
			tx := coin.TransactionDeserialize(v)
			fmt.Printf("k=%+v\n", hex.EncodeToString(k))
			format_transaction(tx)

			return nil
		})
		return nil
	})

}

func get_blogsig() {
	db, err := bolt.Open("/Users/liuguirong/.skycoin/data.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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
func get_uxouts() {
	db, err := bolt.Open("/Users/liuguirong/.skycoin/data.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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

			return nil
		})
		return nil
	})

}
func get_unspent_pool() {
	db, err := bolt.Open("/Users/liuguirong/.skycoin/data.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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

			return nil
		})
		return nil
	})

}
func get_unconfirmed_txns() {
	db, err := bolt.Open("/Users/liuguirong/.skycoin/data.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bucketname := "unconfirmed_txns"

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketname))
		fmt.Printf("bucket-name: %s\n", bucketname)

		b.ForEach(func(k, v []byte) error {
			fmt.Printf("k=%+v, v=%+v\n", k, v)

			return nil
		})
		return nil
	})

}
