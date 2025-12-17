package parser

//Retrieves the value of "category" if it exists; else, it returns an empty string.
func Category(meta []Meta) string {
	for _, m := range meta {
		if m.Type == "category" {
			return m.Value
		}
	}
	return ""
}

//Retrieves the value of "password" attributes if they exist.
func Passwords(meta []Meta) []string {
	var passwordSplice []string 
	for _, m := range meta {
		if m.Type == "password" {
			passwordSplice = append(passwordSplice, m.Value)
		}
	}
	return passwordSplice
}

//Retrieves the value of "tag" attributes if they exist.
func Tags(meta []Meta) []string {
	var tagSplice []string 
	for _, m := range meta {
		if m.Type == "tag" {
			tagSplice = append(tagSplice, m.Value)
		}
	}
	return tagSplice
}

//Retrieves the value of a "title" attribute if it exists; else, it returns an empty string.
func Title(meta []Meta) string {
	for _, m := range meta {
		if m.Type == "title" {
			return m.Value
		}
	}
	return ""
}
