// This file is autogenerated by "go generate". Do not modify it.
// It has been generated from your 'jenkins.yaml' file.
// To update those structure, update the 'jenkins.yaml' and run 'go generate'
package main

// Object app groups structure

// Groups structure

type DeployStruct struct {
	ServiceAddr string `json:"deploy-service-addr"` // Exposed service CNAME or IP address of the expected jenkins instance
	ServicePort string `json:"deploy-service-port"` // Expected jenkins instance port number.
	To          string `json:"deploy-to"`           // Deployment name used to deploy jenkins.
}

type DockerfileStruct struct {
	FromImage        string `json:"dockerfile-from-image"`         // Base Docker image tag name to use in Dockerfile. Must respect [server/repo/]name.
	FromImageVersion string `json:"dockerfile-from-image-version"` // Base Docker image tag version to use in Dockerfile
	Maintainer       string `json:"dockerfile-maintainer"`         // Jenkins image maintainer
}

type FinalImageStruct struct {
	Name             string `json:"final-image-name"`               // Docker image name for your final generated Jenkins Image. Do not set the Server or Repo name. Use final-docker-registry-server and final-docker-repo-name.
	RegistryRepoName string `json:"final-image-registry-repo-name"` // Docker Repository Name where your image will be pushed. If not set, no push will be done.
	RegistryServer   string `json:"final-image-registry-server"`    // Docker registry server name where your image will be pushed. If not set, no push will be done.
	Version          string `json:"final-image-version"`            // Docker image tag version for your generated Jenkins Image.
}

// Object Instance structures

type AppInstanceStruct struct {
	RegistryAuth string `json:"registry-auth"` // List of Docker registry servers authentication separated by coma. One registry server auth string is build as <server>:<token>[:<email>]
	SeedJobRepo  string `json:"seed-job-repo"` // url to the seed job repository. By default, it uses the <YourInfraRepo>. Jobs are defined under job-dsl.

	// Groups

	DeployStruct
	DockerfileStruct
	FinalImageStruct
}

// Object features groups structure

// Groups structure

// Object Instance structures

type FeaturesInstanceStruct struct {
	Name    string `json:"name"`    // name of the jenkins feature
	Options string `json:"options"` // List of feature option to use
}

// Object projects groups structure

// Groups structure

type GitStruct struct {
	RemoteUrl string `json:"git-remote-url"` // with remote-type = 'git', Remote repository url.
}

type GithubStruct struct {
	ApiUrl    string `json:"github-api-url"`    // with remote-type = 'github', Github API Url. By default, it uses public github API.
	Repo      string `json:"github-repo"`       // with remote-type = 'github', Repository name.
	RepoOwner string `json:"github-repo-owner"` // with remote-type = 'github', Repository owner. Can be a user or an organization.
}

// Object Instance structures

type ProjectsInstanceStruct struct {
	Name       string `json:"name"`        // Project name
	RemoteType string `json:"remote-type"` // Define remote source  type. 'github' is used by default. Support 'git', 'github'.

	// Groups

	GitStruct
	GithubStruct
}

// ************************
// Create request structure
// ************************

type CreateReq struct {
	Forj struct {
		Debug              string `json:"debug"`
		ForjjInfra         string `json:"forjj-infra"`
		ForjjInfraUpstream string `json:"forjj-infra-upstream"`
		ForjjInstanceName  string `json:"forjj-instance-name"`
		ForjjOrganization  string `json:"forjj-organization"`
		ForjjSourceMount   string `json:"forjj-source-mount"`
	}
	Objects CreateArgReq
}

type CreateArgReq struct {
	App      map[string]AppInstanceStruct      `json:"app"`      // Object details
	Features map[string]FeaturesInstanceStruct `json:"features"` // Object details
	Projects map[string]ProjectsInstanceStruct `json:"projects"` // Object details
}

// ************************
// Update request structure
// ************************

type UpdateReq struct {
	Forj struct {
		Debug              string `json:"debug"`
		ForjjInfra         string `json:"forjj-infra"`
		ForjjInfraUpstream string `json:"forjj-infra-upstream"`
		ForjjInstanceName  string `json:"forjj-instance-name"`
		ForjjOrganization  string `json:"forjj-organization"`
		ForjjSourceMount   string `json:"forjj-source-mount"`
	}
	Objects UpdateArgReq
}

type UpdateArgReq struct {
	App      map[string]AppInstanceStruct      `json:"app"`      // Object details
	Features map[string]FeaturesInstanceStruct `json:"features"` // Object details
	Projects map[string]ProjectsInstanceStruct `json:"projects"` // Object details
}

// **************************
// Maintain request structure
// **************************

type MaintainReq struct {
	Forj struct {
		Debug              string `json:"debug"`
		ForjjInfra         string `json:"forjj-infra"`
		ForjjInfraUpstream string `json:"forjj-infra-upstream"`
		ForjjInstanceName  string `json:"forjj-instance-name"`
		ForjjOrganization  string `json:"forjj-organization"`
		ForjjSourceMount   string `json:"forjj-source-mount"`
		DeployTo           string `json:"deploy-to"`
	}
	Objects MaintainArgReq
}

type MaintainArgReq struct {
	App map[string]AppMaintainStruct `json:"app"` // Object details
}

type AppMaintainStruct struct {
	RegistryAuth string `json:"registry-auth"` // List of Docker registry servers authentication separated by coma. One registry server auth string is build as <server>:<token>[:<email>]
}

// YamlDesc has been created from your 'jenkins.yaml' file.
const YamlDesc = "---\n" +
	"plugin: \"jenkins\"\n" +
	"version: \"0.1\"\n" +
	"description: \"CI jenkins plugin for FORJJ.\"\n" +
	"runtime:\n" +
	"  docker:\n" +
	"    image: \"forjdevops/forjj-jenkins\"\n" +
	"    dood: true\n" +
	"  service_type: \"REST API\"\n" +
	"  service:\n" +
	"    #socket: \"jenkins.sock\"\n" +
	"    parameters: [ \"service\", \"start\", \"--templates\", \"/templates\"]\n" +
	"created_flag_file: \"{{ .InstanceName }}/forjj-{{ .Name }}.yaml\"\n" +
	"task_flags:\n" +
	"  common:\n" +
	"    forjj-infra-upstream:\n" +
	"      help: \"address of the infra repository upstream\"\n" +
	"    forjj-infra:\n" +
	"      help: \"Name of the Infra repository to use\"\n" +
	"    forjj-instance-name:\n" +
	"    forjj-source-mount:\n" +
	"    forjj-organization:\n" +
	"    debug:\n" +
	"      help: \"To activate jenkins debug information\"\n" +
	"    forjj-source-mount:\n" +
	"      help: \"Where the source dir is located for jenkins plugin.\"\n" +
	"  maintain:\n" +
	"    deploy-to:\n" +
	"      default: docker\n" +
	"      help: \"Where jenkins will be published.\"\n" +
	"objects:\n" +
	"  app:\n" +
	"    default-actions: [\"add\", \"change\"]\n" +
	"    groups:\n" +
	"      dockerfile:\n" +
	"        flags:\n" +
	"          # Information we can define for the Dockerfile.\n" +
	"          from-image:\n" +
	"            help: \"Base Docker image tag name to use in Dockerfile. Must respect [server/repo/]name.\"\n" +
	"            default: forjdevops/jenkins-dood\n" +
	"          from-image-version:\n" +
	"            help: \"Base Docker image tag version to use in Dockerfile\"\n" +
	"          maintainer:\n" +
	"            help: \"Jenkins image maintainer\"\n" +
	"      final-image:\n" +
	"        flags:\n" +
	"          name:\n" +
	"            help: \"Docker image name for your final generated Jenkins Image. Do not set the Server or Repo name. Use final-docker-registry-server and final-docker-repo-name.\"\n" +
	"            default: jenkins\n" +
	"          version:\n" +
	"            help: \"Docker image tag version for your generated Jenkins Image.\"\n" +
	"          registry-server:\n" +
	"            help: \"Docker registry server name where your image will be pushed. If not set, no push will be done.\"\n" +
	"            default: hub.docker.com\n" +
	"          registry-repo-name:\n" +
	"            help: \"Docker Repository Name where your image will be pushed. If not set, no push will be done.\"\n" +
	"      deploy:\n" +
	"        flags:\n" +
	"          to:\n" +
	"            help: \"Deployment name used to deploy jenkins.\"\n" +
	"            default: \"docker\"\n" +
	"          service-addr:\n" +
	"            help: \"Exposed service CNAME or IP address of the expected jenkins instance\"\n" +
	"          service-port:\n" +
	"            default: 8080\n" +
	"            help: \"Expected jenkins instance port number.\"\n" +
	"    flags:\n" +
	"      seed-job-repo:\n" +
	"        help: \"url to the seed job repository. By default, it uses the <YourInfraRepo>. Jobs are defined under job-dsl.\"\n" +
	"        default: \"{{ .Forjj.InfraUpstream }}\"\n" +
	"      registry-auth:\n" +
	"        help: \"List of Docker registry servers authentication separated by coma. One registry server auth string is build as <server>:<token>[:<email>]\"\n" +
	"        secure: true\n" +
	"        envar: \"REGISTRY_AUTH\"\n" +
	"  features:\n" +
	"    default-actions: [\"add\", \"change\", \"remove\"]\n" +
	"    identified_by_flag: name\n" +
	"    flags:\n" +
	"      name:\n" +
	"        help: \"name of the jenkins feature\"\n" +
	"        required: true\n" +
	"      options:\n" +
	"        help: \"List of feature option to use\"\n" +
	"  projects:\n" +
	"    default-actions: [\"add\", \"change\", \"remove\"]\n" +
	"    identified_by_flag: name\n" +
	"    flags:\n" +
	"      name:\n" +
	"        help: \"Project name\"\n" +
	"        required: true\n" +
	"      remote-type:\n" +
	"        help: \"Define remote source  type. 'github' is used by default. Support 'git', 'github'.\"\n" +
	"    groups:\n" +
	"      github:\n" +
	"        flags:\n" +
	"          api-url:\n" +
	"            default: \"https://github.com\"\n" +
	"            help: \"with remote-type = 'github', Github API Url. By default, it uses public github API.\"\n" +
	"          repo-owner:\n" +
	"            help: \"with remote-type = 'github', Repository owner. Can be a user or an organization.\"\n" +
	"          repo:\n" +
	"            help: \"with remote-type = 'github', Repository name.\"\n" +
	"      git:\n" +
	"        flags:\n" +
	"          remote-url:\n" +
	"            help: \"with remote-type = 'git', Remote repository url.\"\n" +
	""
