package main

import "fmt"

// ref: https://lailin.xyz/post/factory.html

type jsonRuleConfigParser struct{}
type yamlRuleConfigParser struct{}

// IRuleConfigParser 工厂用于产生 Parse 的实例
type IRuleConfigParser interface {
	Parse(data []byte)
}

func (j jsonRuleConfigParser) Parse(data []byte) {
	fmt.Println("json parsing ...", data)
}

func (y yamlRuleConfigParser) Parse(data []byte) {
	fmt.Println("yaml parsing ...", data)
}

type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

// 工厂方法规定：每个产品都有一个专属工厂
//这其实就像是将，每一个 NewThing 的具体类型函数从大的工厂类里面抽离出来，最外层工厂为工厂的工厂

type jsonRuleConfigParseFactory struct{}

func (j jsonRuleConfigParseFactory) CreateParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

type yamlRuleConfigParseFactory struct{}

func (y yamlRuleConfigParseFactory) CreateParser() IRuleConfigParser {
	return yamlRuleConfigParser{}
}

func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return jsonRuleConfigParseFactory{}
	case "yaml":
		return yamlRuleConfigParseFactory{}
	}
	return nil
}
