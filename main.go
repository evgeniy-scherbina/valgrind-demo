package main

import "C"
import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/linxGnu/grocksdb"
)

const pathToDB = "/tmp/rocksdb_data"

func memoryLeak(n int) {
	for i := 0; i < n; i++ {
		var uuid *C.uchar
		uuid = (*C.uchar)(C.malloc(16))
		_ = uuid
	}
}

func rocksdb(n int, infinite bool, valueSize int) {
	bbto := grocksdb.NewDefaultBlockBasedTableOptions()
	cache := grocksdb.NewLRUCache(1 << 30)
	bbto.SetBlockCache(cache)
	bbto.SetFilterPolicy(grocksdb.NewBloomFilter(10))

	opts := grocksdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbto)
	opts.SetCreateIfMissing(true)
	opts.EnableStatistics()
	opts.SetStatsDumpPeriodSec(1)
	db, err := grocksdb.OpenDb(opts, pathToDB)
	if err != nil {
		log.Fatal(err)
	}
	_ = db

	simulateLoad(db, n, infinite, valueSize)
}

func simulateLoad(db *grocksdb.DB, n int, infinite bool, valueSize int) {
	ro := grocksdb.NewDefaultReadOptions()
	wo := grocksdb.NewDefaultWriteOptions()

	defaultKey := []byte("foo")
	defaultValue := make([]byte, valueSize)
	for i := 0; i < valueSize; i++ {
		defaultValue[i] = byte(i)
	}

	// if ro and wo are not used again, be sure to Close them.
	err := db.Put(wo, defaultKey, defaultValue)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < n || infinite; i++ {
		value, err := db.Get(ro, defaultKey)
		if err != nil {
			log.Fatal(err)
		}
		_ = value
		//defer value.Free()
	}
}

func main() {
	n := flag.Int("n", 0, "")
	delay := flag.Int("delay", 0, "")
	infinite := flag.Bool("infinite", false, "")
	valueSize := flag.Int("value_size", 0, "")
	flag.Parse()

	fmt.Printf("delay start\n")
	time.Sleep(time.Duration(*delay) * time.Second)
	fmt.Printf("delay end\n")

	fmt.Printf("start\n")
	fmt.Printf("n: %v\n", *n)
	//memoryLeak(*n)
	rocksdb(*n, *infinite, *valueSize)
	fmt.Printf("end\n")
}
