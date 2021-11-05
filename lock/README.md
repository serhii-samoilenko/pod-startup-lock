# Lock service

## How it works

##### 1. Starts and listens for HTTP requests
Responds with `423 Locked` until initialized.
Binding `host` and `port` are configurable.

##### 2. Checks TCP/HTTP endpoints if configured
Waits till all dependent endpoints are accessible. Then allows acquiring of the lock. 
Endpoints are checked constantly. Locking is allowed only if all are OK.

##### 3. First request(s) acquires the lock
Client gets `200 OK`. Lock is acquired for a specific time. Multiple locks can be configured for parallel acquiring.
Custom lock duration may be specified in the request itself.

##### 4. Subsequent requests are denied to acquire the lock
Client gets `423 Locked` until lock timeout exceeds.

##### That's it. Steps 2 - 4 are constantly repeated

## Request custom lock duration
You can configure default lock timeout. But each client may request custom duration with `GET` parameter, for example, 
`http://localhost:8888?duration=60`
To acquire a lock for 60 seconds.

## Dependent Endpoints check
This is useful when you need to wait for certain service(s) start before allowing starting of applications in the cluster.
* You may specify `http/https` endpoint, like `http://myelasticsearch.local:9200`
  Response with HTTP code `2XX` is considered as a success.
* You may specify `tcp` endpoint, like `tcp://mymongodb:27017`
  Established TCP connection is considered as a success.

## Configuration
You may specify additional command line options to override defaults:

| Option      | Default | Description |
| ----------- |---------| ----------- |
| `--host`    | 0.0.0.0 | Address to bind |
| `--port`    | 8888    | Port to bind    |
| `--locks`   | 1       | Number of locks allowed to be acquired at the same time |
| `--timeout` | 10      | Default time until the acquired lock is released, seconds |
| `--check`   | *none*  | List of endpoints to check before allow locking, see example below |
| `--failHc`  | 10      | Pause between endpoint checks if the previous check failed, seconds |
| `--passHc`  | 60      | Pause between endpoint checks if the previous check succeeded, seconds |

## How to run locally
Example with some command line options:
```bash
go run lock/main.go --port 9000 --locks 2 --check http://myelasticsearch:9200 --check tcp://mymongodb:27017
```

## How to deploy to Kubernetes
The preferable way is to deploy as a DaemonSet. Sample deployment YAML (notice checked endpoint):
```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: startup-lock
spec:
  selector:
    matchLabels:
      app: startup-lock
  template:
    metadata:
      labels:
        app: startup-lock
        version: '1.0'
    spec:
      hostNetwork: true
      nodeSelector:
        kubernetes.io/os: linux
      containers:
        - name: startup-lock-container
          image: lisenet/startup-lock
          imagePullPolicy: IfNotPresent
          args: ["--port", "8888", "--locks", "1", "--check", "http://$(HOST_IP):9999"]
          ports:
            - name: http
              containerPort: 8888
          env:
            - name: HOST_IP
              valueFrom:
                fieldRef:
                  fieldPath: status.hostIP
```
