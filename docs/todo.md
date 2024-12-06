# TODO

## client

-
- home実装
- setting実装
- search実装

## server

- 変更点
    - 認証が必要なエンドポイントにuserIDを追加してIssuesを解決
    - 今までIDとかだったエンドポイントをuserID,imageIDとかに変更
    - VerifyをVerifyAddUserに変更し、AddUserエンドポイントを削除。VerifyAddUserでUserを作成
- TODO
    - 機能を整理してディレクトリに分ける
    - main.goの処理を他のディレクトリに分ける