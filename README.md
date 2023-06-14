# Kubernetes Cronjob Example with Multiple Binaries

This example shows a service handling some transactions.
It runs on 8080, and listens for create transaction requests.
It also has a cronjob to retry failed transactions. The cronjob configuration file is `cronjob.yaml`. It is scheduled to run every 30 mins.

To play with the project on your local:

1. Install [minikube](https://minikube.sigs.k8s.io/docs/start) and run.
```
minikube start
````

2. Apply kubernetes resources for service.
```
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

3. Forward go service.
```
minikube service kubernetes-cron-example-multiple-binaries-service
```

4. Send a request and see the logs.
```
curl --location --request POST 'http://{serviceAddress}/transaction'
kubectl get pod
kubectl logs {podName}
```

5. Apply kubernetes resources for cronjob.
```
kubectl apply -f cronjob.yaml
````

6. Check if a job is created.
```
kubectl get pod --watch
kubectl get jobs
kubectl logs {podName}
```