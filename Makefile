test-dbutils:
	docker-compose -f ./test/docker-compose.yml up -V --build --abort-on-container-exit --remove-orphans
	docker-compose -f ./test/docker-compose.yml down -v --remove-orphans --rmi all