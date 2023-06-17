package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/samply/golang-fhir-models/fhir-models/fhir"
)

// FHIRリソースのオブジェクトをJSONシリアライズして、FHIR文書を作成するサンプル
func main() {
	// Bundleの作成
	bundle := fhir.Bundle{}
	//v, err := bundle.MarshalJSON()

	//改行、インデント等して出力
	v, err := json.MarshalIndent(bundle, "", "    ")

	if err != nil {
		log.Fatal(err) //終了
	}

	fmt.Println(string(v))

}
