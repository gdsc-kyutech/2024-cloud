# 2. Cloud Run と Firestore によるサーバーレスREST APIの構築

[Google Cloud Skills Boost - Developing a REST API with Go and Cloud Run](https://www.cloudskillsboost.google/course_templates/741/labs/464421) に取り組んで、Cloud Run と Firestore によるサーバーレス REST API の構築を行いましょう。

アクセスできない場合は、日本語版の[Google Cloud Skills Boost - Go と Cloud Run を使用した REST API の開発](https://www.cloudskillsboost.google/course_templates/741/labs/463386)を参照してください。

## 注意

日本語版ではDockefileにバグがあり、以下のようなエラーがでます。

```sh
Revision '' is not ready and cannot serve traffic. The user-provided container failed to start and listen on the port defined provided by the PORT=8080 environment variable. Logs for this revision might contain more information.
```

このエラーは、Dockerfileのdebianのバージョンが古く、Goが使用しているGLIBCのバージョンと合わないためビルドに失敗したため、サーバーが起動できていないというものです。（参考 : https://tech-lab.sios.jp/archives/35991 ）

英語版ではdebian12を使用しているので、Dockerfileを以下のように修正してください。

```Dockerfile
FROM gcr.io/distroless/base-debian12
WORKDIR /usr/src/app
COPY server .
CMD [ "/usr/src/app/server" ]
```


[目次に戻る](README.md)
