package items

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Storage interface {
	getItem(id int) (record string, err error)
}

// LocalStorage is a structure for interactions with storage
type LocalStorage struct {
	storage map[int]map[string]string
}

// NewLocalStorage creates local storage from csv file in map[string]string format
func NewLocalStorage(filename string) (*LocalStorage, error) {
	storage := make(map[int]map[string]string)

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot open file %s, err: %s", filename, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	var header []string
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("cannot read next record, err: %s\n", err)
			continue
		}

		if len(header) == 0 {
			header = record
			continue
		}

		data := make(map[string]string)
		for i, v := range record {
			data[header[i]] = v
		}

		id, err := strconv.Atoi(data["id"])
		if err != nil {
			log.Printf("cannot parse id, err: %s\n", err)
		}

		storage[id] = data

	}
	localStorage := &LocalStorage{storage: storage}

	return localStorage, nil
}

// getItem searches item in a storage
func (lS *LocalStorage) getItem(id int) (map[string]string, error) {
	record, ok := lS.storage[id]
	if !ok {
		return nil, fmt.Errorf("no such record in a local storage")
	}

	return record, nil
}
