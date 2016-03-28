package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/BurntSushi/toml"
	"github.com/jgcarvalho/zeca-search/search"
)

func run(fnconfig string) {
	var conf search.Config
	md, err := toml.DecodeFile(fnconfig, &conf)
	if err != nil {
		log.Fatal(err)
	}
	if len(md.Undecoded()) > 0 {
		fmt.Printf("Chaves desconhecidas no arquivo de configuração: %q\n", md.Undecoded())
		fmt.Printf("Chaves conhecidas: %q\n", md.Keys())
		fmt.Println("Configuration:", conf)
		return
	}
	fmt.Println("Configuration:", conf)
	search.RunMaster(conf)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// flags
	fnconfig := flag.String("config", "default", "Configuration file")
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to file")
	memprofile := flag.String("memprofile", "", "write memory profile to this file")
	flag.Parse()

	// cpuprofile
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// memprofile
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		defer f.Close()
	}

	if *fnconfig == "default" {
		run(os.Getenv("GOPATH") + "/src/github.com/jgcarvalho/zeca-search-master/config.toml")
	} else {
		run(*fnconfig)
	}

}
