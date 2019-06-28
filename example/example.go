package main


import (
	"context"
	"flag"
	"fmt"

	"github.com/cocktail18/paladin"

	"github.com/BurntSushi/toml"
)

type exampleConf struct {
	Bool    bool
	Int     int64
	Float   float64
	String  string
	Strings []string
}

func (e *exampleConf) Set(text string) error {
	var ec exampleConf
	if err := toml.Unmarshal([]byte(text), &ec); err != nil {
		return err
	}
	*e = ec
	fmt.Println("set val", ec)
	return nil
}

func ExampleClient() {
	if err := paladin.Init(); err != nil {
		panic(err)
	}
	var ec exampleConf
	// var setter
	if err := paladin.Watch("example.toml", &ec); err != nil {
		panic(err)
	}
	if err := paladin.Get("example.toml").UnmarshalTOML(&ec); err != nil {
		panic(err)
	}
	// use exampleConf
	// watch event key
	go func() {
		for event := range paladin.WatchEvent(context.TODO(), "key") {
			fmt.Println(event)
		}
	}()

}

func main(){
	flag.Parse()
	ExampleClient()
}
