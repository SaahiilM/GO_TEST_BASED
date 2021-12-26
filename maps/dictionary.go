package maps

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrWordNotFound = DictionaryErr("could not find the searched word")

	ErrWordExists = DictionaryErr("cannot add word, it already exists")

	ErrWordDoesNotExist = DictionaryErr("cannot update, word does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrWordNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word string, newDefinition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrWordNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}
