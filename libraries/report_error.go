package libraries

import "log"

func ReportError(err error) {
	if err != nil {
		log.Println(err)
	}
}
