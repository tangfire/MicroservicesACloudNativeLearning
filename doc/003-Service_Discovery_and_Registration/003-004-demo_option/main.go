package main

import "fmt"

// option模式

const defaultValueC = 1

// ServiceConfig 我们定义一个服务配置结构体，里面有各自配置项
type ServiceConfig struct {
	A string
	B string
	C int

	X struct{}
	Y Info
}

type Info struct {
	addr string
}

// NewServiceConfig 创建一个ServiceConfig的函数
func NewServiceConfig1(a, b string, c int) *ServiceConfig {
	return &ServiceConfig{
		A: a,
		B: b,
		C: c,
	}
}

// 要求A和B必须传，C爱传不传，不传就用默认值
func NewServiceConfig2(a, b string, c ...int) *ServiceConfig {
	valueC := defaultValueC
	if len(c) > 0 {
		valueC = c[0]
	}
	return &ServiceConfig{
		A: a,
		B: b,
		C: valueC,
	}
}

// option模式

type FuncServiceConfigOption func(*ServiceConfig)

func NewServiceConfig3(a, b string, opts ...FuncServiceConfigOption) *ServiceConfig {
	sc := &ServiceConfig{
		A: a,
		B: b,
		C: defaultValueC,
	}
	// 针对可能传进来的FuncServiceConfigOption参数做处理
	for _, opt := range opts {
		opt(sc)
	}

	return sc
}

func WithC(c int) FuncServiceConfigOption {
	return func(sc *ServiceConfig) {
		sc.C = c
	}
}

func WithY(info Info) FuncServiceConfigOption {
	return func(sc *ServiceConfig) {
		sc.Y = info
	}
}

func main() {
	//f := WithC(10)
	sc := NewServiceConfig3("tangfire", "xierqi", WithC(10), WithY(Info{addr: "127.0.0.1:8080"}))
	fmt.Printf("sc:%#v\n", sc)

	sc.C = 100
	fmt.Printf("sc:%#v\n", sc) // 可以直接改？？

	// 进阶版Option使用
	cfg := NewConfig(18)
	fmt.Printf("cfg:%#v\n", cfg)

	cfg2 := NewConfig(18, WithConfigName("张三"))
	fmt.Printf("cfg:%#v\n", cfg2)

	//cfg2.age = ??? 在其他包中没有办法修改

}

const defaultValueName = "fireshine"

// 进阶版Option
type config struct {
	name string
	age  int
}

func NewConfig(age int, opts ...ConfigOption) *config {
	cfg := &config{
		age:  age,
		name: defaultValueName,
	}

	for _, opt := range opts {
		opt.apply(cfg)
	}

	return cfg
}

type ConfigOption interface {
	apply(*config)
}

type funcOption struct {
	f func(*config)
}

func (f funcOption) apply(cfg *config) {
	f.f(cfg)
}

func NewFuncOption(f func(*config)) *funcOption {
	return &funcOption{f: f}
}

func WithConfigName(name string) ConfigOption {
	return NewFuncOption(func(cfg *config) {
		cfg.name = name
	})
}
