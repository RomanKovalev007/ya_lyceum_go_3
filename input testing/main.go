package inputT

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ParallelMapCtx(ctx context.Context, inputs []int, fn func(int) int, workers int) ([]int, error){
	n := len(inputs)

	wg := sync.WaitGroup{}
	res := make([]int, n)
	ch := make(chan int)

	go func ()  {
		defer close(ch)
		for i := 0; i < n; i ++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()
	
		

	for i := 0; i < workers; i ++ {
		wg.Add(1)
		go func (){
			defer wg.Done()
			for{
				select{
				case <- ctx.Done():
					return
				case idx, ok := <- ch:
					if !ok{
						return
					}
					res[idx] = fn(inputs[idx])
				}
			}
		}()
	}

	wg.Wait()

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func main(){
	inputs := []int{}
	for i := 0; i < 3; i ++{
		inputs = append(inputs, i)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Millisecond)
	defer cancel()

	res, err := ParallelMapCtx(ctx, inputs, func(i int) int { time.Sleep(1 * time.Millisecond); return i*2}, 1)
	if err != nil{
		fmt.Println(err.Error())
	}
	_ = res
	fmt.Println(res)
}