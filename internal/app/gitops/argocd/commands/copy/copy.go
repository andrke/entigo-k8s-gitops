package copy

import (
	"errors"
	"fmt"
	"github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"
	"github.com/entigolabs/entigo-k8s-gitops/internal/app/gitops/argocd/api"
	"github.com/entigolabs/entigo-k8s-gitops/internal/app/gitops/argocd/commands/sync"
	argoCDUpdate "github.com/entigolabs/entigo-k8s-gitops/internal/app/gitops/argocd/commands/update"
	gitCopy "github.com/entigolabs/entigo-k8s-gitops/internal/app/gitops/commands/copy"
	"github.com/entigolabs/entigo-k8s-gitops/internal/app/gitops/common"
)

func Run(flags *common.Flags) {
	appName := flags.App.Name
	waitFailure := flags.ArgoCD.WaitFailure
	if flags.App.Branch != flags.App.SourceBranch {
		appName = appName + "-" + flags.App.Branch
	}
	client := api.NewClientOrDie(flags)
	app := client.GetRequest(appName, flags.ArgoCD.Timeout)
	if app == nil {
		copyFromSource(flags, client)
		flags.ArgoCD.WaitFailure = waitFailure
	}
	flags.App.Name = appName
	argoCDUpdate.Run(flags)
}

func copyFromSource(flags *common.Flags, client api.Client) {
	app := client.GetRequest(flags.App.Name, flags.ArgoCD.Timeout)
	if app == nil {
		common.Logger.Fatal(&common.PrefixedError{Reason: errors.New(fmt.Sprintf("App '%s' not found", flags.App.Name))})
	}
	modifyFlags(flags, app)
	gitCopy.Run(flags)
	flags.ArgoCD.WaitFailure = false
	sync.Run(flags)
}

func modifyFlags(flags *common.Flags, app *v1alpha1.Application) {
	flags.Git.Repo = app.Spec.Source.RepoURL
	flags.Git.Branch = app.Spec.Source.TargetRevision
	flags.App.Path = app.Spec.Source.Path
	flags.App.Name = app.Name
}
