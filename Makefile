
.PHONY: clean build ti tp ta

DIST_DIR=./dist
BIN_NAME=terraform-provider-onelogin
BIN_PATH=${DIST_DIR}/${BIN_NAME}

PLUGINS_DIR=~/.terraform.d/plugins

clean:
	rm -r ${DIST_DIR}

clean-terraform:
	rm terraform.*

build:
	mkdir -p ${DIST_DIR}
	go build -o ${DIST_DIR} ./...
	mv ${DIST_DIR}/cmd ${BIN_PATH}

sideload: clean build
	mkdir -p ${PLUGINS_DIR}
	cp ${BIN_PATH} ${PLUGINS_DIR}/${BIN_NAME}

ti:
	terraform init

tp:
	terraform plan

ta:
	terraform apply
