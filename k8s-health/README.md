# Kubernetes DaemonSet health check service
Typically you would like to postpone application startup until all DaemonSets on Node (or at least mapped to the host network) are ready.
This util constantly performs DaemonSet heallthcheck and respond with `200 OK` if passed.
 
Just add this service as a dependent endpoint to the Lock service, and lock won't be acquired until DaemonSets are ready.  

## How it works

##### 1. Starts and listens for HTTP requests
Responds with `412 Precondition Failed` until healthckeck succeeds.
Binding `host` and `port` are configurable.

##### 2. Uses Kubernetes API to get a list of DaemonSets required on the same node
`NODE_NAME` environment variable must be set.
No additional authentication required to access Kubernetes API if running as a cluster resource.

##### 3. When all DaemonSet Pods are up and running, return success
Responds with `200 OK`. Constantly repeats DaemonSet healthcheck.

## Which DaemonSets to check
You may need to check only certain DaemonSets and ignore the others:
* Add `--hostNet` flag to check only ones with binding to the host network, i.e. with `spec.template.spec.hostNetwork: true`
* Use `--namespace XXX` to check only ones in a specific namespace. All namespaces by default.
* List excluded labels with `--ex` flag. DaemonSets having **at least one** matching label will be excluded.  
* **OR**, list included labels with `--in` flag. DaemonSets having **all** matching labels will be included, rest excluded.
  You can't specify both `--ex` and `--in` flags, choose one.  

## In Cluster / Out Of Cluster configuration
[`kubernetes-go-client`](https://github.com/kubernetes/client-go) is used under the hood.

By default, `kubernetes-go-client` uses *in-cluster* configuration.
It will try to create configuration basing on cluster info from the running pod.
If you do override `baseUrl` parameter, the *out-of-cluster* configuration will be used,
assuming you are running k8s proxy with `kubectl`

## Required Configuration
Environment variable `NODE_NAME` must be exposed to indicate which node health should be checked.

## Additional Configuration
You may specify additional command line options to override defaults:

| Option        | Default | Description |
| ------------- |---------| ----------- |
| `--host`      | 0.0.0.0 | Address to bind |
| `--port`      | 9999    | Port to bind    |
| `--baseUrl`   | *none*  | K8s API Base Url. **Only to run in the out-of-cluster mode** |
| `--namespace` | *none*  | Target K8s namespace where to perform DaemonSets healthcheck. Leave blank for all namespaces |
| `--hostNet`   |         | Check only DaemonSets bind to the `host network` |
| `--failHc`    | 10      | Pause between healthchecks if the previous check failed, seconds |
| `--passHc`    | 60      | Pause between healthchecks if the previous check succeeded, seconds |
| `--in`        | *none*  | DaemonSet labels to include in healthcheck, Format: `label:value` |
| `--ex`        | *none*  | DaemonSet labels to exclude from healthcheck, Format: `label:value` |

## How to run locally
Example with some command line options:
```bash
kubectl proxy -p 57585
export NODE_NAME=10.11.10.11
go run k8s-health/main.go --baseUrl http://127.0.0.1:57585 --in app:test --in version:1.1 --hostNet
```

## How to deploy to Kubernetes
The preferable way is to deploy as a DaemonSet. Sample deployment YAML:
```yaml
apiVersion: extensions/v1beta1
kind: DaemonSet
metadata:
  name: startup-lock-k8s-health
spec:
  template:
    metadata:
      labels:
        app: startup-lock-k8s-health
        version: '1.0'
    spec:
      hostNetwork: true
      nodeSelector:
        kubernetes.io/role: node
      containers:
        - name: startup-lock-k8s-health-container
          image: ssamoilenko/startup-lock-k8s-health
          args: ["--port", "9999", "--hostNet"]
          ports:
            - name: http
              containerPort: 9999
          env:
            - name: NODE_NAME
              valueFrom:
                fieldRef:
                  fieldPath: spec.nodeName
```
