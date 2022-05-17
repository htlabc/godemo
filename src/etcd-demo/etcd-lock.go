package etcd

import (
	"bufio"
	"flag"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	recipe "htl.com/pkg/mod/github.com/coreos/etcd/contrib/recipes"
	"math/rand"

	"log"
	"os"
	"strings"
	"sync"
)

var ( addr = flag.String("addr", "http://127.0.0.1:2379", "etcd addresses")
queueName = flag.String("name", "my-test-queue", "queue name"))


//分布式队列
func etcdMultiQueue(){

	flag.Parse()


	endpoints:=strings.Split(*addr,",")
	cli,err:=clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err!=nil{
		log.Fatal(err)
	}
	defer cli.Close()


	q:=recipe.NewQueue(cli,*queueName)

	consolescanner:=bufio.NewScanner(os.Stdin)
	consolescanner.Text()

	for consolescanner.Scan(){
		action:=consolescanner.Text()
		items:=strings.Split(action," ")
		switch items[0] {
		case "push":
			if len(items)!=2{
				fmt.Println("must set valueto push")
				continue
			}
			q.Enqueue(items[1])
		case "pop"://从队列弹出
		v,err:=q.Dequeue()
		if err!=nil{
			fmt.Println()
		}
		fmt.Println(v)

		case "quit","exit":
			return
		default:
			fmt.Println("unknow")
		}
	}




}


func BarrierTest(){


	package main


	import (
		"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"


	"github.com/coreos/etcd/clientv3"
	recipe "github.com/coreos/etcd/contrib/recipes"
	)


	var (
		addr        = flag.String("addr", "http://127.0.0.1:2379", "etcd addresses")
		barrierName = flag.String("name", "my-test-queue", "barrier name")
	)



		flag.Parse()


		// 解析etcd地址
		endpoints := strings.Split(*addr, ",")


		// 创建etcd的client
		cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
		if err != nil {
			log.Fatal(err)
		}
		defer cli.Close()


		// 创建/获取栅栏
		b := recipe.NewBarrier(cli, *barrierName)


		// 从命令行读取命令
		consolescanner := bufio.NewScanner(os.Stdin)
		for consolescanner.Scan() {
			action := consolescanner.Text()
			items := strings.Split(action, " ")
			switch items[0] {
			case "hold": // 持有这个barrier
				b.Hold()
				fmt.Println("hold")
			case "release": // 释放这个barrier
				b.Release()
				fmt.Println("released")
			case "wait": // 等待barrier被释放
				b.Wait()
				fmt.Println("after wait")
			case "quit", "exit": //退出
				return
			default:
				fmt.Println("unknown action")
			}
		}

}


//DoubleBarrier：计数型栅栏etcd 还提供了另外一种栅栏，叫做 DoubleBarrier，这也是一种非常有用的栅栏。这个栅栏初始化的时候需要提供一个计数 count，如下所示：

package main


import (
"bufio"
"flag"
"fmt"
"log"
"os"
"strings"


"github.com/coreos/etcd/clientv3"
"github.com/coreos/etcd/clientv3/concurrency"
recipe "github.com/coreos/etcd/contrib/recipes"
)


var (
	addr        = flag.String("addr", "http://127.0.0.1:2379", "etcd addresses")
	barrierName = flag.String("name", "my-test-doublebarrier", "barrier name")
	count       = flag.Int("c", 2, "")
)


func main() {
	flag.Parse()


	// 解析etcd地址
	endpoints := strings.Split(*addr, ",")


	// 创建etcd的client
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	// 创建session
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()


	// 创建/获取栅栏
	b := recipe.NewDoubleBarrier(s1, *barrierName, *count)


	// 从命令行读取命令
	consolescanner := bufio.NewScanner(os.Stdin)
	for consolescanner.Scan() {
		action := consolescanner.Text()
		items := strings.Split(action, " ")
		switch items[0] {
		case "enter": // 持有这个barrier
			b.Enter()
			fmt.Println("enter")
		case "leave": // 释放这个barrier
			b.Leave()
			fmt.Println("leave")
		case "quit", "exit": //退出
			return
		default:
			fmt.Println("unknown action")
		}
	}
}



//etcd stm实现

//
//Txn().If(cond1, cond2, ...).Then(op1, op2, ...,).Else(op1’, op2’, …)




package main


import (
"context"
"flag"
"fmt"
"log"
"math/rand"
"strings"
"sync"


"github.com/coreos/etcd/clientv3"
"github.com/coreos/etcd/clientv3/concurrency"
)


var (
	addr = flag.String("addr", "http://127.0.0.1:2379", "etcd addresses")
)


func main() {
	flag.Parse()


	// 解析etcd地址
	endpoints := strings.Split(*addr, ",")


	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()


	// 设置5个账户，每个账号都有100元，总共500元
	totalAccounts := 5
	for i := 0; i < totalAccounts; i++ {
		k := fmt.Sprintf("accts/%d", i)
		if _, err = cli.Put(context.TODO(), k, "100"); err != nil {
			log.Fatal(err)
		}
	}


	// STM的应用函数，主要的事务逻辑
	exchange := func(stm concurrency.STM) error {
		// 随机得到两个转账账号
		from, to := rand.Intn(totalAccounts), rand.Intn(totalAccounts)
		if from == to {
			// 自己不和自己转账
			return nil
		}
		// 读取账号的值
		fromK, toK := fmt.Sprintf("accts/%d", from), fmt.Sprintf("accts/%d", to)
		fromV, toV := stm.Get(fromK), stm.Get(toK)
		fromInt, toInt := 0, 0
		fmt.Sscanf(fromV, "%d", &fromInt)
		fmt.Sscanf(toV, "%d", &toInt)


		// 把源账号一半的钱转账给目标账号
		xfer := fromInt / 2
		fromInt, toInt = fromInt-xfer, toInt+xfer


		// 把转账后的值写回
		stm.Put(fromK, fmt.Sprintf("%d", fromInt))
		stm.Put(toK, fmt.Sprintf("%d", toInt))
		return nil
	}


	// 启动10个goroutine进行转账操作
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				if _, serr := concurrency.NewSTM(cli, exchange); serr != nil {
					log.Fatal(serr)
				}
			}
		}()
	}
	wg.Wait()


	// 检查账号最后的数目
	sum := 0
	accts, err := cli.Get(context.TODO(), "accts/", clientv3.WithPrefix()) // 得到所有账号
	if err != nil {
		log.Fatal(err)
	}
	for _, kv := range accts.Kvs { // 遍历账号的值
		v := 0
		fmt.Sscanf(string(kv.Value), "%d", &v)
		sum += v
		log.Printf("account %s: %d", kv.Key, v)
	}


	log.Println("account sum is", sum) // 总数
}
