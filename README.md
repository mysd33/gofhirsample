# Go言語での FHIRのサンプル

- [診療情報提供書HL7FHIR記述仕様](https://std.jpfhir.jp/)に基づくサンプルデータ（Bundle-BundleReferralExample01.json）に対して、FHIRプロファイルでの検証、パースするサンプルプログラムです。
- また、FHIRリソース(Bundle)として作成したオブジェクトを、FHIRのJSON文字列で出力（シリアライズ）するサンプルプログラムもあります。
## プロファイルの検証（バリデーション）とパース
- FHIRプロファイルでの検証
    - Goの場合、[HL7のConfluenceのページ「Open Source Implementations」](https://confluence.hl7.org/display/FHIR/Open+Source+Implementations)で紹介されている、FHIRのリファレンス実装がありません。
    - また、その他、検索しても、Goでは、JavaのHAPI等と違い、FHIRの構造定義ファイルでの検証を行うライブラリがなさそうです。
    - ですが、[HL7 FHIR v4.0.1:R4のダウンロードページ](https://hl7.org/fhir/R4/downloads.html)に[JSON Schema形式のファイル](https://hl7.org/fhir/R4/fhir.schema.json.zip)が提供されています。        
    - そこで、JSONスキーマによる検証ができる[gojsonschema](https://github.com/xeipuuv/gojsonschema)というライブラリを使って、JSONスキーマの検証をしています。
- 【未実施】JPCoreプロファイル、文書情報プロファイルでの検証
    - [JPCoreプロファイル](https://jpfhir.jp/fhir/core/)のサイトにJPCoreプロファイルの構造定義ファイルがあります。
    - また、[https://std.jpfhir.jp/](https://std.jpfhir.jp/)のサイトに、JPCoreを含むスナップショット形式の[診療情報提供書の文書情報プロファイル（IGpackage2023.4.27 snapshot形式: jp-ereferral-0.9.6-snap.tgz）](https://jpfhir.jp/fhir/eReferral/jp-ereferral-0.9.7-snap.tgz)があります。 
    - いずれも、[FHIR package仕様](https://registry.fhir.org/learn)に従ったnpmパッケージ形式です。    
    - ですが、FHIRプロファイルのようなJSONスキーマ形式では提供されていないようで、Goの場合、JPCoreプロファイル、文書情報プロファイルレベルの検証を実施するのが難しそうです。

- FHIRデータのパース
    - 前述の通り、Goの場合、[HL7のConfluenceのページ「Open Source Implementations」](https://confluence.hl7.org/display/FHIR/Open+Source+Implementations)で紹介されている、FHIRのリファレンス実装がありません。
    - ですが、検索すると、[Golang FHIR Models](https://github.com/samply/golang-fhir-models)というライブラリが出てきたので、それを使って、パースをしてみました。
        - JavaのHAPI FHIRのような実装と比べると、バリデータ機能もないですし、コミュニティとしての信頼性も低いと考えています。
    - 上記は、内部では、[Go標準のJSONライブラリ(encoding/json)](https://pkg.go.dev/encoding/json)を使用します。

## シリアライズ
- FHIRデータからJSONへのシリアライズ
    - パース同様、[Golang FHIR Models](https://github.com/samply/golang-fhir-models)および、[Go標準のJSONライブラリ(encoding/json)](https://pkg.go.dev/encoding/json)を使って、シリアライズします。


## 実行方法
- 検証・パースするサンプルAPの使い方
    - ビルド後、生成されたexeファイルを実行してください。
```sh
# parsingフォルダへ移動
cd parsing
# ビルド
go build
# 実行
parsing-example.exe
```
- JSONシリアライズするサンプルAPの使い方
    - ビルド後、生成されたexeファイルを実行してください。
```sh
# serializingフォルダへ移動
cd serializing
# ビルド
go build
# 実行
serializing-example.exe
```

## 検証・パースの実行結果の例

```sh
>example.exe
# JSONスキーマチェック結果
2023/06/10 23:16:52 JSON Schema Check: ドキュメントは有効です
# テストデータのパース結果
2023/06/10 23:16:52 Bundle type: Document
2023/06/10 23:16:52 文書名: 診療情報提供書
2023/06/10 23:16:52 subject display: 患者リソースPatient
2023/06/10 23:16:52 subject reference Id: urn:uuid:0a48a4bf-0d87-4efb-aafd-d45e0842a4dd
2023/06/10 23:16:52 subject reference type: Patient
2023/06/10 23:16:52 患者番号: 12345
2023/06/10 23:16:52 患者氏名: 田中 太郎
2023/06/10 23:16:52 患者カナ氏名: タナカ タロウ
```

## JSONシリアライズ実行結果の例
- [処方情報のFHIR記述仕様書](https://jpfhir.jp/fhir/ePrescriptionData/igv1/)に従い、JSON文字列のほんの一部分が生成出来てるのが分かります。

```sh
D:\git\gofhirsample\serializing>serializing-example.exe
{
    "meta": {
        "lastUpdated": "2023-06-17T22:59:14+09:00",
        "profile": [
            "http://jpfhir.jp/fhir/ePrescription/StructureDefinition/JP_Bundle_ePrescriptionData"
        ]
    },
    "type": "document",
    "entry": [
        {
            "fullUrl": "urn:uuid:dcb3b52d-717b-4a1e-af7c-6dc86fb32d08",
            "resource": {
                "meta": {
                    "lastUpdated": "2023-06-17T22:59:14+09:00",
                    "profile": [
                        "http://jpfhir.jp/fhir/eReferral/StructureDefinition/JP_Composition_ePrescriptionData"      
                    ]
                },
                "text": {
                    "status": "generated",
                    "div": "\u003cdiv xmlns=\"http://www.w3.org/1999/xhtml\"\u003exxx\u003c/div\u003e"
                },
                "extension": [
                    {
                        "url": "http://hl7.org/fhir/StructureDefinition/composition-clinicaldocument-versionNumber",
                        "valueString": "1.0"
                    }
                ],
                "identifier": {
                    "system": "http://jpfhir.jp/fhir/Common/IdSystem/resourceInstance-identifier",
                    "value": "1311234567-2020-00123456"
                },
                "status": "final",
                "type": {
                    "coding": [
                        {
                            "system": "http://jpfhir.jp/fhir/Common/CodeSystem/doc-typecodes",
                            "code": "57833-6",
                            "display": "処方箋"
                        }
                    ]
                },
                "category": [
                    {
                        "coding": [
                            {
                                "system": "http://jpfhir.jp/fhir/ePrescription/CodeSystem/prescription-category",
                                "code": "01",
                                "display": "処方箋"
                            }
                        ]
                    }
                ],
                "subject": {
                    "reference": "urn:uuid:10ffd50b-ae08-45d2-b496-8a80e69539f8"
                },
                "date": "2023-06-17T22:59:14+09:00",
                "author": [],
                "title": "処方箋",
                "resourceType": "Composition"
            }
        },
        {
            "fullUrl": "urn:uuid:10ffd50b-ae08-45d2-b496-8a80e69539f8",
            "resource": {
                "meta": {
                    "lastUpdated": "2023-06-17T22:59:14+09:00",
                    "profile": [
                        "http://jpfhir.jp/fhir/ePrescription/StructureDefinition/JP_Patient_ePrescriptionData"
                    ]
                },
                "text": {
                    "status": "generated",
                    "div": "\u003cdiv xmlns=\"http://www.w3.org/1999/xhtml\"\u003exxx\u003c/div\u003e"
                },
                "identifier": [
                    {
                        "system": "urn:oid:1.2.392.100495.20.3.51.11311234567",
                        "value": "00000010"
                    },
                    {
                        "system": "http:/jpfhir.jp/fhir/ccs/Idsysmem/JP_Insurance_member/00012345",
                        "value": "00012345:あいう:１８７:05"
                    }
                ],
                "name": [
                    {
                        "extension": [
                            {
                                "url": "http:// hl7.org/fhir/StructureDefinition/iso21090-EN-representation",
                                "valueString": "IDE"
                            }
                        ],
                        "use": "official",
                        "text": "東京　太郎",
                        "family": "東京",
                        "given": [
                            "太郎"
                        ]
                    },
                    {
                        "extension": [
                            {
                                "url": "http:// hl7.org/fhir/StructureDefinition/iso21090-EN-representation",
                                "valueString": "SYL"
                            }
                        ],
                        "use": "official",
                        "text": "トウキョウ　タロウ",
                        "family": "トウキョウ",
                        "given": [
                            "タロウ"
                        ]
                    }
                ],
                "gender": "male",
                "birthDate": "2023-06-17T22:59:14+09:00",
                "address": [
                    {
                        "text": "神奈川県横浜市港区１－２－３",
                        "postalCode": "123-4567",
                        "country": "JP"
                    }
                ],
                "resourceType": "Patient"
            }
        }
    ],
    "resourceType": "Bundle"
}
```