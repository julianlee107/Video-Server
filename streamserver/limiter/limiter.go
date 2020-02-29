package limiter

import "log"

type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the limitation")
		return false
	}
	cl.bucket <- 1
	log.Printf("Successfully got connection")
	return true
}

func (cl *ConnLimiter) ReleaseConn() {

	c := <-cl.bucket
	log.Printf("Release connection %d\n", c)

}
