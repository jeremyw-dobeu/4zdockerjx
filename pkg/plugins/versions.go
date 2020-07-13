package plugins

import (
	jenkinsv1 "github.com/jenkins-x/jx-api/pkg/apis/jenkins.io/v1"
)

const (
	// AdminVersion the version of the jx admin plugin
	AdminVersion = "0.0.39"

	// ApplicationVersion the version of the jx application plugin
	ApplicationVersion = "0.0.6"

	// GitOpsVersion the version of the jx gitops plugin
	GitOpsVersion = "0.0.68"

	// PipelineVersion the version of the jx pipeline plugin
	PipelineVersion = "0.0.4"

	// ProjectVersion the version of the jx project plugin
	ProjectVersion = "0.0.31"

	// PromoteVersion the version of the jx promote plugin
	PromoteVersion = "0.0.54"

	// SecretVersion the version of the jx secret plugin
	SecretVersion = "0.0.38"

	// VerifyVersion the version of the jx verify plugin
	VerifyVersion = "0.0.8"
)

var (
	// Plugins default plugins
	Plugins = []jenkinsv1.Plugin{
		CreateJXPlugin("admin", AdminVersion),
		CreateJXPlugin("application", ApplicationVersion),
		CreateJXPlugin("gitops", GitOpsVersion),
		CreateJXPlugin("pipeline", PipelineVersion),
		CreateJXPlugin("project", ProjectVersion),
		CreateJXPlugin("promote", PromoteVersion),
		CreateJXPlugin("secret", SecretVersion),
		CreateJXPlugin("verify", VerifyVersion),
	}
)