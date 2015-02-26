output "admin url" {
    value = "http://${aws_elb.admin-elb.dns_name}/admin"
}
