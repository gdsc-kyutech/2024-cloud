# 2. Cloud Run と Firestore によるサーバーレスREST APIの構築

[Google Cloud Skills Boost - Developing a REST API with Go and Cloud Run](https://www.cloudskillsboost.google/course_templates/741/labs/464421) に取り組んで、Cloud Run と Firestore によるサーバーレス REST API の構築を行いましょう。

アクセスできない場合は、日本語版の[Google Cloud Skills Boost - Go と Cloud Run を使用した REST API の開発](https://www.cloudskillsboost.google/course_templates/741/labs/463386)を参照してください。

## 注意

日本語版ではDockefileにバグがあり、以下のようなエラーがでます。

```sh
Revision '' is not ready and cannot serve traffic. The user-provided container failed to start and listen on the port defined provided by the PORT=8080 environment variable. Logs for this revision might contain more information.
```

このエラーは、Dockerfileのdebianのバージョンが古く、Goが使用しているGLIBCのバージョンと合わないためビルドに失敗し、サーバーが起動できていないというものです。（参考 : https://tech-lab.sios.jp/archives/35991 ）

英語版ではdebian12を使用しているので、Dockerfileを以下のように修正してください。

```Dockerfile
FROM gcr.io/distroless/base-debian12
WORKDIR /usr/src/app
COPY server .
CMD [ "/usr/src/app/server" ]
```

また、日本語版ではタスク7が最後まで閲覧できない不具合があります。タスク7の内容を以下に示すので、参考にしてください。

### タスク 7. 新しいリビジョンをデプロイする

1. ソースコードを再ビルドします。
    ```sh
    go build -o server
    ```

1. REST APIの新しいイメージをビルドします。
    ```sh
    gcloud builds submit \
    --tag gcr.io/$GOOGLE_CLOUD_PROJECT/rest-api:0.2
    ```
1. 更新したイメージをデプロイします。

    ```sh
    gcloud run deploy rest-api \
    --image gcr.io/$GOOGLE_CLOUD_PROJECT/rest-api:0.2 \
    --platform managed \
    --region "Filled in at lab startup." \
    --allow-unauthenticated \
    ```

1. デプロイが完了すると、以前と同様のメッセージが表示されます。新しいバージョンのデプロイ時には、REST API の URL は変わりませんでした。

1. すでにそのURLで開いてあるブラウザタブに戻ります（末尾に「`/v/`」が付いています）。更新して、APIステータスがまだ実行中であることを示す、以前と同じメッセージが表示されることを確認します。
    ![](https://cdn.qwiklabs.com/Q6zP6oeJbelXp7M24egmn5kOqwM%2FjM2udtHAx9S9k9c%3D)

1. ブラウザのアドレスバーに表示されたアプリケーションURLに「`/customer/22530`」を追加します。次のJSONレスポンスを受け取ります。顧客に関して提案、承認、拒否された治療の合計額がそれぞれ表示されるはずです。
    ![](https://cdn.qwiklabs.com/YA7BORLOP0WqHewA%2B2vxM2Gll12QkAbxZ%2F7PN2IblYI%3D)

1. 22530の代わりに次のような別の顧客IDをURLに入力してみましょう。
  - 34216
  - 70156（金額はすべてゼロになります）
  - 12345（顧客やペットは存在しないため、**Query is null**のようなエラーが返されます）

データベースからの読み取りを行う、スケーラブルでメンテナンスをあまり必要としないサーバーレスREST APIをビルドしました。

[目次に戻る](README.md)
