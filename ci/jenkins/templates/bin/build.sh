#!/bin/sh -e
#
#

IMAGE_NAME={{ .JenkinsImage.Name }}

if [ "$LOGNAME" = jenkins ]
then
   REPO={{ .JenkinsImage.RegistryRepoName }}
   IMAGE_VERSION={{ .JenkinsImage.Version }}
else
   REPO=$LOGNAME
   IMAGE_VERSION=test
fi

if [ -f build_opts.sh ]
then
   source build_opts.sh
fi

TAG_NAME={{ .JenkinsImage.RegistryServer }}/$REPO/$IMAGE_NAME:$IMAGE_VERSION

if [ "$http_proxy" != "" ]
then
   PROXY=" --build-arg http_proxy=$http_proxy --build-arg https_proxy=$https_proxy --build-arg no_proxy=$no_proxy"
   echo "Using your local proxy setting : $http_proxy"
   if [ "$no_proxy" != "" ]
   then
      PROXY="$PROXY --build-arg no_proxy=$no_proxy"
      echo "no_proxy : $no_proxy"
   fi
fi

if [ -z "$MYFORK" ]
then
   MYFORK="forj-oss/jenkins-install-inits"
   echo "Using default Organisation/repo ($MYFORK) for jenkins-install-inits. Add MYFORK= to change it."
fi

if [ -z "$BRANCH" ]
then
   BRANCH=master
   echo "Using current git branch 'master'. Add BRANCH= to change it."
fi

JENKINS_INSTALL_INITS_URL="https://github.com/$MYFORK/raw/$BRANCH/"
FEATURES="--build-arg JENKINS_INSTALL_INITS_URL=$JENKINS_INSTALL_INITS_URL"

set -x
sudo -n docker pull {{ .Dockerfile.FromImage }}{{ if .Dockerfile.FromImageVersion }}:{{ .Dockerfile.FromImageVersion }}{{ end }}
sudo -n docker build -t $TAG_NAME $FEATURES $PROXY $BUILD_OPTS .
set +x


if [ "$AUTO_PUSH" = true ]
then
   set -x
   sudo -n docker push $TAG_NAME
fi
