# 1. Compute Engine による VM のセットアップと REST API の構築

## 全体手順

Skills Boost の注意点
- 初回はNoCostのラボでも二回目以降はCreditが必要である
- ラボで必要ではないリソースや必要数以上のリソースを使用するとアカウントがブロックされる場合がある
- 制限時間があるため余裕持って終われるようにする

以下の箇条書きを補足し、項目ごとに分離してドキュメント化する。

1. [Google Cloud Skills Boost - Compute Engine を使用した Google Cloud でのウェブアプリのホスティング](https://www.cloudskillsboost.google/course_templates/638/labs/480366?locale=ja)
    1. `設定と要件` をやる
        1. `リージョンとゾーンを設定する` はやらなくて良い
1. [Compute Engine での Go のスタートガイド](https://cloud.google.com/go/getting-started/getting-started-on-compute-engine?hl=ja)
    1. `目標`を読む
    1. `料金`を読む（今回はSkillsBoostアカウントを作成して使用しているため、初回は無料で2回目以降が1クレジット）
    1. `準備`を行う
        1. `プロジェクト セレクタに移動`をクリックし、プロジェクトをクリックし、プロジェクト ID をメモしておく
            1. 例 : qwiklabs-gcp-02-84ace8ee5776
    1. `APIを有効にする` をクリックし、Compute Engine, Cloud Build API を有効化する
    1. `Cloud Shellに移動`をクリックし、環境のセットアップ（git clone）を行う
    1. `gcloud config set project YOUR_PROJECT_ID` を実行する
        1. YOUR_PROJECT_ID はメモしておいたプロジェクト ID
        1. ガイドページの環境変数設定機能を使うと便利（赤いペン）
    1. `Cloud Shell でアプリを実行する` をやる
        1. これをやることによって、VM内のローカル環境で動作確認が行える
            1. `ポート8080でプレビュー` でHelloWorld!を確認できたら成功
    1. `単一インスタンスへのデプロイ` をやる
        1. ビルドを行う前に、ファイルを修正する
            1. [golang-samples/getting-started/gce/startup-script.sh](./source/1/startup-script.sh)を編集する
                1. Logging機能が提供終了している？ためか動作しないため
            2. [golang-samples/getting-started/gce/cloudbuild.yaml](./source/1/cloudbuild.yaml)を編集する
                1. バージョン不整合によりコンパイルエラーが発生しないようにビルド環境と実行環境を合わせる
        2. ZONEは `us-central1-a` を使用する（なんでもいい）
        3. 外部 IP アドレスからHelloWorld!を確認できたら成功
    1. [golang-samples/getting-started/gce/main.go](./source/1/main.go)を編集する
        1. `/v1/customer/22530`にアクセスするとJsonが表示される→成功
            1. [Google Cloud Skills Boost - Developing a REST API with Go and Cloud Run](https://www.cloudskillsboost.google/course_templates/741/labs/464421) と同様の動作が確認できる
                1． アクセスできない場合は、[Google Cloud Skills Boost - Develop Serverless Applications on Cloud Run](https://www.cloudskillsboost.google/course_templates/741) の左メニューから`Go と Cloud Run を使用した REST API の開発`をクリック

## Google Cloud Platform の概要


## Google Cloud Skills Boost の概要


## Google Cloud Skills Boost - Compute Engine を使用した Google Cloud でのウェブアプリのホスティング


## Compute Engine での Go のスタートガイド


## DBの導入やスケーラビリティの実現について

Google Cloud Compute Engine によって、VM のセットアップと REST API の構築が完了しました。

しかし、現状のmain.goはコード内に直接データを記述しているため、アルゴリズムとデータが密結合している状態です。また、データベースを導入していないため、データの永続化やアクセスの管理が難しい状況です。

これを解決するためには、データベースを導入し、データとアルゴリズムを分離することが重要です。

また、スケーラビリティの観点からも、現状の Compute Engine では、負荷が増加した際にスケールアウトが難しいという課題があります。

これらの課題を解決するために、次のステップとして、Cloud Run と Firestore によるサーバーレス REST API の構築を行いましょう。

[2. Cloud Run と Firestore によるサーバーレスREST APIの構築](2-sv-less.md)へ進む

[目次に戻る](README.md)
