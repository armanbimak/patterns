package main

import (
	"fmt"
	"sync"
)

func loadConfig(path string) {
	fmt.Println(fmt.Sprintf("loaded config from %s", path))
}

func main() {
	singletonGuarantee := sync.Once{}

	singletonGuarantee.Do(func() {
		loadConfig("/some/path")
	})

	singletonGuarantee.Do(func() {
		fmt.Println("this will not work since Once object flag already set 1")
	})

}
