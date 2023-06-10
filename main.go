package main

import (
	"io/ioutil"
	"log"

	"github.com/samply/golang-fhir-models/fhir-models/fhir"
	"github.com/xeipuuv/gojsonschema"
)

// あえて分かりやすくするため１つのメソッドに手続き的に書いてあるので、本当に実装したい場合は保守性の高いモジュール化されたコード書くこと
// 参考
// https://github.com/samply/golang-fhir-models
func main() {
	// 診療情報提供書のHL7 FHIRのサンプルデータBundle-BundleReferralExample01.jsonを読み込み
	fileData, err := ioutil.ReadFile("Bundle-BundleReferralExample01.json")
	if err != nil {
		log.Fatal(err) //終了
	}
	// HL7 FHIRのJSONスキーマfhir.schema.json
	schemaData, err := ioutil.ReadFile("fhir.schema.json")
	if err != nil {
		log.Fatal(err) //終了
	}

	// サンプルデータのJSONスキーマ検証
	schemaLoader := gojsonschema.NewBytesLoader(schemaData)
	// 診療情報提供書のHL7 FHIRのサンプルデータBundle-BundleReferralExample01.json
	docummentLoader := gojsonschema.NewBytesLoader(fileData)
	result, err := gojsonschema.Validate(schemaLoader, docummentLoader)
	if err != nil {
		log.Fatal(err)
	}
	if result.Valid() {
		log.Println("JSON Schema Check: ドキュメントは有効です")
	} else {
		//検証エラー
		log.Println("JSON Schema Check: ドキュメントに不備があります")
		for _, desc := range result.Errors() {
			log.Printf("- %s\n", desc)
		}
	}

	// BundleリソースのJSONパース
	bundle, err := fhir.UnmarshalBundle(fileData)
	if err != nil {
		log.Fatal(err) //終了
	}
	// Bundleリソース解析
	log.Printf("Bundle type: %s", bundle.Type.Display())
	entries := bundle.Entry
	var subjectRefId string
	for i, v := range entries {
		//TODO: ResourceリソースとしてJSONパースしてもResourceTypeを取得できない
		//resource, err := fhir.UnmarshalResource(v.Resource)
		//if err != nil {
		//    log.Fatal(err) //終了
		//}
		//log.Printf("Resource Id: %s", *resource.Id)

		// 最初のEntryであるCompositionリソースの解析する例
		if i == 0 {
			composition, err := fhir.UnmarshalComposition(v.Resource)
			if err != nil {
				log.Fatal(err) //終了
			}
			title := composition.Title
			log.Printf("文書名: %s", title)
			subjectDisplay := composition.Subject.Display
			subjectRefId = *composition.Subject.Reference
			subjectType := composition.Subject.Type
			log.Printf("subject display: %s", *subjectDisplay)
			log.Printf("subject reference Id: %s", subjectRefId)
			log.Printf("subject reference type: %s", *subjectType)
			continue
		}
		switch *v.FullUrl {
		case subjectRefId:
			// Compostion.subjectが参照するPatientの解析する例
			patient, err := fhir.UnmarshalPatient(v.Resource)
			if err != nil {
				log.Fatal(err) //終了
			}
			// 患者番号の取得
			log.Printf("患者番号: %s", *patient.Identifier[0].Value)
			// 患者氏名の取得
			humanNames := patient.Name
			for _, humanName := range humanNames {
				valueCode := humanName.Extension[0].ValueCode
				if *valueCode == "IDE" {
					log.Printf("患者氏名: %s", *humanName.Text)
				} else {
					log.Printf("患者カナ氏名: %s", *humanName.Text)
				}
			}
		}

	}
}
