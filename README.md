# Wrapper for K8s Jobs + Istio Sidecar

Sidecar containers do not works well with k8s jobs. The job will keep running so long as the sidecar proxy is running. This is a general problem not specific to Istio (see [kubernetes/kubernetes#25908](https://github.com/kubernetes/kubernetes/issues/25908)). The typical sollution includes explicit terminate signaling between app and sidecar container such that the sidecar can exit when the app does.

***K8s Job wrapper*** is very simple and light weight program, written in golang, that start a desired process and terminate istio-proxy sidecar after process is finish. The wrapper will send a signal to Envoy proxy to shutdown, the signal is a HTTP POST to http://localhost:15020/quitquitquit. This will gracefully shutdown both pilot-agent and the Envoy proxy.

## Kubernetes Jobs

A [Job](https://kubernetes.io/docs/concepts/workloads/controllers/job/) creates one or more Pods and will continue to retry execution of the Pods until a specified number of them successfully terminate. As pods successfully complete, the Job tracks the successful completions. When a specified number of successful completions is reached, the task (ie, Job) is complete. 

## Istio Proxy Sidecar

[Sidecar](https://istio.io/latest/docs/reference/config/networking/sidecar/) describes the configuration of the sidecar proxy that mediates inbound and outbound communication to the workload instance it is attached to.

## Job + Sidecar = Pod Status Pending

A Job is not considered complete until all containers have stopped running, and Istio Sidecars run indefinitely.

## Issues

* https://github.com/istio/istio/issues/6324
* https://github.com/istio/istio/issues/11045
* https://github.com/kubernetes/kubernetes/issues/25908

## Distroless Docker Images

In a distroless environment you don’t have access to ```sleep``` or ```curl``` as suggested in [istio/issues/6324](https://github.com/istio/istio/issues/6324).

* https://www.solo.io/blog/challenges-of-running-istio-distroless-images/
* https://stackoverflow.com/questions/54921054/terminate-istio-sidecar-istio-proxy-for-a-kubernetes-job-cronjob

## How to use

Wrapper build.
```
go build -a -o wrapper ./cmd/wrapper
```

[Dockerfile](./Dockerfile) example.

```dockerfile
FROM gcr.io/distroless/base:latest-amd64
...
ENTRYPOINT [ "/wrapper", "<your cmd>", "<arg 1>", "<arg 2>" ]
```

```shell
docker build . -t some-job-wrapper
```

Kubernetes Job.

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: pi
  labels:
    sidecar.istio.io/inject: "true"
spec:
  template:
    spec:
      containers:
      - name: pi
        image: some-job-wrapper
        command: [ "/wrapper", "<your cmd>", "<arg 1>", "<arg 2>" ]
      restartPolicy: Never
  backoffLimit: 4
```

## References

* https://www.envoyproxy.io/docs/envoy/latest/operations/admin#post--quitquitquit
* https://medium.com/redbox-techblog/handling-istio-sidecars-in-kubernetes-jobs-c392661c4af7
