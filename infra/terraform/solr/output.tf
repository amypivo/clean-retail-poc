output "admin url" {
    value = "http://${aws_instance.solr-instance.public_dns}:8080/solr/admin/"
}
