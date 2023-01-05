TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=local
NAMESPACE=rprtr258
NAME=seventv
BINARY=terraform-provider-${NAME}
VERSION=0.0.1
OS_ARCH=linux_amd64

default: install

build:
	go build -o ${BINARY}

release:
	go install github.com/goreleaser/goreleaser@latest
	goreleaser release --rm-dist --snapshot --skip-publish  --skip-sign

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

clear:
	cd examples && rm -rf .terraform* terraform* && terraform init

example-apply: install
	cd examples && TF_LOG=WARN terraform apply

example-plan: install
	cd examples && terraform validate && TF_LOG=WARN terraform plan

example-reflex:
	reflex -r '\.(go|tf)$$' -- make example-plan

