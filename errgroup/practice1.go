package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	//group := errgroup.Group{}
	withContext, _ := errgroup.WithContext(context.Background())
	withContext.Go(func() error {
		return errors.New("liulin")
	})
	withContext.Go(func() error {
		fmt.Println("liulin")
		return nil
	})
	withContext.Go(func() error {
		fmt.Println("cxq")
		return nil
	})
	err := withContext.Wait()
	if err != nil {
		fmt.Println(err.Error())
	}

}
