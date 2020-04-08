package testdata

var HasResources = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: hasResource
  namespace: hasResource
spec:
  template:
    spec:
      containers:
        - name: hasResourceContainer
          resources: {}
`
var NoResources = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: noResource
  namespace: noResource
spec:
  template:
    spec:
      containers:
        - name: noResourceContainer
`
