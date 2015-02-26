# Retail Search Proof of Concept

## Prerequisites

***For local development***, the PoC requires that a virtual machine (VM) environment is installed. Specifically, install the following:

* [`Vagrant`](https://www.vagrantup.com/downloads.html)
* [`VirtualBox`](https://www.virtualbox.org/wiki/Downloads)

***For AWS deployment***, the Poc requires:

* `Terraform`, if you have Homebrew on your system, use `brew install terraform`
* `virtualenv`, use `pip install virtualenv`
* On the AWS console, create/download an EC2 key pair, and credentials for an AWS user
    * put the .pem file for the key pair in the `infra` folder
* `source go_init` to set environment variables needed for deployment (this step asks for the EC2 pair values and the name of the key you put in the infra folder)
* modify the `init/terraform/terraform.tfvars.sample` to contain the AWS key pair values from the credentials file. Save this file as `init/terraform/terraform.tfvars`. This step is necessary because Terraform doesn't yet allow access to the environment variables you set in the last step.
* create a symlink to, or make a copy of, `init/terraform/terraform.tfvars` in `init/terraform/<solution name>/`

In general, the project assumes a MacOSX environment. For development in a windows environment, some adjustments would need to be made.

## The ./go Script

The project provides a `./go` script to enables several convenience commands for deploying locally as well as to the cloud. Here's some background reading on why ./go scripts are [a good idea](http://www.thoughtworks.com/insights/blog/praise-go-script-part-i)

**IMPORTANT**
At the moment we have two separate ./go scripts, one for our solr tech stack and one for elasticsearch. To work on the solr setup use `./go_solr`. To work on ES use `./go_elasticsearch`. You can also set a `$TECH_STACK` environment variable and then call `./go` directly.

To see what commands are available, type the following from the root of the project:

```
./go_elasticsearch
```
(or `./go_solr`)

The following sections describe some of the commands available in greater detail. We'll use the elasticsearch tech stack in the examples, but you can also run all the same commands with `./go_solr` instead.

### ./go_elasticsearch local

Run `./go_elasticsearch local` to launch a local instance of Elasticsearch. This command will load a minimal VM using Vagrant and deploy Elasticsearch. A single Elasticsearch node can be accessed on the host machine on port 9200.

e.g.

```
curl 'http://localhost:9200/_cat?pretty'
```

### ./go aws_provision

Run `./go aws_provision` to stand up a disposable environment in AWS (using Terraform). This command stands up EC2 instances, but does not deploy Elasticsearch.

You will need valid AWS credentials for Terraform to standup EC2 instances. To tell Terraform about your AWS credentials you can set the `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables. Something like this:
	
```
export AWS_ACCESS_KEY_ID=<BLAH>
export AWS_SECRET_ACCESS_KEY=<BLAH>
```

### ./go aws_deploy
Run `./go aws_deploy` to deploy an elasticsearch cluster to the AWS environment (using ansible).

You'll need ansible installed (using `brew install ansible` for example).

### ./go aws_teardown
Run `./go aws_teardown` to nuke the AWS environment. Everything will go away, including any data or custom changes you've made.



## using the elasticsearch cluster

### Admin Dashboard URLs

For the Kopf dashboard: [host URL]/admin/

For the Marvel dashboard (Vagrant only): localhost:9200/_plugins/marvel/kibana/index.html

## Uploading Sample Data
You'll need to install Nodejs on your machine (brew install node or apt-get nodejs).

From the sampleData directory, run postData.js. postData.js takes two arguments, a file containing an array of JSON objects to bulk index, and the URL of your elasticsearch instance. 

example:
```
node postData.js sampleData.json http://localhost:9200/products/product/_bulk
```
