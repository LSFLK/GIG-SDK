#!/bin/bash
cd tests/libraries
go test
cd ../request_handlers
go test
cd ../models/
go test