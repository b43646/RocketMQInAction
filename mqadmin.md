
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


### 2. NameSrv

```
# Get NameSrv Config
[root@VM-0-3-centos bin]# ./mqadmin getNamesrvConfig -n rocketmq-nameserver1:9876
============rocketmq-nameserver1:9876============
defaultThreadPoolNums                             =  16
...

# Update NameSrv Config
[root@VM-0-3-centos bin]# ./mqadmin updateNamesrvConfig -n rocketmq-nameserver1:9876 -k defaultThreadPoolNums -v 32
update name server config success![rocketmq-nameserver1:9876]
defaultThreadPoolNums : 32
[root@VM-0-3-centos bin]# ./mqadmin getNamesrvConfig -n rocketmq-nameserver1:9876
============rocketmq-nameserver1:9876============
defaultThreadPoolNums                             =  32
...
[root@VM-0-3-centos bin]# ./updateNamesrvConfig -n rocketmq-nameserver1:9876 -k defaultThreadPoolNums -v 16
update name server config success![rocketmq-nameserver1:9876]
defaultThreadPoolNums : 16
[root@VM-0-3-centos bin]# ./mqadmin getNamesrvConfig -n rocketmq-nameserver1:9876
============rocketmq-nameserver1:9876============
defaultThreadPoolNums                             =  16
...
```

### 3. broker

```
# show broker status
[root@VM-0-3-centos bin]# ./mqadmin brokerStatus -n rocketmq-nameserver1:9876 -b rocketmq-master1:10911

# Get broker config
[root@VM-0-3-centos bin]# ./mqadmin getBrokerConfig -n rocketmq-nameserver1:9876 -b rocketmq-master1:10911

# update broker config
[root@VM-0-3-centos bin]# ./mqadmin updateBrokerConfig -n rocketmq-nameserver1:9876 -b rocketmq-master1:10911 maxMsgsNumBatch 128

```


### 4. Message

```
[root@VM-0-3-centos bin]# ./mqadmin updatetopic -n rocketmq-nameserver1:9876 -c DefaultCluster -t test -r 4 -w 4
create topic to 172.19.0.3:10911 success.
create topic to 172.19.0.12:10911 success.
TopicConfig [topicName=test, readQueueNums=4, writeQueueNums=4, perm=RW-, topicFilterType=SINGLE_TAG, topicSysFlag=0, order=false, attributes={}]

# send 100 messages
#!/bin/bash

for ((i=1; i<=100; i++))
do
 ./mqadmin sendMessage -n rocketmq-nameserver1:9876 -c DefaultCluster -t test -p "hello world $i"
  sleep 0.5
done

# Consume message
[root@VM-0-3-centos bin]# ./mqadmin consumeMessage -n rocketmq-nameserver1:9876  -t test
MessageQueue [topic=test, brokerName=broker-a, queueId=3] print msg finished. status=NO_NEW_MSG, offset=0
MessageQueue [topic=test, brokerName=broker-b, queueId=3] print msg finished. status=NO_NEW_MSG, offset=0
Consume ok
MSGID: AC130003B6CD5505405709099C730000 MessageExt [brokerName=broker-a, queueId=2, storeSize=245, queueOffset=0, sysFlag=0, bornTimestamp=1735812424821, bornHost=/172.19.0.3:37984, storeTimestamp=1735812424848, storeHost=/172.19.0.3:10911, msgId=AC13000300002A9F0000000000000000, commitLogOffset=0, bodyCRC=459924658, reconsumeTimes=0, preparedTransactionOffset=0, toString()=Message{topic='test', flag=0, properties={MSG_REGION=DefaultRegion, UNIQ_KEY=AC130003B6CD5505405709099C730000, CLUSTER=DefaultCluster, MIN_OFFSET=0, TAGS=DefaultCluster, WAIT=true, TRACE_ON=true, MAX_OFFSET=1}, body=[104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 32, 50, 48, 50, 53], transactionId='null'}] BODY: hello world 2025
MessageQueue [topic=test, brokerName=broker-a, queueId=2] print msg finished. status=NO_NEW_MSG, offset=1
MessageQueue [topic=test, brokerName=broker-b, queueId=2] print msg finished. status=NO_NEW_MSG, offset=0
MessageQueue [topic=test, brokerName=broker-a, queueId=1] print msg finished. status=NO_NEW_MSG, offset=0
MessageQueue [topic=test, brokerName=broker-b, queueId=1] print msg finished. status=NO_NEW_MSG, offset=0
MessageQueue [topic=test, brokerName=broker-a, queueId=0] print msg finished. status=NO_NEW_MSG, offset=0
MessageQueue [topic=test, brokerName=broker-b, queueId=0] print msg finished. status=NO_NEW_MSG, offset=0


# Query Message by Message ID

[root@VM-0-3-centos bin]# ./mqadmin queryMsgById  -n rocketmq-nameserver1:9876 -t test -i AC130003BD9855054057090DBED90000
OffsetID:            AC13000C00002A9F00000000000001EA
Topic:               test
Tags:                [DefaultCluster]
Keys:                [null]
Queue ID:            1
Queue Offset:        0
CommitLog Offset:    490
Reconsume Times:     0
Born Timestamp:      2025-01-02 18:11:35,770
Store Timestamp:     2025-01-02 18:11:35,783
Born Host:           172.19.0.3:50514
Store Host:          172.19.0.12:10911
System Flag:         0
Properties:          {MSG_REGION=DefaultRegion, UNIQ_KEY=AC130003BD9855054057090DBED90000, CLUSTER=DefaultCluster, TAGS=DefaultCluster, WAIT=true, TRACE_ON=true}
Message Body Path:   /tmp/rocketmq/msgbodys/AC130003BD9855054057090DBED90000


# Query Message by offset
[root@VM-0-3-centos bin]# ./mqadmin queryMsgByOffset -b broker-a -n rocketmq-nameserver1:9876 -t test -i 0 -o 1
OffsetID:            AC13000300002A9F00000000000005AF
Topic:               test
Tags:                [DefaultCluster]
Keys:                [null]
Queue ID:            0
Queue Offset:        1
CommitLog Offset:    1455
Reconsume Times:     0
Born Timestamp:      2025-01-02 18:19:05,031
Store Timestamp:     2025-01-02 18:19:05,050
Born Host:           172.19.0.3:57012
Store Host:          172.19.0.3:10911
System Flag:         0
Properties:          {MSG_REGION=DefaultRegion, UNIQ_KEY=AC130003CA2955054057091499C70000, CLUSTER=DefaultCluster, MIN_OFFSET=0, TAGS=DefaultCluster, WAIT=true, TRACE_ON=true, MAX_OFFSET=2}
Message Body Path:   /tmp/rocketmq/msgbodys/AC130003CA2955054057091499C70000


WARN: No Consumer

```

### 5. CLuster

```
[root@VM-0-3-centos bin]# ./mqadmin clusterList -n rocketmq-nameserver1:9876
#Cluster Name           #Broker Name            #BID  #Addr                  #Version              #InTPS(LOAD)     #OutTPS(LOAD)  #Timer(Progress)        #PCWait(ms)  #Hour         #SPACE    #ACTIVATED
DefaultCluster          broker-a                0     172.19.0.3:10911       V5_3_1                 0.10(0,0ms)       0.00(0,0ms)  0-0(0.0w, 0.0, 0.0)               0  0.27          0.1700          true
DefaultCluster          broker-a                1     172.19.0.7:10911       V5_3_1                 0.20(0,0ms)       0.00(0,0ms)  4-0(0.0w, 0.0, 0.0)               0  0.27          0.1600         false
DefaultCluster          broker-b                0     172.19.0.12:10911      V5_3_1                 0.00(0,0ms)       0.00(0,0ms)  0-0(0.0w, 0.0, 0.0)               0  0.20          0.1600          true
DefaultCluster          broker-b                1     172.19.0.10:10911      V5_3_1                 0.00(0,0ms)       0.00(0,0ms)  2-0(0.0w, 0.0, 0.0)               0  0.20          0.1600         false


[root@VM-0-3-centos bin]# ./mqadmin statsAll -n rocketmq-nameserver1:9876
#Topic                                                            #Consumer Group                                                  #Accumulation      #InTPS     #OutTPS   #InMsg24Hour  #OutMsg24Hour
rmq_sys_SYNC_BROKER_MEMBER_VM-0-3-centos                                                                                                      0        0.00                          0    NO_CONSUMER
RMQ_SYS_TRANS_HALF_TOPIC                                                                                                                      0        0.00                          0    NO_CONSUMER
VM-0-3-centos                                                                                                                                 0        0.00                          0    NO_CONSUMER
test                                                                                                                                          0        0.20                         45    NO_CONSUMER
BenchmarkTest                                                                                                                                 0        0.00                          0    NO_CONSUMER
OFFSET_MOVED_EVENT                                                                                                                            0        0.00                          0    NO_CONSUMER
TBW102                                                                                                                                        0        0.00                          0    NO_CONSUMER
rmq_sys_REVIVE_LOG_DefaultCluster                                                                                                             0        0.00                          0    NO_CONSUMER
SELF_TEST_TOPIC                                                                                                                               0        0.00                          0    NO_CONSUMER
DefaultCluster                                                                                                                                0        0.00                          0    NO_CONSUMER
SCHEDULE_TOPIC_XXXX                                                                                                                           0        0.00                          0    NO_CONSUMER
DefaultCluster_REPLY_TOPIC                                                                                                                    0        0.00                          0    NO_CONSUMER
rmq_sys_wheel_timer                                                                                                                           0        0.00                          0    NO_CONSUMER
rmq_sys_SYNC_BROKER_MEMBER_broker-b                                                                                                           0        0.00                          0    NO_CONSUMER
rmq_sys_SYNC_BROKER_MEMBER_broker-a                                                                                                           0        0.00                          0    NO_CONSUMER
RMQ_SYS_TRANS_OP_HALF_TOPIC                                                                                                                   0        0.00                          0    NO_CONSUMER
broker-b                                                                                                                                      0        0.00                          0    NO_CONSUMER
broker-a                                                                                                                                      0        0.00                          0    NO_CONSUMER

```

