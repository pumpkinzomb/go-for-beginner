package dictionary

import "errors"

type Dictionary map[string]string

var (
	errNotFound = errors.New("There is not exists.")
 	errAlreadyExist = errors.New("There is already exist.")
	errCantEdit = errors.New("Can't edit the word, It is not exist.")
)
func (d Dictionary) Get(word string) (string, error) {
	value, exist := d[word];
	if(exist){
		return value , nil
	}else {
		return "", errNotFound 
	}
}

func (a Dictionary) Add(word string, description string)(error){
	_, err := a.Get(word)
	if(err == errNotFound){
		a[word] = description
	}else if(err == nil){
		return errAlreadyExist
	}
	return nil
}

func (a Dictionary) Edit(word string, description string)(error){
	_, err := a.Get(word)
	if(err == errNotFound){
		return errCantEdit
	}else if(err == nil){
		a[word] = description
	}
	return nil
}