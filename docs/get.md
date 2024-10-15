# GET

- /get/user
    - 入力できる値(queryでの受け取り)
    - 与えられたユーザーの情報を返す。(passwordやEmail以外)

    | 変数      |     説明     |
    |-----------|-----------|
    |ID       |  ユーザーID   |

    - 使用例

    ```
    curl -X GET "http://localhost:8080/segon_pix/get/user?ID=1234"
    ```

- /restricted/get/user(token必要)
    - 入力できる値(queryでの受け取り)
    - このエンドポイント以下はJWTトークンが必要になる
    - Userを返す。emailとpasswordからUserを判断する

    | 変数      |     説明     |
    |-----------|-----------|
    |JWTトークン     |  verifyで返ってきたやつ(24時間使用可能)  |
    |  email   |  メールアドレス  |
    |  password   |  パスワード  |


    - 使用例

    ```
    curl -X GET -H "Authorization: Bearer <JWTトークン>" \
"http://localhost:8080/segon_pix_auth/get/user?email=john@example.com&password=password"

    ```

- /get/list/image
    - 入力できる値(queryでの受け取り)
    - 指定されたハッシュタグの部分一致を返す。
    - URLとimageIDのlistを返す。

    | 変数      |     説明     |
    |-----------|-----------|
    |Hashtag       |  検索したいワード   |

    - 使用例

    ```
    curl -X GET "http://localhost:8080/segon_pix/get/list/image?Hashtag=旅行"
    ```

- /get/image
    - 入力できる値(queryでの受け取り)
    - 指定されたimageIDの画像の情報を返す。
    - PostedImageの全部を返す。

    | 変数      |     説明     |
    |-----------|-----------|
    |imageID       |  画像のID   |

    - 使用例

    ```
    curl -X GET "http://localhost:8080/segon_pix/get/image?imageID=1234"