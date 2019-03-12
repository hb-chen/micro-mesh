RBAC
===

- `security-rbac.yaml`访问控制配置
- `security-policy.yaml`认证配置，`RBAC`依赖目标服务开启双向TLS认证
- `destination-rule-tls.yaml`服务请求策略开启TLS
- `service-deployment-account.yaml`添加ServiceAccount

## `destination-rule.yaml`服务请求策略开启TLS
```diff
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: mm-example-api
spec:
  host: mm-example-api
+  trafficPolicy:
+    tls:
+      mode: ISTIO_MUTUAL
  subsets:
  - name: v1
    labels:
      version: v1
---
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: mm-example-srv-1
spec:
  host: mm-example-srv-1
+  trafficPolicy:
+    tls:
+      mode: ISTIO_MUTUAL
  subsets:
  - name: v1
    labels:
      version: v1
---
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: mm-example-srv-2
spec:
  host: mm-example-srv-2
+  trafficPolicy:
+    tls:
+      mode: ISTIO_MUTUAL
  subsets:
  - name: v1
    labels:
      version: v1
---
```

## `service-deployment.yaml`添加ServiceAccount
```diff
##################################################################################################
# SRV_1 service
##################################################################################################
apiVersion: v1
kind: Service
metadata:
  name: mm-example-srv-1
  labels:
    app: mm-example-srv-1
spec:
  ports:
  - port: 9080
    name: grpc
  selector:
    app: mm-example-srv-1
+ ---
+ apiVersion: v1
+ kind: ServiceAccount
+ metadata:
+   name: mm-example-srv-1
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: mm-example-srv-1-v1
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: mm-example-srv-1
        version: v1
    spec:
+      serviceAccountName: mm-example-srv-1
      containers:
      - name: mm-example-srv-1
        command: [
          "/main",
          "-serve_addr=:9080",
          "-remote_addr=:9080"
        ]
        image: hbchen/micro-mesh-example-srv-1:v0.0.3
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 9080
---
##################################################################################################
# SRV_2 service
##################################################################################################
apiVersion: v1
kind: Service
metadata:
  name: mm-example-srv-2
  labels:
    app: mm-example-srv-2
spec:
  ports:
  - port: 9080
    name: grpc
  selector:
    app: mm-example-srv-2
+ ---
+ apiVersion: v1
+ kind: ServiceAccount
+ metadata:
+   name: mm-example-srv-2
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: mm-example-srv-2-v1
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: mm-example-srv-2
        version: v1
    spec:
+      serviceAccountName: mm-example-srv-2
      containers:
      - name: mm-example-srv-2
        command: [
          "/main",
          "-serve_addr=:9080",
          "-remote_addr=:9080"
        ]
        image: hbchen/micro-mesh-example-srv-2:v0.0.3
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 9080
---
##################################################################################################
# API service
##################################################################################################
apiVersion: v1
kind: Service
metadata:
  name: mm-example-api
  labels:
    app: mm-example-api
spec:
  ports:
  - port: 9080
    name: http
  selector:
    app: mm-example-api
+ ---
+ apiVersion: v1
+ kind: ServiceAccount
+ metadata:
+   name: mm-example-api
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: mm-example-api-v1
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: mm-example-api
        version: v1
    spec:
+      serviceAccountName: mm-example-api
      containers:
      - name: mm-example-api
        command: [
          "/main",
          "-serve_addr=:9080",
          "-remote_addr=:9080"
        ]
        image: hbchen/micro-mesh-example-api:v0.0.3
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 9080
---
```