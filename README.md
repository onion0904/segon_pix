# segon_pix

## 説明
これは画像投稿サイトであり、様々な機能の実装を経験するためのリポジトリである。


## リクエストの仕方

### POST

- /add/user

    - 入力できる値(JSONでの受け取り)

    | 変数      |     説明     | 
    |-----------|-----------|
    |Name       |  ユーザー名   |       
    |Profile(任意)    |  プロフィール欄に書くメッセージ  |      
    |Email(任意)      |  メールアドレス  |      
    |Birthday   |  誕生日      |

    - 使用例

    ```
    curl -X POST https://localhost:8080/add/user \
    -H "Content-Type: application/json" \
    -d '{
        "name": "John Doe",
        "email": "john@example.com",
      }'
    ```
- /add/image
    - 入力できる値

    | 変数      | 説明|
    |-----------|-----|
    |ID       |   ユーザーID   |
    |File    |   画像ファイル   |
    |Hashtags     | ハッシュタグ（複数可） | 


    - 使用例

    ```
    curl -X POST "https://localhost:8080/add/image/images?ID=1234" \
  -F "File=@/path/to/your/image.jpg" \
  -F "Hashtags=tag1" \
  -F "Hashtags=tag2" \
  -F "Hashtags=tag3"

    ```
- /add/like
    - 入力できる値

    | 変数      |     説明     | 
    |-----------|-----------|
    |userID     |  ユーザーID  |       
    |imageID    |  画像ID  |      

    - 使用例

    ```
    curl -X POST "https://localhost:8080/add/like?userID=1234&imageID=5678"
    ```
- /add/comment

### GET

- /list/user
    - 入力できる値

    | 変数      |     説明     | 
    |-----------|-----------|
    |ID       |  ユーザーID   |       

    - 使用例

    ```
    curl -X GET "https://localhost:8080/list/user/info?ID=1234"
    ```
- /list/image
    - 入力できる値

    | 変数      |     説明     | 
    |-----------|-----------|
    |Hashtag       |  検索したいワード   |       

    - 使用例

    ```
    curl -X GET "https://localhost:8080/list/image/search?Hashtag=旅行"
    ```

### PUT

- /update/comment
    - 入力できる値

    | 変数      |     説明     | 
    |-----------|-----------|
    |commentID       |  コメントのID  |       
    |imageID    |  画像のID  |      
    |newContent       |  更新後のコメントの内容  |      


    - 使用例

    ```
   curl -X PUT "https://localhost:8080/update/comment?commentID=5678&imageID=1234&newContent=更新されたコメント内容"
    ```

### DELETE

- /delete/user
    - 入力できる値

    | 変数      |    説明     | 
    |-----------|-----------|
    |ID       |  ユーザーID   |       

    - 使用例

    ```
    curl -X DELETE "https://localhost:8080/delete/user?ID=1234"
    ```
- /delete/image
    - 入力できる値

    | 変数      |     説明     | 
    |-----------|-----------|
    |ID       |   画像ID  |       

    - 使用例

    ```
    curl -X DELETE "https://localhost:8080/delete/image/?ID=1234"
    ```
- /delete/like
    - 入力できる値

    | 変数      |     説明     | 
    |-----------|-----------|
    |userID     |  ユーザーID  |       
    |imageID    |  画像ID  |       

    - 使用例

    ```
    curl -X DELETE "https://localhost:8080/delete/like?userID=1234&imageID=5678"
    ```
- /delete/comment
    - 入力できる値

    | 変数      |     説明     | 
    |-----------|-----------|
    |commentID       |  コメントのID  |       
    |imageID    |  画像のID  |      
    - 使用例

    ```
    curl -X DELETE "https://localhost:8080/delete/comment?commentID=5678&imageID=1234"
    ```



## TODO

- IconとHeaderを登録する機能の実装
- AddCommentにmodelsのmessage以外も追加できるようにする
- followする機能の実装
- Update Userの実装
- ログイン機能の実装
- controllerのエラーハンドリング
- コメントを分かりやすくする