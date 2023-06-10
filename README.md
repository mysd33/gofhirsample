# Go言語での FHIRのサンプル

- [Golang FHIR Models](https://github.com/samply/golang-fhir-models)とGo標準のJSONライブラリ使って、[診療情報提供書HL7FHIR記述仕様](https://std.jpfhir.jp/)に基づくサンプルデータ（Bundle-BundleReferralExample01.json）をパースするサンプルプログラムです。
- さらに、[HL7 FHIR v4.0.1:R4のダウンロードページ](https://hl7.org/fhir/R4/downloads.html)にある[JSON Schema](https://hl7.org/fhir/R4/fhir.schema.json.zip)と、[gojsonschema](https://github.com/xeipuuv/gojsonschema)のライブラリを使って、パースする前に、サンプルデータのJSONスキーマの検証も実施しています。

- 使い方
    - ビルド後、生成されたexeファイルを実行してください。
```sh
# ビルド
go build
# 実行
example.exe
```

```sh
# 実行結果の例
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