package MYSQL

import (
	"sync"
)

// SyMutex 增加锁  互斥锁
var SyMutex sync.Mutex

// SyMutexsr 增加锁 读写锁
var SyMutexsr sync.RWMutex

type UserMain struct {
	Account   string
	Password  string
	Authority int
}
