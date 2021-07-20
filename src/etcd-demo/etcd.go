package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

func EtcdDemo() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
	)

	// 创建配置对象，指定server地址并设置超时时间
	// 这里因为我用的是windows系统 docker安装在虚拟中
	// 所以地址填的是虚拟机ip

	config = clientv3.Config{
		Endpoints:   []string{"192.168.56.101:2379"},
		DialTimeout: 5 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {
		// 只是测试一下，有错误就直接panic吧
		panic(err)
	}

	_, err = client.Put(context.TODO(), "/user/Roki", "hello! etcd")
	if err != nil {
		panic(err)
	}

	response, err := client.Get(context.TODO(), "/user/Roki")
	if err != nil {
		panic(err)
	}

	for k, v := range response.Kvs {
		fmt.Println(k, v.Key)
	}
	fmt.Println(response)
}
