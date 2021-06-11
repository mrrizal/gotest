package main

import "github.com/pkg/errors"

type Dictionary map[string]string

func (d Dictionary) Search(keyword string) (string, error) {
	if val, ok := d[keyword]; ok {
		return val, nil
	}
	return "", errors.New("not found")
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	if err != nil {
		d[key] = value
		return nil
	}
	return errors.New("already exists")
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	if err != nil {
		return errors.New("not found")
	}
	d[key] = value
	return nil
}

func main() {

}
