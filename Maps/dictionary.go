package dictionary

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string //Wrapper for String

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word] //dictionary look ups return 2 values, the value itself, and a boolean that returns true or false depending on whether the search was successful or not
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
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

//a map value is a pointer to a hmap structure meaning that when you pass in a map to a function/method you are only copying the pointer address part, which means that you can modify its underlying data!

//you can't create a nil map and then write into it without causing a run time error: looks like this -> var m map[string]string

//instead do this:

// var dictionary = map[string]string{}

// OR

// var dictionary = make(map[string]string)
