# segon_pix

## Contributers
- フロントエンド担当segnities007
- バックエンド担当onion0904

## 説明
これは画像投稿アプリである。


## バックエンド

### 概要 
初めてのチーム開発として、友人と共同で画像投稿アプリ「segon_pix」を開発しました。お互いに大きなアウトプット経験がなかったため、開発プロセス全体を学ぶことも目的としました。Pixivを参考に、ユーザー登録・画像投稿・いいね機能などを実装しました。 

### 工夫した点 
フロントエンドとの連携が初めてだったため、APIの使い方を詳細にまとめたドキュメント（/docs ディレクトリ）を作成しました。各エンドポイントの説明・リクエスト/レスポンス例を記載し、円滑な連携を意識しました。 
→ APIドキュメント:https://github.com/onion0904/segon_pix/tree/main/docs

フロントエンド側がAPIを実機で確認できるよう、ngrokを用いてローカル環境を一時的にインターネット公開しました。

### 使用技術
- Go,echo,gorm,MySQL,docker,GCS


## フロントエンド
フロントエンドを以下に移行しました。
[segon_pix_android](https://github.com/segnities007/segon_pix_android)