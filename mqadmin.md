
### 1. Topic

```
# Create a topic

[root@VM-0-3-centos bin]# ./mqadmin updatetopic -n rocketmq-nameserver1:9876 -c DefaultCluster -t newTopic -r 8 -w 8
create topic to 172.19.0.3:10911 success.
create topic to 172.19.0.12:10911 success.
TopicConfig [topicName=newTopic, readQueueNums=8, writeQueueNums=8, perm=RW-, topicFilterType=SINGLE_TAG, topicSysFlag=0, order=false, attributes={}]

# list all topics in the cluster

[root@VM-0-3-centos bin]# ./mqadmin topiclist -n rocketmq-nameserver1:9876
rmq_sys_SYNC_BROKER_MEMBER_VM-0-3-centos
RMQ_SYS_TRANS_HALF_TOPIC
VM-0-3-centos
BenchmarkTest
OFFSET_MOVED_EVENT
TBW102
rmq_sys_REVIVE_LOG_DefaultCluster
SELF_TEST_TOPIC
DefaultCluster
SCHEDULE_TOPIC_XXXX
DefaultCluster_REPLY_TOPIC
newTopic
rmq_sys_wheel_timer
rmq_sys_SYNC_BROKER_MEMBER_broker-b
rmq_sys_SYNC_BROKER_MEMBER_broker-a
RMQ_SYS_TRANS_OP_HALF_TOPIC
broker-b
broker-a

# show topic status

[root@VM-0-3-centos bin]# ./mqadmin topicstatus -n rocketmq-nameserver1:9876 -t newTopic
#Broker Name                      #QID  #Min Offset           #Max Offset             #Last Updated
broker-a                          0     0                     0
broker-a                          1     0                     0
broker-a                          2     0                     0
broker-a                          3     0                     0
broker-a                          4     0                     0
broker-a                          5     0                     0
broker-a                          6     0                     0
broker-a                          7     0                     0
broker-b                          0     0                     0
broker-b                          1     0                     0
broker-b                          2     0                     0
broker-b                          3     0                     0
broker-b                          4     0                     0
broker-b                          5     0                     0
broker-b                          6     0                     0
broker-b                          7     0                     0

# delete a topic
[root@VM-0-3-centos bin]# ./mqadmin deletetopic  -n rocketmq-nameserver1:9876 -c DefaultCluster -t newTopic
delete topic [newTopic] from cluster [DefaultCluster] success.
delete topic [newTopic] from NameServer success.


```
