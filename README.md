# segon_pix

## 説明
これは画像投稿サイトであり、様々な機能の実装を経験するためのリポジトリである。


## リクエストの仕方

### POST

- /add/user

    - 入力できる値(JSONでの受け取り)
    - ユーザーを登録する。

    | 変数      |     説明     | 
    |-----------|-----------|
    |Name       |  ユーザー名   |       
    |Profile(任意)    |  プロフィール欄に書くメッセージ  |      
    |Email(任意)      |  メールアドレス  |      
    |Birthday   |  誕生日      |

    - 使用例

    ```
    curl -X POST http://localhost:8080/segon_pix/add/user \
    -H "Content-Type: application/json" \
    -d '{
        "name": "John Doe",
        "email": "john@example.com"
      }'
    ```
- /add/image
    - 入力できる値
    - 画像を追加する。ユーザーのPostedImageにも追加する。ハッシュタグは検索するときに使う。

    | 変数      | 説明|
    |-----------|-----|
    |ID       |   ユーザーID   |
    |File    |   画像ファイル   |
    |Hashtags     | ハッシュタグ（複数可） | 


    - 使用例

    ```
    curl -X POST "http://localhost:8080/segon_pix/add/image?ID=1234" \
  -F "File=@/path/to/your/image.jpg" \
  -F "Hashtags=tag1" \
  -F "Hashtags=tag2" \
  -F "Hashtags=tag3"

    ```
- /add/like
    - 入力できる値
    -指定された画像にユーザー情報を入れる。ユーザー情報にいいねした画像を追加する。

    | 変数      |     説明     | 
    |-----------|-----------|
    |userID     |  ユーザーID  |       
    |imageID    |  画像ID  |      

    - 使用例

    ```
    curl -X POST "http://localhost:8080/segon_pix/add/like?userID=1234&imageID=5678"
    ```
- /add/comment

### GET

- /get/user
    - 入力できる値
    - 与えられたユーザーの情報を返す。

    | 変数      |     説明     | 
    |-----------|-----------|
    |ID       |  ユーザーID   |       

    - 使用例

    ```
    curl -X GET "http://localhost:8080/segon_pix/get/user?ID=1234"
    ```
- /get/list/image
    - 入力できる値
    - 指定されたハッシュタグの部分一致を返す。

    | 変数      |     説明     | 
    |-----------|-----------|
    |Hashtag       |  検索したいワード   |       

    - 使用例

    ```
    curl -X GET "http://localhost:8080/segon_pix/get/list/image?Hashtag=旅行"
    ```

### PUT

- /update/comment
    - 入力できる値
    - コメントの内容を更新する。

    | 変数      |     説明     | 
    |-----------|-----------|
    |commentID       |  コメントのID  |       
    |imageID    |  画像のID  |      
    |newContent       |  更新後のコメントの内容  |      


    - 使用例

    ```
   curl -X PUT "http://localhost:8080/segon_pix/update/comment?commentID=5678&imageID=1234&newContent=更新されたコメント内容"
    ```

### DELETE

- /delete/user
    - 入力できる値
    - ユーザー情報を削除する。

    | 変数      |    説明     | 
    |-----------|-----------|
    |ID       |  ユーザーID   |       

    - 使用例

    ```
    curl -X DELETE "http://localhost:8080/segon_pix/delete/user?ID=1234"
    ```
- /delete/image
    - 入力できる値
    - 画像をDBとGCSから削除する。

    | 変数      |     説明     | 
    |-----------|-----------|
    |ID       |   画像ID  |       

    - 使用例

    ```
    curl -X DELETE "http://localhost:8080/segon_pix/delete/image?ID=1234"
    ```
- /delete/like
    - 入力できる値
    - いいねを取り消す。ユーザー情報のいいね欄からも消す。

    | 変数      |     説明     | 
    |-----------|-----------|
    |userID     |  ユーザーID  |       
    |imageID    |  画像ID  |       

    - 使用例

    ```
    curl -X DELETE "http://localhost:8080/segon_pix/delete/like?userID=1234&imageID=5678"
    ```
- /delete/comment
    - 入力できる値
    - コメントを消す

    | 変数      |     説明     | 
    |-----------|-----------|
    |commentID       |  コメントのID  |       
    |imageID    |  画像のID  |      
    - 使用例

    ```
    curl -X DELETE "http://localhost:8080/segon_pix/delete/comment?commentID=5678&imageID=1234"
    ```



## TODO

- IconとHeaderを登録する機能の実装
- AddCommentにmodelsのmessage以外も追加できるようにする
- followする機能の実装
- Update Userの実装
- ログイン機能の実装
- controllerのエラーハンドリング
- コメントを分かりやすくする