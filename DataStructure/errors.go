package DataStructure

type GenericError struct {
	typeOfError, category, description string
}

func (e GenericError) Error() string {
	return e.category + ": " + e.typeOfError + ": " + e.description
}

type ErrorList []GenericError

func (eList ErrorList) New(typeOfError string,
	category string,
	description string) ErrorList {
	return append(eList, GenericError{typeOfError, category, description})
}

func (eList ErrorList) GetErrors() []error {
	k := make([]error, len(eList))
	for i, v := range eList {
		k[i] = v
	}
	return k
}
