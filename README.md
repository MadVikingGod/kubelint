# kubelint

Kubernetes best practices, and dryer lint.

This is an experiment attempting to codify best practices around kubernetes.  The goal is to run `kustomize build | kubelint` or `helm template... | kubelint` and get back a list of problems with your manifests.
