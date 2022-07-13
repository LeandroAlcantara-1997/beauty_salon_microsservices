.PHONY: run
run:
	docker-compose -f ./appointments/build/docker-compose.yaml up -d
	docker-compose -f build/docker-compose.yml up -d