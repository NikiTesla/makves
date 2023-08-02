package items

import "fmt"

type Storage interface {
	getItem(id int) (record string, err error)
}

type LocalStorage struct {
	storage map[int]string
}

func NewLocalStorage(filename string) (*LocalStorage, error) {
	localStorage := &LocalStorage{}

	return localStorage, nil
}

func (lS *LocalStorage) getItem(id int) (string, error) {
	record, ok := lS.storage[id]
	if !ok {
		return "", fmt.Errorf("no such record in a local storage")
	}

	return record, nil
}
