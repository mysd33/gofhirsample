# Go言語での FHIRのサンプル

- [診療情報提供書HL7FHIR記述仕様](https://std.jpfhir.jp/)に基づくサンプルデータ（Bundle-BundleReferralExample01.json）に対して、FHIRプロファイルでの検証、パースするサンプルプログラムです。
- また、FHIRリソース(Bundle)として作成したオブジェクトを、FHIRのJSON文字列で出力（シリアライズ）するサンプルプログラムもあります。
## プロファイルの検証（バリデーション）とパース
- FHIRプロファイルでの検証
    - Goの場合、[HL7のConfluenceのページ「Open Source Implementations」](https://confluence.hl7.org/display/FHIR/Open+Source+Implementations)で紹介されている、FHIRのリファレンス実装がありません。このため、[FHIR package仕様](https://registry.fhir.org/learn)に従ったnpmパッケージ形式での検証方法は、難しそうです。
    - ですが、[HL7 FHIR v4.0.1:R4のダウンロードページ](https://hl7.org/fhir/R4/downloads.html)に、[JSON Schema形式のファイル](https://hl7.org/fhir/R4/fhir.schema.json.zip)が提供されています。そこで、JSONスキーマによる検証ができる[gojsonschema](https://github.com/xeipuuv/gojsonschema)というライブラリを使って、JSON Schemaによる検証をしてみました。
    - なお、HL FHIRでのバリデーションは複数の方法が提供されており、JSON Schemaもその1つです。方法によって検証可能な内容が若干異なり、公式Validator等に比べるとJSON Schemaで検証できる内容は限定されるようです。
        - [HL7 FHIR:R4 Validating Resources](http://hl7.org/fhir/R4/validation.html)

- 【未実施】JPCoreプロファイル、文書情報プロファイルでの検証
    - JSON Schemaでの提供がされていないことから、上記と同様の理由で、 [JPCoreプロファイル](https://jpfhir.jp/fhir/core/)、[https://std.jpfhir.jp/](https://std.jpfhir.jp/)にある[診療情報提供書の文書情報プロファイル（IGpackage2023.4.27 snapshot形式: jp-ereferral-0.9.6-snap.tgz）](https://jpfhir.jp/fhir/eReferral/jp-ereferral-0.9.7-snap.tgz)レベルの検証を実施するのが難しそうです。

- FHIRデータのパース
    - 前述の通り、Goの場合、FHIRのリファレンス実装がありません。
    - ですが、検索すると、[Golang FHIR Models](https://github.com/samply/golang-fhir-models)というライブラリが出てきたので、それを使って、パースをしてみました。
        - JavaのHAPI FHIRのような実装と比べると、バリデータ機能もないですし、コミュニティとしての信頼性も低いため、このライブラリを正式に使うのは推奨できませんが、試してみました。
        - 上記は、内部では、[Go標準のJSONライブラリ(encoding/json)](https://pkg.go.dev/encoding/json)を使用します。

## シリアライズ
- FHIRデータからJSONへのシリアライズ
    - パース同様、[Golang FHIR Models](https://github.com/samply/golang-fhir-models)および、[Go標準のJSONライブラリ(encoding/json)](https://pkg.go.dev/encoding/json)を使って、シリアライズします。

## 注意事項
- なお、[Golang FHIR Models](https://github.com/samply/golang-fhir-models)はHAPIのようなリファレンス実装と比較しての信頼性、今後のR5等のFHIRバージョンアップ対応等の将来性が保証がされないことから、あまりおすすめできない実装手段かと感じています。 [別のリポジトリのgoのサンプルAP](https://github.com/mysd33/gofhirsample2)では、別の手段として、汎用的なJSONライブラリのみでFHIRのパースを実現できないかを検討していますので、こちらを参照ください。
    - この場合、シリアライズは対応できません。

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
>parsing-example.exe
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
- 注意事項
    - [Golang FHIR Models](https://github.com/samply/golang-fhir-models)に、json.Marshalの振る舞いにより「<」や「>」が`\u003c` and `\u003e`になってしまう不具合があり、修正されておらず、自分で直接、生成コードを直さないといけないようです。
        - https://github.com/samply/golang-fhir-models/issues/13
    - ですので、やはり、このライブラリを正式に使うのは推奨できないです。

```sh
>serializing-example.exe
{
    "meta": {
        "lastUpdated": "2023-06-18T17:20:40+09:00",
        "profile": [
            "http://jpfhir.jp/fhir/ePrescription/StructureDefinition/JP_Bundle_ePrescriptionData"
        ]
    },
    "type": "document",
    "entry": [
        {
            "fullUrl": "urn:uuid:1719e3a6-3473-4c2b-979c-b6792756dcb2",
            "resource": {
                "meta": {
                    "lastUpdated": "2023-06-18T17:20:40+09:00",
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
                    "reference": "urn:uuid:9ab62a02-5d14-41f3-93bd-34ecd8772df9"
                },
                "date": "2023-06-18T17:20:40+09:00",
                "author": [],
                "title": "処方箋",
                "resourceType": "Composition"
            }
        },
        {
            "fullUrl": "urn:uuid:9ab62a02-5d14-41f3-93bd-34ecd8772df9",
            "resource": {
                "meta": {
                    "lastUpdated": "2023-06-18T17:20:40+09:00",
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
                        "system": "http://jpfhir.jp/fhir/ccs/Idsysmem/JP_Insurance_member/00012345",
                        "value": "00012345:あいう:１８７:05"
                    }
                ],
                "name": [
                    {
                        "extension": [
                            {
                                "url": "http://hl7.org/fhir/StructureDefinition/iso21090-EN-representation",
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
                                "url": "http://hl7.org/fhir/StructureDefinition/iso21090-EN-representation",
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
                "birthDate": "1920-02-11",
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