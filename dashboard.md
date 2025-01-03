


```
yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
yum install -y docker-ce docker-ce-cli containerd.io
systemctl start docker
docker pull apacherocketmq/rocketmq-dashboard:latest
docker run -d --name rocketmq-dashboard -e "JAVA_OPTS=-Drocketmq.namesrv.addr=129.226.169.221:9876" -p 9080:8080 -t apacherocketmq/rocketmq-dashboard:latest
```

![image](https://github.com/user-attachments/assets/5e7f8784-f054-4d6e-b7f1-43ea9f632d16)

