package main

import "github.com/pkg/errors"

type Dictionary map[string]string

func (d Dictionary) Search(keyword string) (string, error) {
	if val, ok := d[keyword]; ok {
		return val, nil
	}
	return "", errors.New("not found")
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}

func main() {

}
