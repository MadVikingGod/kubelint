# Klint rules

- [x] NakedPodCheck -  Don’t use naked Pods - https://kubernetes.io/docs/concepts/configuration/overview/#naked-pods-vs-replicasets-deployments-and-jobs
- [ ] PodHostPort - Don’t specify a hostPort for a Pod - https://kubernetes.io/docs/concepts/configuration/overview/#services
- [ ] PodHostNetwork - Don’t specify a hostNetwork for a Pod - https://kubernetes.io/docs/concepts/configuration/overview/#services
- [ ] MissingLabels - There should be a minimal set of labels - This should be configurable, with sane defaults
- [x] ImageTagLatest - container images should not be Latest - https://kubernetes.io/docs/concepts/configuration/overview/#container-images
- [x] ImageTagMissing - container images should have a tag -
- [x] ImagePullPolicyAlways - Warn - The image pull policy should be Always - https://kubernetes.io/docs/concepts/configuration/overview/#container-images
- [ ] PodResourcesMissing - Pods should be configured with resources - https://cloud.google.com/blog/products/gcp/kubernetes-best-practices-resource-requests-and-limits
- [ ] PodReadinessMissing - Pods should be configured with readiness checks - https://cloud.google.com/blog/products/gcp/kubernetes-best-practices-setting-up-health-checks-with-readiness-and-liveness-probes
- [ ] PodLivenessMissing - Pods should be configure with liveness checks - https://cloud.google.com/blog/products/gcp/kubernetes-best-practices-setting-up-health-checks-with-readiness-and-liveness-probes
- [X] DeprecatedApi - A number of APIs have been deprecated, should move to newer versions -  https://kubernetes.io/blog/2019/07/18/api-deprecations-in-1-16/


## How would this work?
for each yaml document:

1. marshal into unstructured
1. Use GVK to sort tests
1. run relevant tests
1. collect warnings/errors
1. return error if any test is critical

## Unanswered questions:
- How does someone add an external lint rule?  Should this be supported?
- Would it be easier, and more dry to marshal into proper type?  
  - If so, how do we handel non-native types, eg. CRDs.  
  - Do we need to?
  
## explore nDepend as a potential for custom rules.

## explore CRD as a way to define rules.