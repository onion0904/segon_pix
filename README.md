# segon_pix

## 説明
これは画像投稿サイトであり、様々な機能の実装を経験するためのリポジトリである。

## 実装予定の機能
- 画像表示画面
    - いいね
    - コメント
    - 絵師のプロフィールに移動 

- プロフィール画面
    - フォロー
    - 共有
    - 絵の一覧表示
    - 上にはヘッダー画像
    - いいねした画像一覧

- 画像の投稿
    - 複数枚可能
    - ハッシュタグで紐づけ


- ログイン


## APIの機能

| MethodName  | Input                   | Output                        | Explanation                                         |
|-------------|-------------------------|-------------------------------|-----------------------------------------------------|
| UserInfo    | userid(uint)            | *models.User, error           | 与えられたidのユーザー情報を返す                      |
| SearchImage | Qhashtag(string)        | []models.PostedImage, error   | 与えられたハッシュタグの部分一致の画像のスライスを返す |
| AddUser     | model(*models.User)     | error                         | ユーザーを追加する                                    |
| DeleteUser  | userID(uint)            | error                         | ユーザーとその投稿画像を削除する                       |
| AddPostedImage     | ctx(context.Context), file(io.Reader), filename(string), userID(uint), hashtags([]models.Hashtag)| error | GCSへのアップロードを伴う投稿画像の追加を処理します             |
| DeletePostedImage  | ctx(context.Context), imageID(uint)                                                              | error | GCSの投稿画像と対応するファイルの削除を処理します                |
| AddLike    | userID(uint), imageID(uint)                                                                              | error | ユーザーが画像にいいねをする処理     |
| RemoveLike | userID(uint), imageID(uint)                                                                              | error | ユーザーが画像のいいねを取り消す処理 |
| AddComment     | model(*models.Comment), imageID(uint)              | error  | 新しいコメントをPostedImageに追加する                |
| UpdateComment  | commentID(uint), newContent(string), imageID(uint) | error  | 指定されたコメントの内容を更新する                   |
| DeleteComment  | commentID(int), imageID(uint)                      | error  | コメントを削除し、PostedImageからもコメントを削除する |
