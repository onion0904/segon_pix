# POST

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
    - 入力できる値(queryでの受け取り)
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
