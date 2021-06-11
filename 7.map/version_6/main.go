package main

const (
	ErrNotFound = DictionaryErr("not found")
	ErrWordExists = DictionaryErr("word already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(keyword string) (string, error) {
	if val, ok := d[keyword]; ok {
		return val, nil
	}
	return "", ErrNotFound
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	if err != nil {
		d[key] = value
		return nil
	}
	return ErrWordExists
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	if err != nil {
		return ErrNotFound
	}
	d[key] = value
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	if err != nil {
		return err
	}
	delete(d, key)
	return nil
}

func main() {

}
