package initialize

import (
	"github.com/alitto/pond"
	"runtime"
	"server/global"
)

func NewPool() *pond.WorkerPool {
	maxCap := global.ADTPL_CONFIG.WorkerPool.MaxCapicity
	numCpus := GetCpuNum()
	pool := pond.New(numCpus, maxCap)
	return pool
}

func GetCpuNum() int {
	numCpus := runtime.NumCPU()
	return numCpus
}
