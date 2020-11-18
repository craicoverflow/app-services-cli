---
id: rhoas_kafka_topics_create
title: rhoas kafka topics create
---
## rhoas kafka topics create

Create topic

### Synopsis

Create topic in the current selected Managed Kafka cluster

```
rhoas kafka topics create [flags]
```

### Options

```
  -f, --config-file string   A path to a file containing extra configuration variables. If this option is not supplied, default configurations will be used
  -h, --help                 help for create
  -n, --name string          Topic name (required)
  -p, --partitions int32     Set number of partitions (default 3)
  -r, --replicas int32       Set number of replicas (default 2)
```

### SEE ALSO

* [rhoas kafka topics](rhoas_kafka_topics.md)	 - Manage topics

###### Auto generated by spf13/cobra on 18-Nov-2020