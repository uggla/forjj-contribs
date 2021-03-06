package main

import (
	"github.com/forj-oss/goforjj"
	"log"
	"os"
	"path"
)

// Return ok if the jenkins instance exist
func (r *MaintainReq) check_source_existence(ret *goforjj.PluginData) (status bool) {
	log.Print("Checking Jenkins source code path existence.")

	if _, err := os.Stat(r.Forj.ForjjSourceMount); err != nil {
		ret.Errorf("Unable to maintain jenkins instances. '%s' is inexistent or innacessible. %s", r.Forj.ForjjSourceMount, err)
		return
	}

	ret.StatusAdd("environment checked.")
	status = true
	return
}

// Looping on all instances
// Need to define where to deploy (dev/itg/pro/local/other)
func (r *MaintainReq) InstantiateAll(ret *goforjj.PluginData) (_ bool) {
	instance := r.Forj.ForjjInstanceName
	mount := r.Forj.ForjjSourceMount
	auths := NewDockerAuths(r.Objects.App[instance].RegistryAuth)

	src := path.Join(mount, instance)
	if _, err := os.Stat(path.Join(src, jenkins_file)); err == nil {
		p := new_plugin(src)
		ret.StatusAdd("Maintaining '%s'", instance)
		if err := os.Chdir(src); err != nil {
			ret.Errorf("Unable to enter in '%s'. %s", src, err)
			return
		}
		if !p.InstantiateInstance(instance, auths, ret) {
			return false
		}
	} else {
		log.Printf("'%s' is not a forjj plugin source code model. No '%s' found. ignored.", src, jenkins_file)
	}
	return true
}

func (p *JenkinsPlugin) InstantiateInstance(instance string, auths *DockerAuths, ret *goforjj.PluginData) (status bool) {
	if !p.load_yaml(ret) {
		return
	}

	// start a command as described by the source code.
	if p.yaml.Deploy.Command == "" {
		log.Printf(ret.Errorf("Unable to instantiate to %s. Deploy Command is empty.", p.yaml.Deploy.Deployment.To))
		return
	}

	for server := range auths.Auths {
		if err := auths.authenticate(server); err != nil {
			log.Printf(ret.Errorf("Unable to instantiate. %s", err))
			return
		}
	}

	log.Printf(ret.StatusAdd("Running '%s'", p.yaml.Deploy.Command))

	var env []string
	if v := os.Getenv("DOOD_SRC"); v != "" {
		env = append(os.Environ(), "SRC="+path.Join(v, instance))
	}

	s, err := run_cmd("/bin/sh", env, "-c", p.yaml.Deploy.Command)
	log.Printf(ret.StatusAdd(string(s)))
	if err != nil {
		cur_dir, _ := os.Getwd()
		log.Printf(ret.Errorf("%s (pwd: %s)", err, cur_dir))
	}

	return true
}
