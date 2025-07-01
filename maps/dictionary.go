package main

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("word does not exist")
	ErrWordExists       = DictionaryErr("word already exists")
	ErrWordDoesNotExist = DictionaryErr("word does not exist")
)

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, newDef string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = newDef
	case ErrWordDoesNotExist:
		return ErrWordDoesNotExist
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordDoesNotExist:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return def, nil
}

func (e DictionaryErr) Error() string {
	return string(e)
}
