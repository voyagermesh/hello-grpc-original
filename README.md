# hello-grpc
Hello gRPC !

## Run on Host

Build binary:

```console
$ ./hack/make.py
```

Run server:

```console
$ hello-grpc run
```

Run client (Intro API):

```
$ hello-grpc client --address=127.0.0.1:8080 --name=appscode
I0206 17:40:38.495948    4419 logs.go:19] intro:"hello, appscode!" 
```

Run client (Stream API):

```
$ hello-grpc client --address=127.0.0.1:8080 --name=appscode --stream
I0206 17:41:11.933033    4465 logs.go:19] intro:"0: hello, appscode!" 
I0206 17:41:12.933156    4465 logs.go:19] intro:"1: hello, appscode!" 
I0206 17:41:13.933201    4465 logs.go:19] intro:"2: hello, appscode!" 
I0206 17:41:14.933331    4465 logs.go:19] intro:"3: hello, appscode!" 
...
```

JSON response: visit http://127.0.0.1:8080/apis/hello/v1alpha1/intro/json?name=tamal

## Run in a Kubernetes Cluster
Tested against Minikube v0.25.0 (Kubernetes 1.9.0)

```console
$ kubectl apply -f ./hack/deploy/deploy.yaml

$ kubectl get pods,svc
NAME                             READY     STATUS    RESTARTS   AGE
po/hello-grpc-66b9f67c46-bnd2l   1/1       Running   0          10s
po/hello-grpc-66b9f67c46-wpw4b   1/1       Running   0          10s

NAME             TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)                                      AGE
svc/hello-grpc   LoadBalancer   10.104.191.103   <pending>     80:30596/TCP,443:31816/TCP,56790:30206/TCP   10s
svc/kubernetes   ClusterIP      10.96.0.1        <none>        443/TCP                                      34m

$ minikube service list
|-------------|----------------------|--------------------------------|
|  NAMESPACE  |         NAME         |              URL               |
|-------------|----------------------|--------------------------------|
| default     | hello-grpc           | http://192.168.99.100:30596    |
|             |                      | http://192.168.99.100:31816    |
|             |                      | http://192.168.99.100:30206    |
| default     | kubernetes           | No node port                   |
| kube-system | kube-dns             | No node port                   |
| kube-system | kubernetes-dashboard | http://192.168.99.100:30000    |
|-------------|----------------------|--------------------------------|

```

Now visit: http://192.168.99.100:30596/apis/hello/v1alpha1/intro/json?name=tamal

![hello-grpc](/docs/images/hello-grpc.png)


## Status Endpoint

Hello GRPC server has a `/apis/status/json` endpoint which can be used to probe heatlh of the service.

![hello-grpc-status](/docs/images/hello-grpc-status.png)
