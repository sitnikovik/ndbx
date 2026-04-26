consistency ?= QUORUM
n = 1000

.PHONY: bin
bin:
	@cd ./sensorgen && make build

.PHONY: gen
gen:
	./bin/sensorgen -n=${n} -hosts=localhost -port=9042 -keyspace=iot -consistency=${consistency}
