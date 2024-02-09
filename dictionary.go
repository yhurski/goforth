package main

const DictNameMaxLength = 255

var DictContext *gf_dict

type gf_dict struct {
	name string
	link *gf_dict
	code uint32
}

func InitDictionary() {

}

func CreateDictionaryRecord(name string, code uint32, flag byte) {

}
