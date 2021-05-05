# GitOps Util

This is [gitops](https://www.gitops.tech/) helper CLI

* [Compiling Source](#compiling-source)
* [How to use](#how-to-use)
* [Examples](#examples)
* [ArgoCD Commands](#argocd-commands)
    * [argocd-get](#argocd-get)
    * [argocd-sync](#argocd-sync)
    * [argocd-update](#argocd-update)
    * [argocd-delete](#argocd-delete)
    * [argocd-copy](#argocd-copy)

## Compiling Source

- ```go build -o bin/gitops cmd/gitops/main.go```

## How to Use

This GitOps utility supports 8 commands:

- ```run``` - run update and copy logic
- ```update``` - updates specified images
- ```copy``` - copies from master branch to specified branch
- ```argocd-get``` - gets an ArgoCD application as json output
- ```argocd-sync``` - syncs an ArgoCD application
- ```argocd-update``` - updates an ArgoCD application, combines argocd-get, git update and argocd-sync
- ```argocd-delete``` - deletes an ArgoCD application
- ```argocd-copy``` - copies application to specified branch, combines argocd-get, git copy, git update and argocd-sync

## Examples

#### Update Command Example
##### Example With Application Path Flag
```
./gitops update \
    --git-repo=<repository-ssh-url> \
    --git-branch=<repository-branch> \
    --git-key-file=<key-file-location> \
    --images=<image:tag,image2:tag> \ 
    --app-path=<app-path>
```

##### Example With Tokenized Path Flags 

Tokenized path flags: 
1) ```--app-prefix=<app-prefix>``` 
2) ```--app-namespace=<app-namespace>```
3) ```--app-name=<app-name>```

```
./gitops update \
    --git-repo=<repository-ssh-url> \
    --git-branch=<repository-branch> \
    --git-key-file=<key-file-location> \
    --images=<image:tag,image2:tag> \ 
    --app-prefix=<app-prefix> \
    --app-namespace=<app-namespace> \
    --app-name=<app-name>
```
**Valuated ```--app-path``` with other than default value will override tokenized path flags**

## ArgoCD Commands

### argocd-get

Gets an ArgoCD application as json in standard output.

OPTIONS:
* **app-name** - **Required**, name of the ArgoCD application [$APP_NAME]
* **server** - **Required**, server tcp address with port [$ARGO_CD_SERVER]
* **auth-token** - **Required**, authentication token [$ARGO_CD_TOKEN]
* **insecure** - insecure connection (default: **false**) [$ARGO_CD_INSECURE]
* **timeout** - timeout for single ArgoCD operation (default: **300**) [$ARGO_CD_TIMEOUT]
* **grpc-web** - Enables gRPC-web protocol. Useful if Argo CD server is behind proxy which does not support HTTP2 (default: **false**) [$ARGO_CD_GRPC_WEB]

Minimal example

```./gitops argocd-get --app-name='application-name' --server='localhost:443' --auth-token='auth-token'```

Full example

```./gitops argocd-get --app-name='application-name' --server='localhost:443' --auth-token='auth-token' --timeout=300 --insecure=false --grpc-web=false```

### argocd-sync

Syncs an ArgoCD application and waits for it to reach synced and healthy status.

OPTIONS:
* **app-name** - **Required**, name of the ArgoCD application [$APP_NAME]
* **server** - **Required**, server tcp address with port [$ARGO_CD_SERVER]
* **auth-token** - **Required**, authentication token [$ARGO_CD_TOKEN]
* **insecure** - insecure connection (default: **false**) [$ARGO_CD_INSECURE]
* **timeout** - timeout for single ArgoCD operation (default: **300**) [$ARGO_CD_TIMEOUT]
* **grpc-web** - Enables gRPC-web protocol. Useful if Argo CD server is behind proxy which does not support HTTP2 (default: **false**) [$ARGO_CD_GRPC_WEB]
* **async** - don't wait for sync to complete (default: **false**) [$ARGO_CD_ASYNC]
* **wait-failure** - Fail the command when waiting for the sync to complete exceeds the timeout (default: **true**) [$ARGO_CD_WAIT_FAILURE]

Minimal example

```./gitops argocd-sync --app-name='application-name' --server='localhost:443' --auth-token='auth-token'```

Full example

```./gitops argocd-sync --app-name='application-name' --server='localhost:443' --auth-token='auth-token' --timeout=300 --insecure=false --grpc-web=false --async=false --wait-failure=true```

### argocd-update

Updates an ArgoCD application. Combines argocd-get, git update and argocd-sync commands.

OPTIONS:
* **app-name** - **Required**, name of the ArgoCD application [$APP_NAME]
* **server** - **Required**, server tcp address with port [$ARGO_CD_SERVER]
* **auth-token** - **Required**, authentication token [$ARGO_CD_TOKEN]
* **git-key-file** - **Required**, SSH private key location [$GIT_KEY_FILE]
* **images** [-i] - **Required**, images with tags, comma separated list [$IMAGES]
* **insecure** - insecure connection (default: **false**) [$ARGO_CD_INSECURE]
* **timeout** - timeout for single ArgoCD operation (default: **300**) [$ARGO_CD_TIMEOUT]
* **grpc-web** - Enables gRPC-web protocol. Useful if Argo CD server is behind proxy which does not support HTTP2 (default: **false**) [$ARGO_CD_GRPC_WEB]
* **git-strict-host-key-checking** - strict host key checking (default: **false**) [$GIT_STRICT_HOST_KEY_CHECKING]
* **git-push** - push changes (default: **true**) [$GIT_PUSH]
* **git-author-name value** - Git author name (default: **jenkins**) [$GIT_AUTHOR_NAME]
* **git-author-email value** - Git author email (default: **jenkins@localhost**) [$GIT_AUTHOR_EMAIL]
* **keep-registry** [-k] - keeps registry part of the changeable image (default: **false**) [$KEEP_REGISTRY]
* **deployment-strategy** [-s] - change deployment strategy (RollingUpdate | Recreate) (default: if not defined then strategy will remain unchanged) [$DEPLOYMENT-STRATEGY]
* **recursive** - updates directories and their contents recursively (default: **false**) [$RECURSIVE]
* **sync** - Sync the application after update (default: **true**) [$ARGO_CD_SYNC]
* **async** - don't wait for argoCD sync to complete (default: **false**) [$ARGO_CD_ASYNC]
* **wait-failure** - Fail the command when waiting for the sync to complete exceeds the timeout (default: **true**) [$ARGO_CD_WAIT_FAILURE]

Minimal example

```./gitops argocd-update --app-name='application-name' --server='localhost:443' --auth-token='auth-token' --git-token-file='file' --images='image:1'```

### argocd-delete

Deletes an ArgoCD application.

OPTIONS:
* **app-name** - **Required**, name of the ArgoCD application [$APP_NAME]
* **server** - **Required**, server tcp address with port [$ARGO_CD_SERVER]
* **auth-token** - **Required**, authentication token [$ARGO_CD_TOKEN]
* **insecure** - insecure connection (default: **false**) [$ARGO_CD_INSECURE]
* **timeout** - timeout for single ArgoCD operation (default: **300**) [$ARGO_CD_TIMEOUT]
* **grpc-web** - Enables gRPC-web protocol. Useful if Argo CD server is behind proxy which does not support HTTP2 (default: **false**) [$ARGO_CD_GRPC_WEB]
* **cascade** - perform a cascaded deletion of all application resources (default: **true**) [$ARGO_CD_CASCADE]

Minimal example

```./gitops argocd-delete --app-name='application-name' --server='localhost:443' --auth-token='auth-token'```

Full example

```./gitops argocd-sync --app-name='application-name' --server='localhost:443' --auth-token='auth-token' --timeout=300 --insecure=false --grpc-web=false --cascade=true```

### argocd-copy

Copies application to specified branch, combines argocd-get, conditional git copy, git update and argocd-sync.

OPTIONS:
* **app-name** - **Required**, name of the ArgoCD application [$APP_NAME]
* **server** - **Required**, server tcp address with port [$ARGO_CD_SERVER]
* **auth-token** - **Required**, authentication token [$ARGO_CD_TOKEN]
* **git-key-file** - **Required**, SSH private key location [$GIT_KEY_FILE]
* **app-branch** - **Required**, application branch name [$APP_BRANCH]
* **app-source-branch** - **Required**, application source branch name [$APP_SOURCE_BRANCH]
* **images** [-i] - **Required**, images with tags, comma separated list [$IMAGES]
* **insecure** - insecure connection (default: **false**) [$ARGO_CD_INSECURE]
* **timeout** - timeout for single ArgoCD operation (default: **300**) [$ARGO_CD_TIMEOUT]
* **grpc-web** - Enables gRPC-web protocol. Useful if Argo CD server is behind proxy which does not support HTTP2 (default: **false**) [$ARGO_CD_GRPC_WEB]
* **git-strict-host-key-checking** - strict host key checking (default: **false**) [$GIT_STRICT_HOST_KEY_CHECKING]
* **git-push** - push changes (default: **true**) [$GIT_PUSH]
* **git-author-name value** - Git author name (default: **jenkins**) [$GIT_AUTHOR_NAME]
* **git-author-email value** - Git author email (default: **jenkins@localhost**) [$GIT_AUTHOR_EMAIL]
* **app-prefix-argo** - ArgoCD app path (default: **argoapps**) [$APP_PREFIX_ARGO]
* **app-prefix-yaml** - yaml configurations path (default: **yaml**) [$APP_PREFIX_YAML]
* **keep-registry** [-k] - keeps registry part of the changeable image (default: **false**) [$KEEP_REGISTRY]
* **deployment-strategy** [-s] - change deployment strategy (RollingUpdate | Recreate) (default: if not defined then strategy will remain unchanged) [$DEPLOYMENT-STRATEGY]
* **recursive** - updates directories and their contents recursively (default: **false**) [$RECURSIVE]
* **sync** - Sync the application after update (default: **true**) [$ARGO_CD_SYNC]
* **async** - don't wait for argoCD sync to complete (default: **false**) [$ARGO_CD_ASYNC]
* **wait-failure** - Fail the command when waiting for the sync to complete exceeds the timeout (default: **true**) [$ARGO_CD_WAIT_FAILURE]

```./gitops argocd-update --app-name='application-name' --server='localhost:443' --auth-token='auth-token' --git-token-file='file' --app-branch='branch-1' --app-source-branch='master' --images='image:1'```