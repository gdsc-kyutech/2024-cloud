# はじめに

ここでは，Google Cloud Platform及びGoogle Cloud Skills Boostについて解説しています．  
ご存じの方は読み飛ばしていただいて構いません．

## Google Cloud Platform の概要

Google Cloud Platform(GCP)とは，Googleが提供するクラウドコンピューティングサービスの総称です．  
<!-- Google社内で利用されているインフラを自分のサービスに活用することが可能です． -->

### クラウドコンピューティングとは

コンピュータ処理をネットワーク経由でサービスとして提供することです．  
手元に実機を用意する必要のある従来型のシステム構成と比較して，様々なメリットがあります．

- ネットワーク経由で利用するため，いつでもどこでも使うことができる
- 必要なときに必要なだけ利用できるため，状況に応じて使用量を増減できる
- クラウド特有の便利な機能を使用することができる

### Google Cloud Platformの代表的なサービス例

- **Compute Engine**  
    CPUやGPUなどのスペックを指定して，自由に仮想マシン(VM)を利用できる．
- **Cloud Storage**  
    あらゆるデータを保存できるオブジェクトストレージ．
- **BigQuery**  
    多様なデータを分析できるプラットフォーム．
- **Cloud Run**  
    コンテナ化されたアプリケーションを容易に動かすことのできる実行環境．Dockerを動作させるまでの要素を気にせずにコンテナを動作させることができる．
- **Kubernetes Engine**  
    Kubernetes実行環境．Linuxサーバなど，Docker/Kubernetesを動かすために必要な要素を気にせずにコンテナ群を動かすことができる．
- **Cloud Build**  
    実装したコードを迅速に適用(デプロイ)するためのプラットフォーム．

## Google Cloud Skills Boost の概要

[Google Cloud Skills Boost](https://www.cloudskillsboost.google/) とは，Google Cloud Platformをかんたんに学べるプラットフォームです．  
目的別に多様なコースが用意されており，自身のレベルに応じたコースを選んで，自分のペースでGoogle Cloud Platformの仕組みや操作を学ぶことができます．

それぞれのコースには動画・クラウドリソースへ実際にアクセスするハンズオン・理解度チェックのテストなどが含まれています．  
Creditが必要な有料のコースと費用のかからない無料のコースがあり，今回のハンズオンでは無料のコースを使用して進行します．

2024年6月現在，学生は[こちらのフォーム](https://services.google.com/fb/forms/googlecloudskillsbooststudenttrainingcreditsapplication/)から申請することで 1年間有効な 200 Creditsを無料で得ることができます．  
フォームには英語で記す必要があります．以下のように入力してください．

- Email Address：大学から配布された`kyutech.jp`で終わるメールアドレス
- Name of Academic Institution：`Kyushu Institute of Technology`

手動で確認・承認が行われるため，申請してから利用可能になるまで約2〜3週間程度かかります．

[1. 仮想マシン(VM)ベースアーキテクチャによるREST API実行環境の構築](1-vm.md)へ進む

[目次に戻る](README.md)
