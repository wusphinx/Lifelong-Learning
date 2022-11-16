# Jenkins

## 常用
### 容器镜像标签
使用Jenkins制作容器镜像时，经常需要打标签，镜像通常跟源码是有关联的，为了便于管理，可将当前编译镜像的仓库commit id作为镜像的tag
```
IMG_TAG=$(git rev-parse --short ${GIT_COMMIT})
```