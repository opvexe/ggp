# ggp

Introduction: Encapsulates the client-go resource invocation method of k8s.

## Kind

```
vi kind-config.yaml
-------------------
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
- role: worker
```

## Install (Macpro M1 Support version)
```shell
kind create cluster --config cluster.yaml --image kindest/node:v1.19.11
```
