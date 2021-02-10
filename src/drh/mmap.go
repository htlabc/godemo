package drh

//只能在Linux上使用
//func MmapDemo(){
//	var err error
//	var fn string
//	var fd int
//	var fi os.FileInfo
//	var data []byte
//	//打开文件
//	fn = "/home/jues/go/src/GolangMmapDemo/1.data"
//	fd, err = syscall.Open(fn, syscall.O_RDWR, 0)
//	if nil != err {
//		fmt.Println("open file fail!!!")
//	}
//	//获取文件大小
//	fi, err = os.Stat(fn)
//	if nil != err {
//		fmt.Println("get file size fail!!!")
//	}
//	//映射到内存
//	data, err = syscall.Mmap(fd, 0, int(fi.Size()), syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC, syscall.MAP_SHARED)
//
//	//被 madvise 标记的这段地址空间，该进程仍然可以访问（不会segment fault），但是当读/写其中某一页时（例如malloc分配新的内存，或 Go 创建新的对象），内核会 重新分配 一个 用全0填充 的新页面。
//	//
//	//如果进程大量读写这段地址空间（即 release notes 说的 “a large fraction of the heap remains live”，堆空间大部分活跃），内核需要频繁分配页面、并且将页面内容清零，这会导致分配的延迟变高。
//	if err := syscall.Madvise(data, syscall.MADV_RANDOM); err != nil {
//		 fmt.Errorf("madvise: %s", err)
//	}
//
//
//	//
//	if nil != err {
//		fmt.Println("mmap fail!!!")
//		return
//	}
//	//
//	addr := &data[0]
//	fmt.Println("mmap success,addr=", addr, "size=", len(data))
//
//	//取消映射
//	syscall.Munmap(data)
//}
