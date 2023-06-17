package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/samply/golang-fhir-models/fhir-models/fhir"
)

// FHIRリソースのオブジェクトをJSONシリアライズして、FHIR文書を作成するサンプル
func main() {
	// Bundleの作成
	bundle := createBundle()
	compositionEntry := createCompositionEntry()

	// Compositionの作成
	bundle.Entry = append(bundle.Entry, compositionEntry)

	// JSONシリアライズ
	// v, err := bundle.MarshalJSON()
	// 改行、インデント等して出力
	v, err := json.MarshalIndent(bundle, "", "    ")

	if err != nil {
		log.Fatal(err) //終了
	}

	fmt.Println(string(v))

}

// Bundleリソースを作成
func createBundle() fhir.Bundle {
	bundle := fhir.Bundle{}
	now := time.Now().Format("2006-01-02T15:04:05Z07:00")
	bundle.Meta = &fhir.Meta{
		Profile:     []string{"http://jpfhir.jp/fhir/ePrescription/StructureDefinition/JP_Bundle_ePrescriptionData"},
		LastUpdated: &now,
	}
	return bundle
}

// CompositionリソースのEntryを作成
func createCompositionEntry() fhir.BundleEntry {
	composition := fhir.Composition{}
	//TODO: error
	r, _ := composition.MarshalJSON()

	compositionEntity := fhir.BundleEntry{
		Resource: r,
	}

	return compositionEntity
}
