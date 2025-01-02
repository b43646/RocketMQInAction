
Cluster deployment Architecture

| IP 地址 | NameServer 名称 | Broker 名称 |
|--|--|--|
| 129.226.179.102 | rocketmq-nameserver1 | rocketmq-master1 |
| 150.109.110.152 | rocketmq-nameserver2 | rocketmq-master2 |
| 150.109.105.75 | rocketmq-nameserver3 | rocketmq-master1-slave |
| 43.135.102.182 | rocketmq-nameserver4 | rocketmq-master2-slave |


Basic config for all servers
```
yum install -y java-11-openjdk java-11-openjdk-devel
java -version
systemctl stop firewalld

```

```

[root@rocketmq-nameserver1 ~]# ssh-keygen
[root@rocketmq-nameserver1 ~]# ssh-copy-id -i .ssh/id_ed25519.pub root@rocketmq-nameserver2
[root@rocketmq-nameserver1 ~]# ssh-copy-id -i .ssh/id_ed25519.pub root@rocketmq-nameserver3
[root@rocketmq-nameserver1 ~]# ssh-copy-id -i .ssh/id_ed25519.pub root@rocketmq-nameserver4

[root@rocketmq-nameserver1 ~]# cat /etc/hosts
129.226.179.102  rocketmq-nameserver1 rocketmq-master1
150.109.110.152  rocketmq-nameserver2 rocketmq-master2
150.109.105.75   rocketmq-nameserver3 rocketmq-master1-slave
43.135.102.182   rocketmq-nameserver4 rocketmq-master2-slave

[root@rocketmq-nameserver2 ~]# cat /etc/hosts
129.226.179.102  rocketmq-nameserver1 rocketmq-master1
150.109.110.152  rocketmq-nameserver2 rocketmq-master2
150.109.105.75   rocketmq-nameserver3 rocketmq-master1-slave
43.135.102.182   rocketmq-nameserver4 rocketmq-master2-slave

[root@rocketmq-nameserver3 ~]# cat /etc/hosts
129.226.179.102  rocketmq-nameserver1 rocketmq-master1
150.109.110.152  rocketmq-nameserver2 rocketmq-master2
150.109.105.75   rocketmq-nameserver3 rocketmq-master1-slave
43.135.102.182   rocketmq-nameserver4 rocketmq-master2-slave

[root@rocketmq-nameserver4 ~]# cat /etc/hosts
129.226.179.102  rocketmq-nameserver1 rocketmq-master1
150.109.110.152  rocketmq-nameserver2 rocketmq-master2
150.109.105.75   rocketmq-nameserver3 rocketmq-master1-slave
43.135.102.182   rocketmq-nameserver4 rocketmq-master2-slave
```


```
[root@rocketmq-nameserver1 ~]# wget https://dist.apache.org/repos/dist/release/rocketmq/5.3.1/rocketmq-all-5.3.1-bin-release.zip
[root@rocketmq-nameserver1 ~]# unzip rocketmq-all-5.3.1-bin-release.zip
[root@rocketmq-nameserver1 ~]# mv rocketmq-all-5.3.1-bin-release /usr/local/rocketmq

# Adjust memeory for JVM for testing, recommend better configuration for production environment
[root@rocketmq-nameserver1 ~]# vi /usr/local/rocketmq/bin/runbroker.sh

JAVA_OPT="${JAVA_OPT} -server -Xms2g -Xmx2g"

[root@rocketmq-nameserver1 ~]# vi /usr/local/rocketmq/bin/runserver.sh

JAVA_OPT="${JAVA_OPT} -server -Xms2g -Xmx2g -Xmn1g -XX:MetaspaceSize=128m -XX:MaxMetaspaceSize=320m"

# copy all files to other servers
[root@rocketmq-nameserver1 ~]# scp -r /usr/local/rocketmq/ root@rocketmq-nameserver4:/usr/local/;scp -r /usr/local/rocketmq/ root@rocketmq-nameserver3:/usr/local/;scp -r /usr/local/rocketmq/ root@rocketmq-nameserver2:/usr/local/
```

start nameserver
```
[root@rocketmq-nameserver1 ~]# cd /usr/local/rocketmq/bin/
# apply os tuning script
[root@rocketmq-nameserver1 ~]# sh os.sh
[root@rocketmq-nameserver1 ~]# nohup sh mqnamesrv &
[root@rocketmq-nameserver1 ~]# jps

[root@rocketmq-nameserver2 ~]# cd /usr/local/rocketmq/bin/
# apply os tuning script
[root@rocketmq-nameserver2 ~]# sh os.sh
[root@rocketmq-nameserver2 ~]# nohup sh mqnamesrv &
[root@rocketmq-nameserver2 ~]# jps

[root@rocketmq-nameserver3 ~]# cd /usr/local/rocketmq/bin/
# apply os tuning script
[root@rocketmq-nameserver3 ~]# sh os.sh
[root@rocketmq-nameserver3 ~]# nohup sh mqnamesrv &
[root@rocketmq-nameserver3 ~]# jps

[root@rocketmq-nameserver4 ~]# cd /usr/local/rocketmq/bin/
# apply os tuning script
[root@rocketmq-nameserver4 ~]# sh os.sh
[root@rocketmq-nameserver4 ~]# nohup sh mqnamesrv &
[root@rocketmq-nameserver4 ~]# jps
```

start broker+proxy

```
[root@rocketmq-nameserver1 ~]# cd /usr/local/rocketmq/bin/
[root@rocketmq-nameserver1 ~]# nohup sh mqbroker -n "rocketmq-nameserver1:9876;rocketmq-nameserver2:9876;rocketmq-nameserver3:9876;rocketmq-nameserver4:9876" -c ../conf/2m-2s-sync/broker-a.properties --enable-proxy &
[root@rocketmq-nameserver2 ~]# jps

[root@rocketmq-nameserver2 ~]# cd /usr/local/rocketmq/bin/
[root@rocketmq-nameserver2 ~]# nohup sh mqbroker -n "rocketmq-nameserver1:9876;rocketmq-nameserver2:9876;rocketmq-nameserver3:9876;rocketmq-nameserver4:9876" -c ../conf/2m-2s-sync/broker-b.properties --enable-proxy &
[root@rocketmq-nameserver2 ~]# jps

[root@rocketmq-nameserver3 ~]# cd /usr/local/rocketmq/bin/
[root@rocketmq-nameserver3 ~]# nohup sh mqbroker -n "rocketmq-nameserver1:9876;rocketmq-nameserver2:9876;rocketmq-nameserver3:9876;rocketmq-nameserver4:9876" -c ../conf/2m-2s-sync/broker-a-s.properties --enable-proxy &
[root@rocketmq-nameserver3 ~]# jps

[root@rocketmq-nameserver4 ~]# cd /usr/local/rocketmq/bin/
[root@rocketmq-nameserver4 ~]# nohup sh mqbroker -n "rocketmq-nameserver1:9876;rocketmq-nameserver2:9876;rocketmq-nameserver3:9876;rocketmq-nameserver4:9876" -c ../conf/2m-2s-sync/broker-b-s.properties --enable-proxy &
[root@rocketmq-nameserver4 ~]# jps
```

