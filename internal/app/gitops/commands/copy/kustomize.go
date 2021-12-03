// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"os"
	"strconv"
	"text/template"
    "io/ioutil"

    "github.com/entigolabs/entigo-k8s-gitops/internal/app/gitops/common"
	"sigs.k8s.io/kustomize/kyaml/errors"
	"sigs.k8s.io/kustomize/kyaml/fn/framework"
	"sigs.k8s.io/kustomize/kyaml/kio/filters"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

// define the input API schema as a struct
type API struct {
	Metadata *common.AppFlags

	Files struct {
        Deployment  string
        Service     string
	}
}

func installViaKustomize(flags *common.Flags) {
	functionConfig := &API{}
	resourceList := &framework.ResourceList{FunctionConfig: functionConfig}

	cmd := framework.Command(resourceList, func() error {
		// initialize API defaults
		if err := initAPI(functionConfig); err != nil {
			return err
		}

		// execute the service template
		buff := &bytes.Buffer{}
		 b, err := ioutil.ReadFile(functionConfig.Files.Service) // just pass the file name
            if err != nil {
                fmt.Print(err)
         }
        serviceTemplate := string(b) // convert content to a 'string'

		t := template.Must(template.New(fmt.Sprintf("%s-%s-service",flags.App.Name, flags.App.Branch)).Parse(serviceTemplate))
		if err := t.Execute(buff, functionConfig); err != nil {
			return err
		}
		s, err := yaml.Parse(buff.String())
		if err != nil {
			return err
		}

		// execute the deployment template
		buff = &bytes.Buffer{}
         b, err := ioutil.ReadFile(functionConfig.Files.Deployment) // just pass the file name
            if err != nil {
                fmt.Print(err)
         }
        deploymentTemplate := string(b) // convert content to a 'string'

		t = template.Must(template.New(fmt.Sprintf("%s-%s-deployment",flags.App.Name, flags.App.Branch)).Parse(deploymentTemplate))
		if err := t.Execute(buff, functionConfig); err != nil {
			return err
		}
		d, err := yaml.Parse(buff.String())
		if err != nil {
			return err
		}

		// add the template generated Resources to the output -- these will get merged by the next
		// filter
		resourceList.Items = append(resourceList.Items, s, d)

		// merge the new copies with the old copies of each resource
		resourceList.Items, err = filters.MergeFilter{}.Filter(resourceList.Items)
		if err != nil {
			return err
		}

		// apply formatting
		resourceList.Items, err = filters.FormatFilter{}.Filter(resourceList.Items)
		if err != nil {
			return err
		}

		return nil
	})
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func initAPI(api *API, flags *common.Flags) error {

	if api.Metadata.Name == "" {
		return errors.Errorf("must specify metadata.name\n")
	}

	deploymentFile := fmt.Sprintf("%s/%s",flags.App.KustomizeTemplatePath,flags.App.KustomizeDeployFile)
	if _, err := os.Stat(deploymentFile); err == nil {
        api.Files.Deployment = deploymentFile
    } else if errors.Is(err, os.ErrNotExist) {
        return errors.Errorf(fmt.Sprintf("%s is missing\n", deploymentFile))
    } else {
        return errors.Errorf(fmt.Sprintf("%s is missing\n", deploymentFile))
    }

    serviceFile := fmt.Sprintf("%s/%s",flags.App.KustomizeTemplatePath,flags.App.KustomizeServiceFile)
	if _, err := os.Stat(serviceFile); err == nil {
        api.Files.Service = serviceFile
    } else if errors.Is(err, os.ErrNotExist) {
        return errors.Errorf(fmt.Sprintf("%s is missing\n", serviceFile))
    } else {
        return errors.Errorf(fmt.Sprintf("%s is missing\n", serviceFile))
    }

	return nil
}
