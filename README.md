## kubectl-connect

```
 _        _            _   _                           _
| |___  _| |__  ___ __| |_| |  __ ___ _ _  _ _  ___ __| |_
| / / || | '_ \/ -_) _|  _| | / _/ _ \ ' \| ' \/ -_) _|  _|
|_\_\\_,_|_.__/\___\__|\__|_| \__\___/_||_|_||_\___\__|\__|
```

## Installation
You can install `kubectl connect` either via source, or by using the [Krew](https://github.com/kubernetes-sigs/krew) plugin manager.

#### Installation from Source:
```
## Clone This Repo
% git clone git@github.com:awslabs/kubectl-connect && cd kubectl-connect

## Pull Necessary Dependencies
% go get kubectl-connect

## Build Package
% go build .

## Move Plug-In Binary Into System Path
% sudo mv kubectl-connect /usr/local/bin
```
> Note that if you need to compile this package for a non-native system architecture, such as building for x86_64/Linux from aarch64/macOS for example; modify the above `go build .` command as follows: `GOOS=linux GOARCH=x86_64 go build .`

#### Installation using Krew Plug-In Manager:
```
## Install Package
% kubectl krew install kubectl-connect
```

## Dependencies
- [kubectl](https://kubernetes.io/docs/tasks/tools/) - Kubernetes CLI Tool
- [krew](https://github.com/kubernetes-sigs/krew) - Plug-In Package Manager for `kubectl`
- [Go](https://www.golang.org/) - Version 1.21

## Bug reports

* If you have a problem with `kubectl-connect` itself, please file an
  issue in this repository.

## Security

See [CONTRIBUTING](CONTRIBUTING.md#security-issue-notifications) for more information.

## License

This library is licensed under the MIT-0 License. See the LICENSE file.

