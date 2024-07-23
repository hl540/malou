package worker

import (
	"github.com/hl540/malou/utils"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type WorkerPool struct {
	pool map[string]string
	lock sync.Mutex
}

func (wp *WorkerPool) ResetSize(size int) {
	wp.lock.Lock()
	defer wp.lock.Unlock()

	length := len(wp.pool)

	// 如果预期的大小和当前一致，不做处理
	if length == size {
		return
	}

	// 如果预期大小比当前大，需要扩容
	if length < size {
		for i := 0; i <= size-length; i++ {
			workID := utils.StringWithCharset(10, utils.Charset2)
			wp.pool[workID] = ""
		}
		return
	}

	// 如果预期大小比当前小，需要缩减
	// 正在运行的work不会被缩减，所有有时候并不会达到预期大小，但是这个操作回持续进行
	var keysToDelete []string
	for k := range wp.pool {
		if len(keysToDelete) == length-size {
			break
		}
		keysToDelete = append(keysToDelete, k)
	}
	for _, key := range keysToDelete {
		delete(wp.pool, key)
	}
}

// TryWorker 尝试获取work，占用这个work
func (wp *WorkerPool) TryWorker() string {
	wp.lock.Lock()
	defer wp.lock.Unlock()
	wp.lock.TryLock()

	for workID, available := range wp.pool {
		if available == "" {
			wp.pool[workID] = time.Now().String()
			return workID
		}
	}
	return ""
}

// Worker 使用work，拉取到pipeline后使用正式的pipelineID填充
func (wp *WorkerPool) Worker(workID, pipelineID string) bool {
	wp.lock.Lock()
	defer wp.lock.Unlock()
	wp.lock.TryLock()

	// 检查work是否存在
	if _, ok := wp.pool[workID]; !ok {
		return false
	}
	wp.pool[workID] = pipelineID
	return true
}

// Release 归还令牌
func (wp *WorkerPool) Release(workID string) {
	wp.lock.Lock()
	defer wp.lock.Unlock()

	// 检查work是否存在
	if _, ok := wp.pool[workID]; ok {
		wp.pool[workID] = ""
	}
}

// Status 获取WorkerPool状态
func (wp *WorkerPool) Status() map[string]string {
	// 返回副本
	status := make(map[string]string)
	for k, v := range wp.pool {
		status[k] = v
	}
	return status
}

// WithDone 等待全部worker执行完成
func (wp *WorkerPool) WithDone(timeout int64) {
	withTimeout := time.NewTimer(time.Duration(timeout) * time.Second)
	for {
		select {
		case <-withTimeout.C:
			return
		default:
			time.Sleep(time.Second)
			size := len(wp.pool)
			free := wp.FreeNumber()
			logrus.New().Infof("running workers: %d", size-free)
			if size == free {
				return
			}
		}
	}
}

// FreeNumber 空闲的worker数量
func (wp *WorkerPool) FreeNumber() int {
	wp.lock.Lock()
	defer wp.lock.Unlock()
	number := 0
	for _, w := range wp.pool {
		if w == "" {
			number++
		}
	}
	return number
}

var Pool *WorkerPool

func InitWorkerPool(poolSize int) {
	Pool = &WorkerPool{pool: make(map[string]string)}
	for i := 0; i < poolSize; i++ {
		workID := utils.StringWithCharset(10, utils.Charset2)
		Pool.pool[workID] = ""
	}
}
