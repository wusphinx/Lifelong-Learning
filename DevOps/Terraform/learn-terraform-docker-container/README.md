# Terraform
## docker
```
$ terraform init 
$ terraform fmt
$ terraform validate
$ terraform apply
$ terraform state list
$ terraform destroy
```
第一感觉，用docker-compose能搞定，使用apply更新以及destroy删除容器后，意识到terraform是在定标准，统一可以对docker这类基础设施提供的操作，由provider这类插件接入就好了，整体配置不多，可读性也不错，虽然第一次，但感觉还挺好用的。