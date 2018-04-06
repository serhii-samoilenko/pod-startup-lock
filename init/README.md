# Init container for the Lock service

Application repeatedly queries Lock service endpoint until it gets `200 OK` response.
Then exits letting main application to start.

**Designed to be deployed as an Init Container**

## Additional Configuration
You may specify additional command line options to override defaults:

| Option      | Default   | Description |
| ----------- |-----------| ----------- |
| `--port`    | 8888      | Lock Service's HTTP port |
| `--host`    | localhost | Lock Service's hostname |
| `--pause`   | 10        | Pause between Lock acquiring attempts, seconds |
| `--timeout` | *none*    | Custom lock timeout to request, seconds |

## How to run locally
Example with some command line options:
```bash
go run init/main.go --port 9000 --timeout 15
```

## How to deploy to Kubernetes
Should be deployed as an Init Container. Sample deployment YAML:
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myapp-pod
  labels:
    app: myapp
spec:
  containers:
    - name: myapp-container
      image: busybox
      command: ['sh', '-c', 'echo The app is running! && sleep 3600']
  initContainers:
    - name: startup-lock-init-container
      image: ssamoilenko/startup-lock-init
      args: ["--host", "$(NODE_NAME)", "--port", "8888", "--timeout", "15"]
      env:
        - name: NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
```
