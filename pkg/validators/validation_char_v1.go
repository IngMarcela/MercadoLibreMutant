package validators

type RequestValidationDna struct{}

func (rv RequestValidationDna) IsValidateChar(char string) bool {
	return char == "A" || char == "C" || char == "G" || char == "T"
}

func NewValidateChar() RequestValidationDna {
	return RequestValidationDna{}
}
