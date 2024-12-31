#!/usr/bin/env bash
# deploy.sh
#

set -e

env=$1

[[ "$env" != "test" ]] && [[ "$env" != "prod" ]] || [[ "$#" != 1 ]] && {
    echo "Usage: $0 <test|prod>"
    exit 1
}

SED=sed
[[ $(uname) = "Darwin" ]] && SED=gsed

script_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

project_name="devinggo"
app_name=${project_name}
supervisor_service="${app_name}"

deploy_user="vagrant"
deploy_server="devinggo.sre.im" # 请提前配置发布机到目标服务器之间的ssh key信任
deploy_dir=/opt/$project_name
supervisorctl="supervisorctl"

[[ "$env" = "test" ]] && {
    deploy_dir=${deploy_dir}_test
    supervisor_service="${app_name}_test"
}

cd "$script_dir" || exit 1

directory=$(awk -F= '/directory/{print $2}' ${supervisor_service}.ini)
$SED -i "s#$directory#$deploy_dir#" ${supervisor_service}.ini

# shellcheck disable=SC2029
ssh $deploy_user@$deploy_server "
    sudo mkdir -p $deploy_dir;
    sudo chown -R $deploy_user. $deploy_dir;
    rm -f $deploy_dir/$app_name
"

scp ${supervisor_service}.ini $deploy_user@$deploy_server:$deploy_dir

cd ../../../ || exit 1

make build

scp bin/linux_amd64/${app_name} $deploy_user@$deploy_server:$deploy_dir

# shellcheck disable=SC2029
ssh $deploy_user@$deploy_server "
    sudo \cp $deploy_dir/${supervisor_service}.ini /etc/supervisord.d/
    sudo $supervisorctl reload
"

exit 0
