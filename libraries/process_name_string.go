package libraries

import "strings"

func ProcessNameString(stringValue string) string {
	signature := strings.ToLower(stringValue)
	signature = strings.NewReplacer(
		"%", "",
		"/", "",
		"~", "",
		"?", "",
		"&", "",
		"'", "",
		".", "",
		"(", "",
		")", "",
		"[", "",
		"]", "",
		" etc ", "",
		" etc. ", "",
		" from ", " ",
		",", " ",
		"-", " ",
		" and ", " ",
		" the ", " ",
		" of ", " ",
		" an ", " ",
		" a ", " ",
		" for ", " ",
		" in ", " ",
		" at ", " ",
		" on ", " ",
	).Replace(signature)
	return signature
}
