package main

import (
	"fmt"
	"github.com/sevico/pingpongTest/model"
)

type BenchService struct {
	Size       int
	ReminTimes int
}

func (s *BenchService) Start(arg *model.BenchArgs, reply *int) error {
	s.Size = arg.Size
	s.ReminTimes = arg.Times
	*reply = 0
	fmt.Printf("Size:%d ReminTimes:%d\n", arg.Size, arg.Times)
	return nil
}
