# Wrapper for K8s Jobs + Sidecars

Sidecar containers do not works well with k8s jobs. The job will keep running so long as the sidecar proxy is running. This is a general problem not specific to Istio (see [kubernetes/kubernetes#25908](https://github.com/kubernetes/kubernetes/issues/25908)). The typical sollution includes explicit terminate signaling between app and sidecar container such that the sidecar can exit when the app does.

***K8s Job wrapper*** is very simple and light weight program, written in golang, that start a desired process and execute a hook function to terminate the sidecar,  after process is finish. 

There is a hook implementation for Envoy Proxy sidecars,where a HTTP POST request is sent to http://localhost:15020/quitquitquit. This will gracefully shutdown both pilot-agent and the Envoy Proxy.

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

## Wrapper plugin implementation

Echo plugin example.

```golang
package main

import "log"

// GetPlugin returns a plugin instance.
// Called by plugin loader.
func GetPlugin() interface{} {
	return &echoPlugin{}
}

type echoPlugin struct{}

// Run is an implementation of PluginRunner interface.
// Your business logic should be added here.
func (p echoPlugin) Run() {
	log.Println("Echo demo plugin")
}
```

## How to use

Wrapper build.
```shell
go build -a -o wrapper ./cmd/wrapper
```

Plugin build.

```shell
go build -buildmode=plugin -a -o istio-proxy-plugin.so ./cmd/plugin/istio-proxy
```

Plugin build, in debug mode withou optimizations and inline.

```shell
go build -buildmode=plugin -gcflags="all=-N -l" -a -o istio-proxy-plugin.so ./cmd/plugin/istio-proxy
```

### Option 1 - Dockerfile ENTRYPOINT

[Dockerfile](./Dockerfile) example.

```dockerfile
FROM gcr.io/distroless/base:latest-amd64
...
COPY wrapper /
COPY istio-proxy-plugin.so /
...
ENV WRAPPER_PLUGIN_PATH=/istio-proxy-plugin.so
...
ENTRYPOINT [ "/wrapper", "ls", "-l" ]
```

```shell
docker build . -t some-job-wrapper
```

### Option 2 - Kubernetes manifest

The docker image, used bellow, must have the wrapper and wrapper plugin inside.

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
        command: [ "/wrapper", "ls", "-l" ]
        env:
        - name: WRAPPER_PLUGIN_PATH
          value: "/istio-proxy-plugin.so"
      restartPolicy: Never
  backoffLimit: 4
```

The [plugin](./cmd/plugin) for Istio Proxy sidecar will perform a request to http://localhost:15020/quitquitquit, a local administration interface that can be used to cleanly exit the server.
