package update

import (
	"fmt"
	"github.com/entigolabs/entigo-k8s-gitops/internal/app/gitops/common"
	"github.com/entigolabs/entigo-k8s-gitops/internal/app/gitops/git"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type installInput struct {
	changeData []string
	fileNames  []string
}

func getInstallInput(repo *git.Repository) string {
	input := installInput{changeData: strings.Split(repo.Images, ",")}
	input.fileNames = getFileNames(repo.Recursive)
	return composeInstallInput(input)
}

func getFileNames(recursively bool) []string {
	if recursively {
		return getFilesRecursively()
	}
	return getFiles()
}

func getFilesRecursively() []string {
	var yamlNames []string
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(info.Name(), ".yaml") {
				yamlNames = append(yamlNames, path)
			}
			return nil
		})
	if err != nil {
		common.Logger.Println(common.PrefixedError{Reason: err})
	}
	return yamlNames
}

func getFiles() []string {
	yamlNames, err := readDirFiltered(".", ".yaml")
	if err != nil {
		common.Logger.Println(common.PrefixedError{Reason: err})
	}
	return yamlNames
}

func composeInstallInput(input installInput) string {
	composedInput := ""
	for _, d := range input.changeData {
		yamlNamesStr := strings.Join(input.fileNames, ",")
		editLine := fmt.Sprintf("edit %s %s %s", yamlNamesStr, getUpdateInstallLocations(), d)
		composedInput += fmt.Sprintf("%s\n", editLine)
	}
	return composedInput
}

func getUpdateInstallLocations() string {
	deploymentStatefulSetDaemonSetJobLocation := "spec.template.spec.containers.*.image"
	podLocation := "spec.containers.*.image"
	cronJobLocation := "spec.jobTemplate.spec.template.spec.containers.*.image"
	return strings.Join([]string{deploymentStatefulSetDaemonSetJobLocation, podLocation, cronJobLocation}, ",")
}

func readDirFiltered(path string, suffix string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, file := range files {
		if file.IsDir() == false && strings.HasSuffix(file.Name(), suffix) {
			result = append(result, file.Name())
		}
	}
	return result, err
}
