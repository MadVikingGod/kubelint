package testdata

var HasImageTagYaml = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: hasImageTag
  namespace: hasImageTag
spec:
  template:
    spec:
      containers:
        - name: hasImageTagContainer
          image: ThisHasATag:Thisisabogustag
`

var LatestImageTagYaml = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: latestImageTag
  namespace: latestImageTag
spec:
  template:
    spec:
      containers:
        - name: latestImageTagContainer
          image: ImageWithLatest:latest
`

var NoImageTagYaml = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: noImageTag
  namespace: noImageTag
spec:
  template:
    spec:
      containers:
        - name: noImageTagContainer
          image: thisDoesnthaveaTag
`

var NoImageTagInitContainer = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: noImageTagInit
  namespace: noImageTagInit
spec:
  template:
    spec:
      containers:
        - name: hasImageTagContainer
          image: ThisHasATag:Thisisabogustag
      initContainers:
        - name: noImageTagInitContainer
          image: ThisDoesNotHaveATag
`
