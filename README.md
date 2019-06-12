# Runtime

## Development

### Test

```bash
make test
```

### Setup development environment (mac)

start a beefy minikube

```bash
minikube start \
  --memory=12288 \
  --cpus=4 \
  --kubernetes-version=v1.12.0 \
  --vm-driver=virtualbox \
  --disk-size=30g \
  --extra-config=apiserver.enable-admission-plugins="LimitRanger,NamespaceExists,NamespaceLifecycle,ResourceQuota,ServiceAccount,DefaultStorageClass,MutatingAdmissionWebhook"
```

install istio

```bash
kubectl apply \
  --filename https://raw.githubusercontent.com/knative/serving/v0.5.2/third_party/istio-1.0.7/istio-crds.yaml &&
curl -L https://raw.githubusercontent.com/knative/serving/v0.5.2/third_party/istio-1.0.7/istio.yaml \
  | sed 's/LoadBalancer/NodePort/' \
  | kubectl apply --filename -
```

install knative

```bash
kubectl apply --selector knative.dev/crd-install=true \
   --filename https://github.com/knative/serving/releases/download/v0.6.0/serving.yaml \
   --filename https://github.com/knative/build/releases/download/v0.6.0/build.yaml \
   --filename https://github.com/knative/eventing/releases/download/v0.6.0/release.yaml \
   --filename https://github.com/knative/eventing-sources/releases/download/v0.6.0/eventing-sources.yaml \
   --filename https://github.com/knative/serving/releases/download/v0.6.0/monitoring.yaml \
   --filename https://raw.githubusercontent.com/knative/serving/v0.6.0/third_party/config/build/clusterrole.yaml
```

install knative part2

```bash
kubectl apply --filename https://github.com/knative/serving/releases/download/v0.6.0/serving.yaml --selector networking.knative.dev/certificate-provider!=cert-manager \
   --filename https://github.com/knative/build/releases/download/v0.6.0/build.yaml \
   --filename https://github.com/knative/eventing/releases/download/v0.6.0/release.yaml \
   --filename https://github.com/knative/eventing-sources/releases/download/v0.6.0/eventing-sources.yaml \
   --filename https://github.com/knative/serving/releases/download/v0.6.0/monitoring.yaml \
   --filename https://raw.githubusercontent.com/knative/serving/v0.6.0/third_party/config/build/clusterrole.yaml
```

modify `config/config.yaml` to include your docker.io credentials (base64 encoded) and update the dockerregistry value to your docker.io username

### Local Deployment

#### Manager running locally

Install the CRD to a local Kubernetes cluster:

```bash
make install
```

Run the controller on your machine:

```bash
make run
```

#### Manager running inside k8s cluster

```bash
eval $(minikube docker-env)
make docker-build
make install
make deploy
```

### Prod Deployment

Uncomment `manager_image_patch_dev` in `kustomization.yaml`
Then run the following commands:

```bash
make install
make docker-build
make docker-push
make deploy
```

### Run the examples

```bash
kubectl apply -f config/samples/runtime_v1alpha1_function.yaml
```

access the function

```bash
curl -v -H "Host: $(kubectl get ksvc sample --no-headers | awk '{print $2}')" http://$(minikube ip):$(kubectl get svc istio-ingressgateway --namespace istio-system --output 'jsonpath={.spec.ports[?(@.port==80)].nodePort}')
```
