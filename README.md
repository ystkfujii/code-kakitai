# Go言語で構築するクリーンアーキテクチャ設計

このリポジトリは[『Go言語で構築するクリーンアーキテクチャ設計』](
https://techbookfest.org/product/9a3U54LBdKDE30ewPS6Ugn)に出てくるサンプルアプリケーションのリポジトリになります。
書籍では一部のコードしか記載できなかったため、こちらのリポジトリでより詳しくコードを確認できます。

## この書籍について
[『Go言語で構築するクリーンアーキテクチャ設計』](
https://techbookfest.org/product/9a3U54LBdKDE30ewPS6Ugn)
「Go言語で構築するクリーンアーキテクチャ設計」は、Go言語を使用したアプリケーション開発においてクリーンアーキテクチャの原則をどのように適用するかを解説した実践的なガイドです。本書を通じて、読者は以下のような知見を得ることができます：

 - そもそもクリーンアーキテクチャとは？
 - パッケージ構成と各レイヤーの責務がわかる
 - モックを活用したレイヤーの単体テストが知れる
 - ユースケースでトランザクション制御するには？
 - ドメインモデルとドメインサービスの実装が掴める

## 動作確認

本リポジトリのコードを動かすためには以下の手順で環境構築を行ってください：

```bash
make init
```
こちらのコマンドで各種コンテナの起動やDBのマイグレーションが行われます。

```bash
make run
```
こちらのコマンドで、Goサーバーの起動が行われます。


Swaggerを用いてAPIドキュメントを確認する場合は以下のコマンドを実行します：
```bash
make swagger
```
こちらのコマンドで、Swaggerのコンテナが起動します。
localhost:8080にて確認可能です。

## 著者
- hiroaki
- sugamaan
- gyuu
- tomoya