package cli

import (
	"github.com/entigolabs/entigo-k8s-gitops/internal/app/gitops/common"
	"github.com/urfave/cli/v2"
	"strconv"
)

func cliFlags(cmd common.Command) []cli.Flag {
	var flags []cli.Flag
	flags = appendBaseFlags(flags)
	flags = appendCmdSpecificFlags(flags, cmd)
	return flags
}

func appendBaseFlags(flags []cli.Flag) []cli.Flag {
	return append(flags,
		&loggingFlag,
		&gitRepoFlag,
		&gitBranchFlag,
		&gitKeyFileFlag,
		&gitStrictHostKeyCheckingFlag,
		&gitPushFlag,
		&gitAuthorNameFlag,
		&gitAuthorEmailFlag,
		&appPrefixFlag,
		&appNamespaceFlag,
		&appNameFlag)
}

func appendCmdSpecificFlags(baseFlags []cli.Flag, cmd common.Command) []cli.Flag {
	switch cmd {
	case common.RunCmd:
	case common.UpdateCmd:
		baseFlags = updateSpecificFlags(baseFlags)
	case common.CopyCmd:
		baseFlags = copyAndDeleteSpecificFlags(baseFlags)
	case common.DeleteCmd:
		baseFlags = copyAndDeleteSpecificFlags(baseFlags)
	}
	return baseFlags
}

func updateSpecificFlags(baseFlags []cli.Flag) []cli.Flag {
	baseFlags = append(baseFlags, &imagesFlag)
	baseFlags = append(baseFlags, &keepRegistryFlag)
	baseFlags = append(baseFlags, &appPathFlag)
	baseFlags = append(baseFlags, &deploymentStrategyFlag)
	baseFlags = append(baseFlags, &recursiveFlag)
	return baseFlags
}

func copyAndDeleteSpecificFlags(baseFlags []cli.Flag) []cli.Flag {
	baseFlags = append(baseFlags, &appBranchFlag)
	baseFlags = append(baseFlags, &appPrefixArgoFlag)
	baseFlags = append(baseFlags, &appPrefixYamlFlag)
    baseFlags = append(baseFlags, &appDomainFlag)
    baseFlags = append(baseFlags, &appPathFlag)
    baseFlags = append(baseFlags, &appArgoAppsPathFlag)
    baseFlags = append(baseFlags, &appUsePrefixedPathFlag)
    baseFlags = append(baseFlags, &appUseKustomizeFlag)
    baseFlags = append(baseFlags, &appKustomizeTemplatePathFlag)
    baseFlags = append(baseFlags, &appKustomizeServiceFileFlag)
    baseFlags = append(baseFlags, &appKustomizeDeployFileFlag)
	return baseFlags
}

var loggingFlag = cli.StringFlag{
	Name:        "logging",
	Aliases:     []string{"l"},
	EnvVars:     []string{"LOGGING"},
	DefaultText: "prod",
	Value:       "prod",
	Usage:       "set `logging level` (prod | dev)",
	Destination: &flags.LoggingLevel,
}

var gitRepoFlag = cli.StringFlag{
	Name:        "git-repo",
	EnvVars:     []string{"GIT_REPO"},
	DefaultText: "",
	Usage:       "Git repository `SSH URL`",
	Destination: &flags.Git.Repo,
	Required:    true,
}

var gitBranchFlag = cli.StringFlag{
	Name:        "git-branch",
	EnvVars:     []string{"GIT_BRANCH"},
	DefaultText: "",
	Usage:       "branch `name`",
	Destination: &flags.Git.Branch,
	Required:    true,
}

var gitKeyFileFlag = cli.StringFlag{
	Name:        "git-key-file",
	EnvVars:     []string{"GIT_KEY_FILE"},
	DefaultText: "",
	Usage:       "SSH private key `location`",
	Destination: &flags.Git.KeyFile,
	Required:    true,
}

var gitStrictHostKeyCheckingFlag = cli.BoolFlag{
	Name:        "git-strict-host-key-checking",
	EnvVars:     []string{"GIT_STRICT_HOST_KEY_CHECKING"},
	DefaultText: strconv.FormatBool(false),
	Usage:       "strict host key checking",
	Destination: &flags.Git.StrictHostKeyChecking,
}

var gitPushFlag = cli.BoolFlag{
	Name:        "git-push",
	EnvVars:     []string{"GIT_PUSH"},
	DefaultText: strconv.FormatBool(true),
	Value:       true,
	Usage:       "push changes",
	Destination: &flags.Git.Push,
}

var gitAuthorNameFlag = cli.StringFlag{
	Name:        "git-author-name",
	EnvVars:     []string{"GIT_AUTHOR_NAME"},
	DefaultText: "jenkins",
	Value:       "jenkins",
	Usage:       "Git author name",
	Destination: &flags.Git.AuthorName,
}

var gitAuthorEmailFlag = cli.StringFlag{
	Name:        "git-author-email",
	EnvVars:     []string{"GIT_AUTHOR_EMAIL"},
	DefaultText: "jenkins@localhost",
	Value:       "jenkins@localhost",
	Usage:       "Git author email",
	Destination: &flags.Git.AuthorEmail,
}

var appPathFlag = cli.StringFlag{
	Name:        "app-path",
	EnvVars:     []string{"APP_PATH"},
	DefaultText: "",
	Usage:       "`path` to application folder",
	Destination: &flags.App.Path,
}

var appArgoAppsPathFlag = cli.StringFlag{
	Name:        "app-argoapps-path",
	EnvVars:     []string{"APP_ARGOAPPS_PATH"},
	DefaultText: "",
	Usage:       "`path` to argoapps folder",
	Destination: &flags.App.ArgoAppsPath,
}

var appKustomizeTemplatePathFlag = cli.StringFlag{
	Name:        "app-kustomize-template-path",
	EnvVars:     []string{"APP_KUSTOMIZE_TEMPLATE_PATH"},
	DefaultText: "",
	Usage:       "`path` to kustomize template folder",
	Destination: &flags.App.KustomizeTemplatePath,
}

var appKustomizeServiceFileFlag = cli.StringFlag{
	Name:        "app-kustomize-service-template",
	EnvVars:     []string{"APP_KUSTOMIZE_SERVICE_TEMPLATE"},
	DefaultText: "service.yaml",
	Usage:       "filename of kustomize service template",
	Destination: &flags.App.KustomizeServiceFile,
}

var appKustomizeDeployFileFlag = cli.StringFlag{
	Name:        "app-kustomize-deployment-template",
	EnvVars:     []string{"APP_KUSTOMIZE_DEPLOYMENT_TEMPLATE"},
	DefaultText: "deployment.yaml",
	Usage:       "filename of kustomize deployment template",
	Destination: &flags.App.KustomizeDeployFile,
}
var appUsePrefixedPathFlag = cli.BoolFlag{
	Name:        "app-use-prefixed-path",
	EnvVars:     []string{"APP_USE_PREFIXED_PATH"},
	DefaultText: strconv.FormatBool(true),
	Value:       true,
	Usage:       "Use prefixed path",
	Destination: &flags.App.UsePrefixedPath,
}

var appUseKustomizeFlag = cli.BoolFlag{
	Name:        "app-use-kustomize",
	EnvVars:     []string{"APP_USE_KUSTOMIZE"},
	DefaultText: strconv.FormatBool(false),
	Value:       false,
	Usage:       "Use kustomize to copy app",
	Destination: &flags.App.UseKustomize,
}

var appPrefixFlag = cli.StringFlag{
	Name:        "app-prefix",
	EnvVars:     []string{"APP_PREFIX"},
	DefaultText: "",
	Usage:       "`path` prefix to apply",
	Destination: &flags.App.Prefix,
}

var appNamespaceFlag = cli.StringFlag{
	Name:        "app-namespace",
	EnvVars:     []string{"APP_NAMESPACE"},
	DefaultText: "",
	Usage:       "application namespace `name`",
	Destination: &flags.App.Namespace,
}

var appNameFlag = cli.StringFlag{
	Name:        "app-name",
	EnvVars:     []string{"APP_NAME"},
	DefaultText: "",
	Usage:       "application name",
	Destination: &flags.App.Name,
}

var appBranchFlag = cli.StringFlag{
	Name:        "app-branch",
	EnvVars:     []string{"APP_BRANCH"},
	DefaultText: "",
	Usage:       "application branch `name`",
	Destination: &flags.App.Branch,
	Required:    true,
}

var appDomainFlag = cli.StringFlag{
	Name:        "app-domain",
	EnvVars:     []string{"APP_DOMAIN"},
	DefaultText: "localhost",
	Usage:       "application domain",
	Destination: &flags.App.Domain,
}

var appPrefixArgoFlag = cli.StringFlag{
	Name:        "app-prefix-argo",
	EnvVars:     []string{"APP_PREFIX_ARGO"},
	DefaultText: "argoapps",
	Value:       "argoapps",
	Usage:       "Argo app `path`",
	Destination: &flags.App.PrefixArgo,
}

var appPrefixYamlFlag = cli.StringFlag{
	Name:        "app-prefix-yaml",
	EnvVars:     []string{"APP_PREFIX_YAML"},
	DefaultText: "yaml",
	Value:       "yaml",
	Usage:       "yaml configurations `path`",
	Destination: &flags.App.PrefixYaml,
}

var imagesFlag = cli.StringFlag{
	Name:        "images",
	Aliases:     []string{"i"},
	EnvVars:     []string{"IMAGES"},
	DefaultText: "",
	Usage:       "images with tags",
	Destination: &flags.Images,
	Required:    true,
}

var keepRegistryFlag = cli.BoolFlag{
	Name:        "keep-registry",
	Aliases:     []string{"k"},
	EnvVars:     []string{"KEEP_REGISTRY"},
	DefaultText: strconv.FormatBool(false),
	Usage:       "keeps registry part of the changeable image",
	Destination: &flags.KeepRegistry,
}

var deploymentStrategyFlag = cli.StringFlag{
	Name:        "deployment-strategy",
	Aliases:     []string{"s"},
	EnvVars:     []string{"DEPLOYMENT-STRATEGY"},
	DefaultText: "if not defined then strategy will remain unchanged",
	Usage:       "change deployment strategy (RollingUpdate | Recreate)",
	Destination: &flags.DeploymentStrategy,
}

var recursiveFlag = cli.BoolFlag{
	Name:        "recursive",
	EnvVars:     []string{"RECURSIVE"},
	DefaultText: strconv.FormatBool(false),
	Usage:       "updates directories and their contents recursively",
	Destination: &flags.Recursive,
}
