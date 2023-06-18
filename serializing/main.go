package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/samply/golang-fhir-models/fhir-models/fhir"
)

// FHIR形式に変換対象の入力データを疑似した構造体
type PrescriptionData struct {
	PrescriptionNo   string //処方箋番号
	PatientNo        string //患者番号
	HokenjaNo        string // 保険者番号
	HihokennshaKigo  string // 被保険者記号
	HihokennshaBango string // 被保険者番号
	HihokennshaEdaNo string // 被保険者枝番
	KanjiLastName    string // 漢字姓
	KanjiFirstName   string // 漢字名
	KanaLastName     string // かな姓
	KanaFirstName    string // かな名
	GenderCode       int    // 性別コード(0:男性、1:女性、2:その他、3:不明）
	Birthday         string // 誕生日
	Zip              string // 郵便番号
	Address          string // 住所
	LastUpdated      string // 最終更新日
}

// FHIRリソースのオブジェクトをJSONシリアライズして、FHIR文書を作成するサンプル
func main() {
	// FHIR形式に変換対象の入力データ
	input := PrescriptionData{
		PrescriptionNo:   "1311234567-2020-00123456",
		PatientNo:        "00000010",
		HokenjaNo:        "00012345",
		HihokennshaKigo:  "あいう",
		HihokennshaBango: "１８７",
		HihokennshaEdaNo: "05",
		KanjiLastName:    "東京",
		KanjiFirstName:   "太郎",
		KanaLastName:     "トウキョウ",
		KanaFirstName:    "タロウ",
		GenderCode:       0,
		Birthday:         "1920-02-11",
		Zip:              "123-4567",
		Address:          "神奈川県横浜市港区１－２－３",
		LastUpdated:      time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}

	// Patientの作成
	patientEntry, err := createPatienEntry(input)
	if err != nil {
		log.Fatal(err) //終了
	}

	// Compositionの作成
	compositionEntry, err := createCompositionEntry(input, patientEntry)
	if err != nil {
		log.Fatal(err) //終了
	}

	// Bundleの作成
	bundle := createBundle(input)
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
func createBundle(prescriptionData PrescriptionData) *fhir.Bundle {
	bundle := fhir.Bundle{}

	// meta
	bundle.Meta = &fhir.Meta{
		Profile:     []string{"http://jpfhir.jp/fhir/ePrescription/StructureDefinition/JP_Bundle_ePrescriptionData"},
		LastUpdated: &prescriptionData.LastUpdated,
	}
	return &bundle
}

// CompositionリソースのEntryを作成
func createCompositionEntry(prescriptionData PrescriptionData, patientEntry *fhir.BundleEntry) (*fhir.BundleEntry, error) {
	composition := fhir.Composition{}
	// meta
	composition.Meta = &fhir.Meta{
		Profile:     []string{"http://jpfhir.jp/fhir/eReferral/StructureDefinition/JP_Composition_ePrescriptionData"},
		LastUpdated: &prescriptionData.LastUpdated,
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
	identifierValue := prescriptionData.PrescriptionNo
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
	composition.Date = prescriptionData.LastUpdated

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
func createPatienEntry(prescriptionData PrescriptionData) (*fhir.BundleEntry, error) {
	patient := fhir.Patient{}

	// meta
	patient.Meta = &fhir.Meta{
		Profile:     []string{"http://jpfhir.jp/fhir/ePrescription/StructureDefinition/JP_Patient_ePrescriptionData"},
		LastUpdated: &prescriptionData.LastUpdated,
	}

	// text
	patient.Text = &fhir.Narrative{
		Status: fhir.NarrativeStatusGenerated,
		//TODO:　Divの「<」、「>」をjson.MarshalJSON時に自動エスケープしてしまうの修正が必要
		Div: "<div xmlns=\"http://www.w3.org/1999/xhtml\">xxx</div>",
	}

	// identifier
	identifierSystem := "urn:oid:1.2.392.100495.20.3.51.11311234567"
	insuranceSystem := fmt.Sprintf("http://jpfhir.jp/fhir/ccs/Idsysmem/JP_Insurance_member/%s", prescriptionData.HokenjaNo)
	insuranceValue := fmt.Sprintf("%s:%s:%s:%s", prescriptionData.HokenjaNo, prescriptionData.HihokennshaKigo,
		prescriptionData.HihokennshaBango, prescriptionData.HihokennshaEdaNo)
	patient.Identifier = []fhir.Identifier{
		// 患者番号
		{
			System: &identifierSystem,
			Value:  &prescriptionData.PatientNo,
		},
		// 保険個人識別子
		{
			System: &insuranceSystem,
			Value:  &insuranceValue,
		},
	}

	// name
	useOfficial := fhir.NameUseOfficial
	kanjiText := fmt.Sprintf("%s　%s", prescriptionData.KanjiLastName, prescriptionData.KanjiFirstName)
	ide := "IDE"
	readingText := fmt.Sprintf("%s　%s", prescriptionData.KanaLastName, prescriptionData.KanaFirstName)
	syl := "SYL"
	patient.Name = []fhir.HumanName{
		// 漢字
		{
			Use:    &useOfficial,
			Text:   &kanjiText,
			Family: &prescriptionData.KanjiLastName,
			Given:  []string{prescriptionData.KanjiFirstName},
			Extension: []fhir.Extension{
				{
					Url:         "http://hl7.org/fhir/StructureDefinition/iso21090-EN-representation",
					ValueString: &ide,
				},
			},
		},
		// よみ
		{
			Use:    &useOfficial,
			Text:   &readingText,
			Family: &prescriptionData.KanaLastName,
			Given:  []string{prescriptionData.KanaFirstName},
			Extension: []fhir.Extension{
				{
					Url:         "http://hl7.org/fhir/StructureDefinition/iso21090-EN-representation",
					ValueString: &syl,
				},
			},
		},
	}

	// gender
	var gender fhir.AdministrativeGender
	switch prescriptionData.GenderCode {
	case 0:
		gender = fhir.AdministrativeGenderMale
	case 1:
		gender = fhir.AdministrativeGenderFemale
	case 2:
		gender = fhir.AdministrativeGenderOther
	default:
		gender = fhir.AdministrativeGenderUnknown
	}
	patient.Gender = &gender

	// birthdate
	patient.BirthDate = &prescriptionData.Birthday

	// address
	country := "JP"
	patient.Address = []fhir.Address{
		{
			Text:       &prescriptionData.Address,
			PostalCode: &prescriptionData.Zip,
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
