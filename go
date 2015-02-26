#!/bin/bash
set -e -u

echo 
echo tech stack: ${TECH_STACK:?"must supply a TECH_STACK env var"}
echo

function check_for_aws_creds { 
  VARIABLES_SET=1
  if [ -z ${AWS_ACCESS_KEY_ID:-""} ]; then
    VARIABLES_SET=0
    echo "An AWS_ACCESS_KEY_ID must be set to a non-null value in the environment"
  fi
  if [ -z ${AWS_SECRET_ACCESS_KEY:-""} ]; then
    VARIABLES_SET=0
    echo "An AWS_SECRET_ACCESS_KEY must be set to a non-null value in the environment"
  fi
  if [ ${VARIABLES_SET} == 0 ]; then
    echo
    echo "To set variables, source the go_init file"
    exit
  fi
}

function using_isolated_ansible {
  infra/local_scripts/ensure_ansible.sh
  set +u
  source infra/managed_tools/python_env/bin/activate
  set -u
}

function run_terraform {
  cd infra/terraform/${TECH_STACK}
  shift
  terraform $*
}

function run_ansible {
  using_isolated_ansible

  cd infra/ansible 
  export ANSIBLE_HOST_KEY_CHECKING=False # with disposable/phoenix servers this is not really valuable
  ansible-playbook -i ec2.py -u ubuntu --private-key=../${AWS_KEY_NAME:-"rei_search"}.pem ec2_${TECH_STACK}_playbook.yml
}

function provision_vagrant {
  using_isolated_ansible

  vagrant up --no-provision $TECH_STACK && vagrant provision $TECH_STACK
}

case "${1:-}" in 

'')
  echo -e "valid commands are:\n\tlocal\n\taws_provision\n\taws_deploy\n\taws_teardown"
  ;;
aws_provision)
  run_terraform elasticsearch apply
  ;;
aws_teardown)
  run_terraform elasticsearch destroy
  ;;
terraform)
  shift
  run_terraform elasticsearch $*
  ;;
aws_deploy)
  check_for_aws_creds
  run_ansible
  ;;
local_es)
  echo '`local_es` is deprecated, please use `local`'
  ;;
local)
  provision_vagrant
  ;;
*)
  echo 'unrecognized command'
  ;;
esac
