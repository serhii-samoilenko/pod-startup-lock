# Simple time-based lock service with HTTP interface.

[![Go Report Card](https://goreportcard.com/badge/github.com/lisenet/pod-startup-lock)](https://goreportcard.com/report/github.com/lisenet/pod-startup-lock)

Designed at [Oath](https://www.oath.com) to solve the [Thundering herd problem](https://en.wikipedia.org/wiki/Thundering_herd_problem) during multiple applications startup in the [Kubernetes](https://kubernetes.io) cluster. 

## The Problem

Starting multiple applications simultaneously on the same host may cause a performance bottleneck.
In Kubernetes this usually happens when applications are automatically deployed to a newly added Node.
In the worst-case scenario, application startup may be slowed down so dramatically that they fail to pass the healthcheck. 
They are then restarted by Kubernetes just to start fighting for shared resources again, in an endless loop.

## The Solution

Kubernetes allows a Pod to have additional, [Init container](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/#examples),
and postpone application startup until Init container finishes execution.
The solution is to deploy Lock service as a DaemonSet on a Pod, and each init container will sequentially acquire this lock.
So moments of application container starts will be distributed in time.

## Components

See Readmes in subfolders for details.

* #### [Lock](lock/README.md)
  HTTP service to be deployed one instance per Node (as a DaemonSet).
  Returns code `200 OK` as a response to the first request.
  Returns `423 Locked` to the subsequent requests until timeout exceeded.
  May depend on additional endpoint check.
  
* #### [Init](init/README.md)
  Lightweight client for the Lock service. To be deployed as Init Container alongside the main application container.
  Periodically tries to acquire the lock. Once succeeded, terminates, allowing the main container to start running.
  
* #### [K8s-health](k8s-health/README.md)
  Optional component. Performs healthcheck of Kubernetes DaemonSets.
  May be used by Lock service to postpone lock acquiring until all DaemonSets on the Node are up and running.

## How to build locally

The project is built using [Make](https://www.gnu.org/software/make/).

### 0. Install [Go](https://golang.org/dl/)

```bash
wget https://golang.org/dl/go1.18.10.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.18.10.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

### 1. Install [Dep](https://golang.github.io/dep) and update dependencies

The project uses Dep for dependency management. You can find installation instructions on the project page.

```bash
mkdir -p ~/go/{bin,src}
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```

```bash
cd ~/go/src/
git clone https://github.com/lisenet/pod-startup-lock.git
cd ./pod-startup-lock/
```

Then run Make: 

```bash
make dep
go mod init
go mod tidy
go mod vendor
```

### 2. Set target platform for Go binaries

Optional step, default is `linux`. You can use Make task as shown below.

Unix example:

```bash
export GOOS=openbsd
```

Windows:

```bash
set GOOS=windows
```

### 3. Build binaries

Run Make:

```bash
make
```

Or with target platform specified: 

```bash
make darwin build
```

Binaries will be located in sbfolder's `bin` folders:
* `init/bin/init`
* `k8s-health/bin/health`
* `lock/bin/lock`

### 4. Or Build Docker images

First, you need to specify Docker user name as a variable: `DOCKER_USER`.

```bash
export DOCKER_USER=${USERNAME}
```

Then run Make:

```bash
make docker-build
```

### 5. Or Build Docker images and push them to the repo

First, you need to specify Docker credentials as environment variables: `DOCKER_URL`, `DOCKER_USER`, and `DOCKER_PW`.

Then run Make:

```bash
make docker-push
```

## Release Notes

* `1.0.3`
    - Bumped Go version to 1.18.
* `1.0.2`
    - Bumped Go version to 1.16. Bumped Kubernetes version to 1.16.
* `1.0.1`
    - Added connection timeouts for http and tcp connections; Added keep-alive for http connections. 
* `1.0.0`
    - Initial version.
    
## Contributing

Please feel free to submit issues, fork the repository and send pull requests!
