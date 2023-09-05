#!/bin/sh

set -e

current_version=$(sed -n -e 's/version:[ "]*\([^"]*\).*/\1/p' $(dirname $0)/plugin.yaml)
CHART_VERSION_VERSION=${CHART_VERSION_VERSION:-$current_version}

dir=${HELM_PLUGIN_DIR:-"$(helm home)/plugins/chart-version"}
os=$(uname -s | tr '[:upper:]' '[:lower:]')
release_file="chart-version_${os}_${CHART_VERSION_VERSION}.tar.gz"
url="https://github.com/muandane/chart-version/releases/download/v${CHART_VERSION_VERSION}/${release_file}"

mkdir -p $dir

if command -v wget
then
  wget -O ${dir}/${release_file} $url
elif command -v curl; then
  curl -L -o ${dir}/${release_file} $url
fi

tar xvf ${dir}/${release_file} -C $dir

chmod +x ${dir}/chart-version

rm ${dir}/${release_file}