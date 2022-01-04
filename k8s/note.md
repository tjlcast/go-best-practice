## 环境的准备

docker daemon
kubectl
kind # brew install kind


## kind

kind version

kind delete cluster --name my-cluster

kind create cluster --config=kind-cluster.yaml

kind load docker-image jialtang/greeter:v1.0.0 --name my-cluster


## kubectl

kubectl version --client

kubectl config get-contexts

kubectl config use-contexts minikube

kubectl get pods [-A] [-Oyaml]

kubectl apply -f greeter.yaml

kubectl delete -f greeter.yaml

kubectl describe po greeter

kubectl logs greeter [-f]

## 创建容器化应用 greeter

cd app && docker build . -t jialtang/greeter:v1.0.0 && docker push jialtang/greeter:1.0.0

## 新增加一个pod来部署greeter container
