package main

import "testing"

func TestFactory(t *testing.T) {
	jsonFactory := NewIRuleConfigParserFactory("json")
	dojsonParse := jsonFactory.CreateParser().Parse
	dojsonParse([]byte("hello"))

	doyamlParse := NewIRuleConfigParserFactory("yaml").CreateParser().Parse
	doyamlParse([]byte(", world"))
}