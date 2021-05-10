module github.com/tom721/from-helm-server

go 1.15

require (
	github.com/fluxcd/helm-operator v1.2.0
	github.com/gorilla/mux v1.8.0
	k8s.io/apimachinery v0.21.0
	k8s.io/client-go v11.0.0+incompatible
	sigs.k8s.io/controller-runtime v0.8.3
)

replace (
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
	github.com/docker/docker => github.com/moby/moby v1.4.2-0.20200203170920-46ec8731fbce
	github.com/fluxcd/flux => github.com/fluxcd/flux v1.19.0
	github.com/fluxcd/flux/pkg/install => github.com/fluxcd/flux/pkg/install v0.0.0-20200402061723-01a239a69319
	github.com/fluxcd/helm-operator/pkg/install => github.com/fluxcd/helm-operator/pkg/install v0.0.0-20200407140510-8d71b0072a3e
)

replace (
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.4.0
	k8s.io/api => k8s.io/api v0.18.8
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.8
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.8
	k8s.io/client-go => k8s.io/client-go v0.18.8
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20200410145947-bcb3869e6f29
)
