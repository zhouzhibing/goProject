package config

import (
	"strings"
	"os"
	"bufio"
	"io"
)

//const middle = "========="
const middle = "="

type Config struct {
	Mymap  map[string]string
	strcet string
}

func NewConfig(path string) * Config{
	this := new(Config)
	this.initConfig(path)
	return this
}

func (c *Config) initConfig(path string) {
	c.Mymap = make(map[string]string)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := strings.TrimSpace(string(b))
		//fmt.Println(s)
		if strings.Index(s, "#") == 0 {
			continue
		}

		//n1 := strings.Index(s, "[")
		//n2 := strings.LastIndex(s, "]")
		//if n1 > -1 && n2 > -1 && n2 > n1+1 {
		//	c.strcet = strings.TrimSpace(s[n1+1 : n2])
		//	continue
		//}
		//
		//if len(c.strcet) == 0 {
		//	continue
		//}
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		frist := strings.TrimSpace(s[:index])
		if len(frist) == 0 {
			continue
		}
		second := strings.TrimSpace(s[index+1:])

		//pos := strings.Index(second, "\t#")
		//if pos > -1 {
		//	second = second[0:pos]
		//}
		//
		//pos = strings.Index(second, " #")
		//if pos > -1 {
		//	second = second[0:pos]
		//}
		//
		//pos = strings.Index(second, "\t//")
		//if pos > -1 {
		//	second = second[0:pos]
		//}
		//
		//pos = strings.Index(second, " //")
		//if pos > -1 {
		//	second = second[0:pos]
		//}
		//
		//if len(second) == 0 {
		//	continue
		//}
		//
		//key := c.strcet + middle + frist
		key := frist
		c.Mymap[key] = strings.TrimSpace(second)
	}
}
//
//func (c Config) Read(node, key string) string {
//	key = node + middle + key
//	v, found := c.Mymap[key]
//	if !found {
//		return ""
//	}
//	return v
//}

func (c Config) GetValue(key string) string {
	v, found := c.Mymap[key]
	if !found {
		return ""
	}
	return v
}
