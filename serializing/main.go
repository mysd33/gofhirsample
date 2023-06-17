package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/samply/golang-fhir-models/fhir-models/fhir"
)

var now = time.Now().Format("2006-01-02T15:04:05Z07:00")

// FHIRリソースのオブジェクトをJSONシリアライズして、FHIR文書を作成するサンプル
func main() {
	// Patientの作成
	patientEntry, err := createPatienEntry()
	if err != nil {
		log.Fatal(err) //終了
	}

	// Compositionの作成
	compositionEntry, err := createCompositionEntry(patientEntry)
	if err != nil {
		log.Fatal(err) //終了
	}

	// Bundleの作成
	bundle := createBundle()
	bundle.Entry = append(bundle.Entry,
		*compositionEntry,
		*patientEntry)

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
func createBundle() *fhir.Bundle {
	bundle := fhir.Bundle{}

	// meta
	bundle.Meta = &fhir.Meta{
		Profile:     []string{"http://jpfhir.jp/fhir/ePrescription/StructureDefinition/JP_Bundle_ePrescriptionData"},
		LastUpdated: &now,
	}
	return &bundle
}

// CompositionリソースのEntryを作成
func createCompositionEntry(patientEntry *fhir.BundleEntry) (*fhir.BundleEntry, error) {
	composition := fhir.Composition{}
	// meta
	composition.Meta = &fhir.Meta{
		Profile:     []string{"http://jpfhir.jp/fhir/eReferral/StructureDefinition/JP_Composition_ePrescriptionData"},
		LastUpdated: &now,
	}
	// text
	composition.Text = &fhir.Narrative{
		Status: fhir.NarrativeStatusGenerated,
		//TODO:　Divの「<」、「>」をjson.MarshalJSON時に自動エスケープしてしまうの修正が必要
		Div: "<div xmlns=\"http://www.w3.org/1999/xhtml\">xxx</div>",
	}
	// extention
	versionNumber := "1.0"
	composition.Extension = append(composition.Extension, fhir.Extension{
		Url:         "http://hl7.org/fhir/StructureDefinition/composition-clinicaldocument-versionNumber",
		ValueString: &versionNumber,
	})

	// identifier
	identifierSystem := "http://jpfhir.jp/fhir/Common/IdSystem/resourceInstance-identifier"
	identifierValue := "1311234567-2020-00123456" // 処方箋番号入れる
	composition.Identifier = &fhir.Identifier{
		System: &identifierSystem,
		Value:  &identifierValue,
	}

	// status
	composition.Status = fhir.CompositionStatusFinal

	// type
	typeCode := "57833-6"
	typeSystem := "http://jpfhir.jp/fhir/Common/CodeSystem/doc-typecodes"
	typeDisplay := "処方箋"
	composition.Type = fhir.CodeableConcept{
		Coding: []fhir.Coding{
			{
				Code:    &typeCode,
				System:  &typeSystem,
				Display: &typeDisplay,
			},
		},
	}

	// category
	categoryCode := "01"
	categorySystem := "http://jpfhir.jp/fhir/ePrescription/CodeSystem/prescription-category"
	categoryDisplay := "処方箋"
	composition.Category = []fhir.CodeableConcept{
		{
			Coding: []fhir.Coding{
				{
					Code:    &categoryCode,
					System:  &categorySystem,
					Display: &categoryDisplay,
				},
			},
		},
	}

	// subject
	composition.Subject = &fhir.Reference{
		Reference: patientEntry.FullUrl,
	}

	// TODO: 項目追加
	// encounter

	// date
	composition.Date = now

	// TODO: 項目追加
	// author
	composition.Author = []fhir.Reference{}

	// title
	composition.Title = "処方箋"

	// TODO: 項目追加
	// custodian
	// event
	// section

	r, err := composition.MarshalJSON()

	if err != nil {
		return nil, err
	}

	fullUrl := fmt.Sprintf("urn:uuid:%s", uuid.NewString())
	compositionEntity := fhir.BundleEntry{
		FullUrl:  &fullUrl,
		Resource: r,
	}

	return &compositionEntity, nil
}

// PatientリソースのEntryを作成する
func createPatienEntry() (*fhir.BundleEntry, error) {
	patient := fhir.Patient{}

	// meta
	patient.Meta = &fhir.Meta{
		Profile:     []string{"http://jpfhir.jp/fhir/ePrescription/StructureDefinition/JP_Patient_ePrescriptionData"},
		LastUpdated: &now,
	}

	// text
	patient.Text = &fhir.Narrative{
		Status: fhir.NarrativeStatusGenerated,
		//TODO:　Divの「<」、「>」をjson.MarshalJSON時に自動エスケープしてしまうの修正が必要
		Div: "<div xmlns=\"http://www.w3.org/1999/xhtml\">xxx</div>",
	}

	// identifier
	identifierSystem := "urn:oid:1.2.392.100495.20.3.51.11311234567"
	identifierValue := "00000010"
	hokenjaNo := "00012345"
	hihokennshaKigo := "あいう"
	hihokennshaBango := "１８７"
	hihokennshaEdaNo := "05"
	insuranceSystem := fmt.Sprintf("http:/jpfhir.jp/fhir/ccs/Idsysmem/JP_Insurance_member/%s", hokenjaNo)
	insuranceValue := fmt.Sprintf("%s:%s:%s:%s", hokenjaNo, hihokennshaKigo, hihokennshaBango, hihokennshaEdaNo)
	patient.Identifier = []fhir.Identifier{
		// 患者番号
		{
			System: &identifierSystem,
			Value:  &identifierValue,
		},
		// 保険個人識別子
		{
			System: &insuranceSystem,
			Value:  &insuranceValue,
		},
	}

	// name
	useOfficial := fhir.NameUseOfficial
	kanjiText := "東京　太郎"
	kanjiFamily := "東京"
	kanjiGiven := "太郎"
	ide := "IDE"
	readingText := "トウキョウ　タロウ"
	readingFamily := "トウキョウ"
	readingGiven := "タロウ"
	syl := "SYL"
	patient.Name = []fhir.HumanName{
		// 漢字
		{
			Use:    &useOfficial,
			Text:   &kanjiText,
			Family: &kanjiFamily,
			Given:  []string{kanjiGiven},
			Extension: []fhir.Extension{
				{
					Url:         "http:// hl7.org/fhir/StructureDefinition/iso21090-EN-representation",
					ValueString: &ide,
				},
			},
		},
		// よみ
		{
			Use:    &useOfficial,
			Text:   &readingText,
			Family: &readingFamily,
			Given:  []string{readingGiven},
			Extension: []fhir.Extension{
				{
					Url:         "http:// hl7.org/fhir/StructureDefinition/iso21090-EN-representation",
					ValueString: &syl,
				},
			},
		},
	}

	// gender
	male := fhir.AdministrativeGenderMale
	patient.Gender = &male

	// birthdate
	patient.BirthDate = &now

	// address
	addressText := "神奈川県横浜市港区１－２－３"
	postalCode := "123-4567"
	country := "JP"
	patient.Address = []fhir.Address{
		{
			Text:       &addressText,
			PostalCode: &postalCode,
			Country:    &country,
		},
	}

	r, err := patient.MarshalJSON()
	if err != nil {
		return nil, err
	}

	fullUrl := fmt.Sprintf("urn:uuid:%s", uuid.NewString())
	patientEntity := fhir.BundleEntry{
		FullUrl:  &fullUrl,
		Resource: r,
	}

	return &patientEntity, nil
}
