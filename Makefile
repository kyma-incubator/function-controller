
# Image URL to use all building/pushing image targets
IMG ?= runtime-controller:latest

all: test manager

# Run tests
test: generate fmt vet
	go test -v ./pkg/... ./cmd/... -coverprofile cover.out

qt:
	go test -v ./pkg/... ./cmd/... -coverprofile cover.out

# Build manager binary
manager: generate fmt vet
	go build -o bin/manager github.com/kyma-incubator/function-controller/cmd/manager

# NOTE: does not work anymore until https://github.com/kubernetes-sigs/kubebuilder/issues/400 is implemented
# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet
	go run ./cmd/manager/main.go

# Install CRDs into a cluster
install: 
	kubectl apply -f config/crds/runtime_v1alpha1_function.yaml

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy: manifests
	kubectl apply -f config/crds
	kustomize build config/default | kubectl apply -f -
	kubectl apply -f config/config.yaml

# Generate manifests e.g. CRD, RBAC etc.
manifests:
	go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go all

# Run go fmt against code
fmt:
	go fmt ./pkg/... ./cmd/...

# Run go vet against code
vet:
	go vet ./pkg/... ./cmd/...

# Generate code
generate: dep
	go generate ./pkg/... ./cmd/...

dep:
	dep ensure --vendor-only

# Build the docker image
# docker-build: test
.PHONY: docker-build
docker-build: dep fmt
	docker build . -t ${IMG}
	@echo "updating kustomize image patch file for manager resource"
	sed -i'' -e 's@image: .*@image: '"${IMG}"'@' ./config/default/manager_image_patch_remote_dev.yaml ./config/default/manager_image_patch_local_dev.yaml

# Push the docker image
.PHONY: docker-push
docker-push: docker-build
	docker push ${IMG}
