package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type Model struct {
	name string
	kind reflect.Kind
	// reflectValue reflect.Value
}

func (m Model) InitModel(input any) Model {
	reflectInput := reflect.ValueOf(input)

	if(reflectInput.Kind().String() == "ptr"){
		reflectInput = reflectInput.Elem()
	}else {
		panic(errors.New("init model need a pointer input"))
	}

	if reflectInput.Kind().String() == "struct" {
		model := Model{
			name: strings.ToLower(reflectInput.Type().Name()),
			kind: reflectInput.Kind(),
		}
		return model
	} else {
		panic(errors.New("init model need a struct input"))
	}
}

func (m Model) Read(input any){
	reflectInput := reflect.ValueOf(input)
	if(reflectInput.Kind().String() == "ptr"){
		reflectInput = reflectInput.Elem()
	}
	if reflectInput.Kind().String() == "struct" {
		var path  string= "data/" + strings.ToLower(m.name)+".json"
		rawFile, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		// var data = []users.User{}
		var data []interface{}
		// reflect.ValueOf(input).Elem().Set(data)
		// make
		err = json.Unmarshal(rawFile, &data)

		// reflectInput.Elem().Set(reflect.ValueOf(data))
		reflectInput.Set(reflect.ValueOf(data[0]))
		// fmt.Println("input: ", input)

		if err != nil {
			panic(err)
		}
		fmt.Println("rawFile: ", rawFile)
		// fmt.Println(data)
	} else if reflectInput.Kind().String() == "slice"{
		var path  string= "data/" + strings.ToLower(m.name)+".json"
		rawFile, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		// var data = []users.User{}
		err = json.Unmarshal(rawFile, input)
		if err != nil {
			panic(err)
		}
	} else {
		panic(errors.New("Read receive struct"))
	}
}

// func Read(input any){
// 	reflectInput := reflect.ValueOf(input)
// 	if(reflectInput.Kind().String() == "ptr"){
// 		reflectInput = reflectInput.Elem()
// 	}
// 	fmt.Println(reflectInput, reflectInput.Type(), reflectInput.Kind())
// 	if reflectInput.Kind().String() == "struct" {
// 		var modelName string = reflectInput.Type().Name()
// 		var path  string= "data/" + strings.ToLower(modelName)+".json"
// 		rawFile, err := ioutil.ReadFile(path)
// 		if err != nil {
// 			panic(err)
// 		}
// 		// var data = []users.User{}
// 		var data = []users.User{}
// 		// reflect.ValueOf(input).Elem().Set(data)
// 		// make
// 		err = json.Unmarshal(rawFile, &data)
// 		// reflectInput.Elem().Set(reflect.ValueOf(data))
// 		reflectInput.Set(reflect.ValueOf(data[0]))
// 		fmt.Println("input: ", input)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println("rawFile: ", rawFile)
// 		// fmt.Println(data)
// 	} else if reflectInput.Kind().String() == "slice"{
// 		fmt.Println(reflectInput, reflectInput.Type(), reflectInput.Kind())
// 		var modelName string = reflectInput.Type().Elem().String()
// 		fmt.Println("modelname", modelName)
// 		var path  string= "data/" + strings.ToLower(modelName)+".json"
// 		rawFile, err := ioutil.ReadFile(path)
// 		if err != nil {
// 			panic(err)
// 		}
// 		// var data = []users.User{}
// 		err = json.Unmarshal(rawFile, input)
// 		if err != nil {
// 			panic(err)
// 		}
// 	} else {
// 		panic(errors.New("Read receive struct"))
// 	}
// }

func (m Model) Save(input any) {
	reflectInput := reflect.ValueOf(input)
	if reflectInput.Kind().String() == "slice" {
		file, _ := json.MarshalIndent(input, "", "  ")
		err := ioutil.WriteFile(strings.ToLower("data/"+m.name)+".json", file, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		panic(errors.New("save must slice"))
	}
}

// func checkId(val reflect.Value) bool {
// 	if(val.FieldByName("id") != reflect.Value{}){
// 		return false
// 	} else{
// 		return true
// 	}
// } 