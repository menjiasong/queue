package main

import (

	que "github.com/menjiasong00/queue"
)

func main() {

	que.NewConfig([]string{"10.10.18.130","5672","guest","guest"}).Push("TestJob","xxxxxx")

}