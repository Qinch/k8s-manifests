# k8s-manifests


- Minikube工具:
  - Minikube is local Kubernetes, focusing on making it easy to learn and develop for Kubernetes.

- [上篇](https://zhuanlan.zhihu.com/p/1995144704958935811)相关yaml，见part1
	- some files from https://github.com/luksa/kubernetes-in-action
- [中篇](https://zhuanlan.zhihu.com/p/2045272589883507421)相关代码，见part2
	- app
		- 1 cd ./part2/app; make DOCKER_USER=qinchaowhut DOCKER_PWD=dockerhub密码 VERSION=v0.1.2
			- `注：`DOCKER_USER 和 DOCKER_PWD替换成你自己的Dockerhub账号/密码
		- 2 cd ./part2/app/deployments/manifests/ap-test; 
			- 2.1 kubectl create -f 1_namespace_test.yaml 
			- 2.2 kubectl create -f apps_v1_deployment_app.yaml
			- 2.3 kubectl create -f v1_service_app.yaml 
	- app-gateway
		- 1 cd ./part2/app-gateway; make DOCKER_USER=qinchaowhut DOCKER_PWD=dockerhuh密码 VERSION=v0.1.2
   		- 2 cd ./part2/app-gateway/deployments/manifests/ap-test;
   			- 2.1 kubectl create -f 1_namespace_test.yaml
			- 2.2 kubectl create -f apps_v1_deployment_app.yaml
   		- 2.3 kubectl create -f v1_service_app.yaml
	- 查看运行结果：
	<img width="963" height="198" alt="Clipboard_Screenshot_1780468556" src="https://github.com/user-attachments/assets/df9c84a6-6b0b-4e9a-becd-45304d5c174d" />
	<img width="699" height="102" alt="Clipboard_Screenshot_1780469766" src="https://github.com/user-attachments/assets/3bada8ac-e221-492c-853a-b2ff35d35691" />
	<img width="1186" height="114" alt="Clipboard_Screenshot_1780469454" src="https://github.com/user-attachments/assets/61a86828-cb5c-4286-9abd-833e5b89e61c" />

  	- TODO
 		- 代码中加入[viper](https://github.com/spf13/viper)读取配置文件。 

- 下篇
	- TODO
