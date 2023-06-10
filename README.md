# Go言語での FHIRのサンプル

- [診療情報提供書HL7FHIR記述仕様](https://std.jpfhir.jp/)に基づくサンプルデータ（Bundle-BundleReferralExample01.json）に対して、FHIRプロファイルでの検証、パースするサンプルプログラムです。
- FHIRプロファイルでの検証
    - Goの場合、[HL7のConfluenceのページ「Open Source Implementations」](https://confluence.hl7.org/display/FHIR/Open+Source+Implementations)で紹介されている、FHIRのリファレンス実装がありません。
    - また、その他、検索しても、Goでは、JavaのHAPI等と違い、FHIRの構造定義ファイルでの検証を行うライブラリがなさそうです。
    - ですが、[HL7 FHIR v4.0.1:R4のダウンロードページ](https://hl7.org/fhir/R4/downloads.html)に[JSON Schema形式のファイル](https://hl7.org/fhir/R4/fhir.schema.json.zip)が提供されています。        
    - そこで、JSONスキーマによる検証ができる[gojsonschema](https://github.com/xeipuuv/gojsonschema)というライブラリを使って、JSONスキーマの検証をしています。
- 【未実施】JPCoreプロファイル、文書プロファイルでの検証
    - [JPCoreプロファイル](https://jpfhir.jp/fhir/core/)のサイトにJPCoreプロファイルの構造定義ファイルがあります。
    - また、[https://std.jpfhir.jp/](https://std.jpfhir.jp/)のサイトに、JPCoreを含むスナップショット形式の[診療情報提供書の文書プロファイル（IGpackage2023.4.27 snapshot形式: jp-ereferral-0.9.6-snap.tgz）](https://jpfhir.jp/fhir/eReferral/jp-ereferral-0.9.7-snap.tgz)があります。 
    - いずれも、[FHIR package仕様](https://registry.fhir.org/learn)に従ったパッケージです。    
    - ですが、FHIRプロファイルのようなJSONスキーマ形式では提供されていないようで、Goの場合、JPCoreプロファイル、文書プロファイルレベルの検証を実施するのが難しそうです。

- FHIRデータのパース
    - 前述の通り、Goの場合、[HL7のConfluenceのページ「Open Source Implementations」](https://confluence.hl7.org/display/FHIR/Open+Source+Implementations)で紹介されている、FHIRのリファレンス実装がありません。
    - ですが、検索すると、[Golang FHIR Models](https://github.com/samply/golang-fhir-models)というライブラリが出てきたので、それを使って、パースをしてみました。
        - JavaのHAPI FHIRのような実装と比べると、バリデータ機能もないですし、コミュニティとしての信頼性も低いと考えています。
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