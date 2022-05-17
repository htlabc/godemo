package singleflight

//这节课，我来给你介绍两个非常重要的扩展并发原语：SingleFlight 和 CyclicBarrier。SingleFlight 的作用是将并发请求合并成一个请求，以减少对下层服务的压力；而 CyclicBarrier 是一个可重用的栅栏并发原语，用来控制一组请求同时执行的数据结构。
