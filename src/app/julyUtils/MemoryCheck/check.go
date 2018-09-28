package MemoryCheck

//func startCPUProfile() {
//
//	if *cpuProfile != "" {
//		f, err := os.Create(*cpuProfile)
//		if err != nil {
//			fmt.Fprintf(os.Stderr, "Can not create cpu profile output file: %s",
//				err)
//			return
//		}
//		if err := pprof.StartCPUProfile(f); err != nil {
//			fmt.Fprintf(os.Stderr, "Can not start cpu profile: %s", err)
//			f.Close()
//			return
//		}
//	}
//}
//
//func stopCPUProfile() {
//	if *cpuProfile != "" {
//		pprof.StopCPUProfile() // 把记录的概要信息写到已指定的文件
//	}
//}