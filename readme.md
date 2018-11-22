# Pingaling operator

## Build and run the operator

Register the CRD with the kubernetes apiserver

```s
$ kubectl create -f deploy/crds/pingaling_v1alpha1_pingaling_crd.yaml
customresourcedefinition.apiextensions.k8s.io/pingalings.pingaling.io created
```

Build and push the operator image

```s
$ operator-sdk build duyuyang/pingaling-operator:v0.0.1

$ docker push duyuyang/pingaling-operator:v0.0.1
The push refers to repository [docker.io/duyuyang/pingaling-operator]
...
```

Modify the deploy/operator.yaml to update the pingaling-operator image name

Setup RBAC and deploy the pingaling-operator:

```s
$ kubectl create -f deploy/service_account.yaml
$ kubectl create -f deploy/role.yaml
$ kubectl create -f deploy/role_binding.yaml
$ kubectl create -f deploy/operator.yaml
deployment.apps/pingaling-operator created
```

Verify the pingaling-operator is up and running:

```s
$ kubectl get deployment
NAME                 DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
pingaling-operator   1         1         1            1           9m
```

## Deploy pingaling service

```s
$ kubectl create -f deploy/crds/pingaling_v1alpha1_pingaling_cr.yaml
pingaling.pingaling.io/pingaling-dev created
```