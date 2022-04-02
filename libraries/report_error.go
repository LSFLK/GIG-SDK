package libraries

import "log"

func ReportError(err error, v ...interface{}) {
	if err != nil {
		log.Println(err, v)
	}
}
