# export

[Develop Serverless Applications on Cloud Run](https://www.cloudskillsboost.google/course_templates/741/labs/464421) で使用しているサンプルデータをJsonに書き起こす

出力したサンプルデータである[customers.json](../customers.json) はすでにレポジトリ内に格納されているため、このコードを再度実行する必要はない

1. "Task 3. Import test customer data" を手順通りに行い、Firestoreにデータを格納する
2. [main.go](./main.go) を実行し、Firestoreからデータを取得し、Jsonファイルに書き出す
3. 同階層に生成された `customers.json` がサンプルデータとなる
