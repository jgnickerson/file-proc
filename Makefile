MASTER_IMAGE_BASE ?= local-registry/file-proc-master
MASTER_IMAGE_TAG ?= latest
MASTER_IMAGE ?= $(MASTER_IMAGE_BASE):$(MASTER_IMAGE_TAG)

SLAVE_IMAGE_BASE ?= local-registry/file-proc-slave
SLAVE_IMAGE_TAG ?= latest
SLAVE_IMAGE ?= $(SLAVE_IMAGE_BASE):$(SLAVE_IMAGE_TAG)

.PHONY: phony

build-master: phony
	docker build -t $(MASTER_IMAGE) ./master/

build-slave: phony
	docker build -t $(SLAVE_IMAGE) ./slave/

build: phony build-master build-slave

up: phony build
	docker-compose up --remove-orphans
