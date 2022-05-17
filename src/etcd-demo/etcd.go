package etcd

import (
	"bufio"
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"log"
	"os"
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




func EtcdElection(){
	cli,err:=clientv3.New(clientv3.Config{Endpoints: })
	if err!=nil{
		log.Fatal(err)
	}

	defer cli.Close()
	session,err:=concurrency.NewSession(cli)

	defer session.Close()

	var electName="test"
	e1:=concurrency.NewElection(session,electName)

	e1.Campaign(context.TODO(),"test")
	consolescanner:=bufio.NewScanner(os.Stdin)
	for consolescanner.Scan(){
		action:consolescanner.Text()
		switch action {
		case "elect":
			go elect()
		}
	}

}


var count int
// 选主
func elect(e1 *concurrency.Election, electName string) {
	log.Println("acampaigning for ID:", *nodeID)
	// 调用Campaign方法选主,主的值为value-<主节点ID>-<count>
	if err := e1.Campaign(context.Background(), fmt.Sprintf("value-%d-%d", *nodeID, count)); err != nil {
		log.Println(err)
	}
	log.Println("campaigned for ID:", *nodeID)
	count++
}

//new

// 为主设置新值
func proclaim(e1 *concurrency.Election, electName string) {
	log.Println("proclaiming for ID:", *nodeID)
	// 调用Proclaim方法设置新值,新值为value-<主节点ID>-<count>
	if err := e1.Proclaim(context.Background(), fmt.Sprintf("value-%d-%d", *nodeID, count)); err != nil {
		log.Println(err)
	}
	log.Println("proclaimed for ID:", *nodeID)
	count++
}
// 重新选主，有可能另外一个节点被选为了主
func resign(e1 *concurrency.Election, electName string) {
	log.Println("resigning for ID:", *nodeID)
	// 调用Resign重新选主
	if err := e1.Resign(context.TODO()); err != nil {
		log.Println(err)
	}
	log.Println("resigned for ID:", *nodeID)
}




//etcd 提供读写锁跟 mutex分布式锁


