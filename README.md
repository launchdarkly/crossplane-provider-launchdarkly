# Provider LaunchDarkly

`provider-launchdarkly` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
LaunchDarkly API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/launchdarkly/provider-launchdarkly):

<!-- x-release-please-start-version -->

```
up ctp provider install xpkg.upbound.io/launchdarkly/provider-launchdarkly:v0.2.1
```

<!-- x-release-please-end -->

Alternatively, you can use declarative installation:

<!-- x-release-please-start-version -->

```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-launchdarkly
spec:
  package: xpkg.upbound.io/launchdarkly/provider-launchdarkly:v0.2.1
EOF
```

<!-- x-release-please-end -->

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/launchdarkly/crossplane-provider-launchdarkly).

## Developing

### Initial setup

```bash
make submodules
```

Run code-generation pipeline:

```console
make generate
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

### Installing Provider/CRDs into your local k8s cluster

1. Ensure Crossplane is installed on your local cluster (instructions [here](https://docs.crossplane.io/latest/software/install/))
2. Run the following:

   ```bash
   kubectl config use-context <name-of-your-local-k8s-context>
   kubectl apply -f ./package/crds
   ```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/launchdarkly/crossplane-provider-launchdarkly/issues).
