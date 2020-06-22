all:
	make build && make run

build:
	docker build . -t python-ao

run:
	docker run --rm -p 5000:5000  python-ao


