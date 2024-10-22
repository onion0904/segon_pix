# POST

- /add/user(token必要)

    - 入力できる値(JSONでの受け取り)
    - 絶対に
    - ユーザーを登録する。

    | 変数      |     説明     |
    |-----------|-----------|
    |Name       |  ユーザー名   |
    |Description（任意）   |  プロフィール欄に書くメッセージ  |
    |Email      |  メールアドレス  |
    |Password       | パスワード  |
    |Birthday   |  誕生日      |

    - 使用例

```bash
curl -X POST -H "Authorization: Bearer <JWTトークン>" \
-H "Content-Type: application/json" \
-d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password",
    "birthday": 20241015
}' \
http://localhost:8080/segon_pix_auth/add/user
```


- /add/image(token必要)
    - 入力できる値
    - 画像を追加する。ユーザーのPostedImageにも追加する。ハッシュタグは検索するときに使う。

    | 変数      | 説明|
    |-----------|-----|
    |ID       |   ユーザーID   |
    |File    |   画像ファイル   |
    |Hashtags     | ハッシュタグ（複数可） |


    - 使用例

```bash
curl -X POST -H "Authorization: Bearer <JWTトークン>" \
"http://localhost:8080/segon_pix_auth/add/image?ID=1234" \
-F "File=@/path/to/your/image.jpg" \
-F "Hashtags=tag1" \
-F "Hashtags=tag2" \
-F "Hashtags=tag3"
```

- /add/like(token必要)
    - 入力できる値(queryでの受け取り)
    -指定された画像にユーザー情報を入れる。ユーザー情報にいいねした画像を追加する。

    | 変数      |     説明     |
    |-----------|-----------|
    |userID     |  ユーザーID  |
    |imageID    |  画像ID  |

    - 使用例

    ```
    curl -X POST -H "Authorization: Bearer <JWTトークン>" \"http://localhost:8080/segon_pix_auth/add/like?userID=1234&imageID=5678"
    ```


- /add/comment(token必要)
    - 入力できる値(queryでの受け取り)
    -指定された画像にユーザー情報とコメントを追加する。

    | 変数      |     説明     |
    |-----------|-----------|
    |userID     |  ユーザーID  |
    |imageID    |  画像ID  |
    |comment | コメント内容|

    - 使用例

    ```
    curl -X POST -H "Authorization: Bearer <JWTトークン>" \"http://localhost:8080/segon_pix_auth/add/comment?userID=1234&imageID=5678&comment=aiueo"
    ```
