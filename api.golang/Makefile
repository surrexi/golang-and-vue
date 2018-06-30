GCC=go
GCMD=run
GPATH=main.go

run:
#	make build
	$(GCC) $(GCMD) $(GPATH)

build:
	make build_db

build_db:
	rm pkg/db/db_structs.go
	go run pkg/db/main.go -json=./pkg/db/config.json
	mv db_structs.go pkg/db