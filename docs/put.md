# PUT

- /update/comment(token必要)
    - 入力できる値(queryでの受け取り)
    - コメントの内容を更新する。

    | 変数      |     説明     |
    |-----------|-----------|
    |userID    |  userのID  |
    |commentID       |  コメントのID  |
    |imageID    |  画像のID  |
    |newContent       |  更新後のコメントの内容  |


    - 使用例

    ```
   curl -X PUT -H "Authorization: Bearer <JWTトークン>" \"http://localhost:8080/segon_pix_auth/update/comment?userID=1234&commentID=5678&imageID=1234&newContent=更新されたコメント内容"
    ```

- /update/user/header(token必要)
    - 入力できる値
    - ユーザー情報にheader画像を追加する
    - もともとユーザーがあるときに使える

    | 変数      | 説明|
    |-----------|-----|
    |userID       |   ユーザーID   |
    |File    |   画像ファイル   |


    - 使用例

    ```
    curl -X POST -H "Authorization: Bearer <JWTトークン>" \"http://localhost:8080/segon_pix_auth/update/user/header?userID=1234" \
  -F "File=@/path/to/your/image.jpg"
    ```

- /update/user/icon(token必要)
    - 入力できる値
    - ユーザー情報にicon画像を追加する
    - もともとユーザーがあるときに使える

    | 変数      | 説明|
    |-----------|-----|
    |userID       |   ユーザーID   |
    |File    |   画像ファイル   |


    - 使用例

    ```
    curl -X POST -H "Authorization: Bearer <JWTトークン>" \"http://localhost:8080/segon_pix_auth/update/user/icon?userID=1234" \
  -F "File=@/path/to/your/image.jpg"
    ```

- /update/user(token必要)
    - 入力できる値(queryでの受け取り)
    - コメントの内容を更新する。

    | 変数      |     説明     |
    |-----------|-----------|
    |userID       |  userのID  |
    |name       |  userのname  |
    |description       |   プロフィールメッセージ |
    |birthday      |   誕生日 |
    |email      |   email |


    - 使用例

    ```
   curl -X PUT -H "Authorization: Bearer <JWTトークン>" \"http://localhost:8080/segon_pix/update/user?userID=1234&name=onion&description=更新された内容&email=更新された内容"