#### kind create cluster --name=goexpert

#### kubectl cluster-info --context kind-goexpert

#### kubectl get nodes

#### kubectl apply -f k8s/deployment.yaml

#### kubectl apply -f k8s/service.yaml

#### kubectl port-forward svc/servicesvc 8080:8080
