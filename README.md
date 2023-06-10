# Go言語での FHIRのサンプル

- [診療情報提供書HL7FHIR記述仕様](https://std.jpfhir.jp/)に基づくサンプルデータ（Bundle-BundleReferralExample01.json）のFHIRプロファイルでの検証、パースするサンプルプログラムです。
- FHIRプロファイルでの検証
    - Goでは、JavaのHAPI等と違い、FHIRの構造定義ファイルでの検証を行うライブラリがなさそうです。
    - そこで、[HL7 FHIR v4.0.1:R4のダウンロードページ](https://hl7.org/fhir/R4/downloads.html)に[JSON Schema形式のファイル](https://hl7.org/fhir/R4/fhir.schema.json.zip)が提供されています。        
    - [gojsonschema](https://github.com/xeipuuv/gojsonschema)のライブラリを使って、JSONスキーマの検証
- 【未実施】JPCoreプロファイルでの検証
    - [JPCoreプロファイル](https://jpfhir.jp/fhir/core/)のサイトにJPCoreプロファイルの構造定義ファイルがありますが、FHIRプロファイルのようにJSONスキーマ形式では提供されていないようで、Goの場合、JPCoreプロファイルの検証を実施するのが難しそうです。
- FHIRデータのパース
    - [Golang FHIR Models](https://github.com/samply/golang-fhir-models)を使って、パースをしています。
    - 上記は、内部では、[Go標準のJSONライブラリ(encoding/json)](https://pkg.go.dev/encoding/json)を使用します。
- サンプルAPの使い方
    - ビルド後、生成されたexeファイルを実行してください。
```sh
# ビルド
go build
# 実行
example.exe
```

* 実行結果の例
```sh
>example.exe
# JSONスキーマチェック結果
2023/06/10 15:27:34 JSON Schema Check: ドキュメントは有効です
# テストデータのパース結果
2023/06/10 15:27:34 Bundle type: Document
2023/06/10 15:27:34 文書名: 診療情報提供書
2023/06/10 15:27:34 subject display: 患者リソースPatient
2023/06/10 15:27:34 subject reference Id: urn:uuid:0a48a4bf-0d87-4efb-aafd-d45e0842a4dd
2023/06/10 15:27:34 subject reference type: Patient
2023/06/10 15:27:34 患者番号: 12345
2023/06/10 15:27:34 患者氏名: 田中　太郎
2023/06/10 15:27:34 患者カナ氏名: タナカ　タロウ
```