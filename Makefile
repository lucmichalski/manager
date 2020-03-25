
# Image URL to use all building/pushing image targets
CONTROLLER_TAG ?= controller:latest
SERVER_TAG ?= server:latest
# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:trivialVersions=true"

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif


all: manager

setup: ; $(info $(M) setting up env variables for test…) @ ## Setup env variables
export LOCAL=true

# Run tests
test: setup generate fmt vet manifests
	go test ./... -coverprofile cover.out

# Build manager binary
manager: generate fmt vet
	go build -o bin/manager main.go

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet manifests
	go run ./main.go

# Install CRDs into a cluster
install: manifests
	kustomize build config/crd | kubectl apply -f -

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy: manifests
	cd config/server && kustomize edit set image server=${SERVER_TAG}
	cd config/manager && kustomize edit set image controller=${CONTROLLER_TAG}
	kustomize build config/default | kubectl apply -f -

undeploy:
	kustomize build config/default | kubectl delete -f -

# Generate manifests e.g. CRD, RBAC etc.
manifests: controller-gen
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Generate code
generate: controller-gen
	$(CONTROLLER_GEN) object:headerFile=./hack/boilerplate.go.txt paths="./..."


# Build the docker image
docker-build: test
	docker build . -t ${CONTROLLER_TAG}
	docker build . -f Dockerfile-server -t ${SERVER_TAG}

# Push the docker image
docker-push:
	docker push ${CONTROLLER_TAG}
	docker push ${SERVER_TAG}

# find or download controller-gen
# download controller-gen if necessary
controller-gen:
ifeq (, $(shell which controller-gen))
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.5
CONTROLLER_GEN=$(GOBIN)/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif

.PHONY: client-gen-tools
# find or download client-gen
# download client-gen if necessary
client-gen-tools:
ifeq (, $(shell which client-gen))
	go get k8s.io/code-generator/cmd/client-gen
CLIENT_GEN=$(GOBIN)/client-gen
else
CLIENT_GEN=$(shell which client-gen)
endif

proto-gen-tools:
ifeq (, $(shell which protoc-gen-go))
	brew install protobuf
	go get -u github.com/golang/protobuf/protoc-gen-go
PROTO_GEN=$(GOBIN)/protoc-gen-go
PROTO_COMPILER=$(shell which protoc)
else
PROTO_GEN=$(shell which protoc-gen-go)
endif

proto: proto-gen-tools
	@echo "Generating protogen files:"
 	$(shell protoc --go_out=paths=source_relative,plugins=grpc:. pkg/grpc/proto/cluster/cluster.proto)

.PHONY: clientgen
clientgen: client-gen-tools
	$(shell rm -rf pkg/k8s/customclient) \
    $(CLIENT_GEN) -h ./hack/boilerplate.go.txt -n customclient --input-base github.com/keikoproj/manager/api --input custom/v1alpha1 -p github.com/keikoproj/manager/pkg/k8s --fake-clientset=false

