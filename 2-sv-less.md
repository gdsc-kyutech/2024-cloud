# 2. Cloud Run と Firestore によるサーバーレスREST APIの構築

[Google Cloud Skills Boost - Developing a REST API with Go and Cloud Run](https://www.cloudskillsboost.google/course_templates/741/labs/464421) に取り組んで、Cloud Run と Firestore によるサーバーレス REST API の構築を行いましょう。

## 注意

日本語版ではDockefileにバグがあります。以下のDockerfileに修正してください。

```Dockerfile
FROM gcr.io/distroless/base-debian12
WORKDIR /usr/src/app
COPY server .
CMD [ "/usr/src/app/server" ]
```


[目次に戻る](README.md)
