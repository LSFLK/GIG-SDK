package libraries

import "log"

func ReportErrorWithoutCrash(err error) {
	if err != nil {
		log.Println(err)
	}
}
