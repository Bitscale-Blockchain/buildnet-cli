package utils

func ContainsOnlyAllowedValues(arr []string, allowedValues map[string]bool) bool {
	for _, s := range arr {
		if !allowedValues[s] {
			// If the string is not in the allowed values, it's invalid
			return false
		}
	}
	return true
}
