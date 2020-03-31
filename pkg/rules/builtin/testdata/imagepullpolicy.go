package testdata

var ImagePullPolicyAlwaysYaml = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: alwaysImagePullPolicy
  namespace: alwaysImagePullPolicy
spec:
  template:
    spec:
      containers:
        - name: alwaysImagePullPolicyContainer
          imagePullPolicy: Always
`

var ImagePullPolicyMultiYaml = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: multiImagePullPolicy
  namespace: multiImagePullPolicy
spec:
  template:
    spec:
      containers:
        - name: multiImagePullPolicyContainer
          imagePullPolicy: Always
        - name: multiImagePullPolicyContainer-fail
          imagePullPolicy: Never

`

var NeverImagePullPolicyYaml = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: neverImagePullPolicy
  namespace: neverImagePullPolicy
spec:
  template:
    spec:
      containers:
        - name: neverImagePullPolicyContainer
          imagePullPolicy: Never
`

var NoImagePullPolicyYaml = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: noImagePullPolicy
  namespace: noImagePullPolicy
spec:
  template:
    spec:
      containers:
        - name: noImagePullPolicyContainer
`
var SSNoImagePullPolicyYaml = `apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: noImagePullPolicy
  namespace: noImagePullPolicy
spec:
  template:
    spec:
      containers:
        - name: noImagePullPolicyContainer
`