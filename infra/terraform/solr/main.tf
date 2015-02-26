provider "aws" {
    access_key = "${var.aws_access_key}"
    secret_key = "${var.aws_secret_access_key}"
    region = "us-east-1"
}

resource "aws_instance" "solr-instance" {
    ami = "ami-98aa1cf0"
    instance_type = "t1.micro"
		key_name = "rei_search"
    tags { 
				Role = "solr-instance"
    }
  
    security_groups = ["${aws_security_group.allow_ssh.name}","${aws_security_group.allow_tomcat_http.name}"]
}

resource "aws_security_group" "allow_tomcat_http" {
  name = "allow_tomcat_http"
  description = "Allow Tomcat HTTP on 8080"

  ingress {
    from_port = 8080
    to_port = 8080
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "allow_ssh" {
  name = "allow_ssh"
  description = "Allow SSH"

  ingress {
      from_port = 22
      to_port = 22
      protocol = "tcp"
      cidr_blocks = ["0.0.0.0/0"]
  }
}
