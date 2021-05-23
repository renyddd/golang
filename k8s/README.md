# About Kubernetes Develop

## Some Reference
1. [极客时间 - 张磊 - 深入解析声明式API](https://time.geekbang.org/column/article/41876)
2. [kubebuilder tutorial](https://book.kubebuilder.io/introduction.html)
3. [Kubernetes Operator 快速入门教程](https://www.qikqiak.com/post/k8s-operator-101/)


## Some Step
1. [Run Dashboard with Docker Desktop](https://andrewlock.net/running-kubernetes-and-the-dashboard-with-docker-desktop/)

注意所使用的不通 yaml 文件，而生成的不通 namespace 等。
```bash
kubectl proxy
```
access at: http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/

