package validators

func IsInvalidRequest(request int) bool {
	return request < 1
}
