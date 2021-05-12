package pool

var (
	slb SmoothWeightedLoadBalancer
)

func DefaultLoadBalancer() LoadBalancer{
	return &slb
}

// LoadBalancer 负载均衡器
type LoadBalancer interface {
	GetPool(pools []NetPool)NetPool
}

// SmoothWeightedLoadBalancer 平滑加权轮询算法
type SmoothWeightedLoadBalancer struct {
	
}

func (lb *SmoothWeightedLoadBalancer) GetPool(pools []NetPool)NetPool {
	if len(pools)==0{
		return nil
	}
	// 默认采用平滑加权轮询算法
	for _, pool := range pools {
		pool.SetCurrentWeights(pool.InitWeights()+pool.CurrentWeights())
	}
	var maxValue = -1
	var pickPool NetPool
	var sum = 0
	for _, pool := range pools {
		if maxValue==-1||pool.CurrentWeights()>maxValue {
			maxValue = pool.CurrentWeights()
			pickPool = pool
		}
		sum += pool.CurrentWeights()
	}
	if pickPool == nil {
		return nil
	}
	pickPool.SetCurrentWeights(pickPool.CurrentWeights()-sum)
	return pickPool
}