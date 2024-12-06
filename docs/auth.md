# auth

- /signup

    - 入力できる値
    - メアドを入力して認証コードをメアド先に送信するに変更

    | 変数      |     説明     |
    |-----------|-----------|
    |Email     |  メールアドレス  |

    - 使用例

    ```
    curl -X POST http://localhost:8080/signup?email=user@example.com
    ```

- /verifyAddUser

    - 入力できる値(JSONでの受け取り)
    - 認証コードから判断してtokenを返す

    | 変数      |     説明     |
    |-----------|-----------|
    |Name       |  ユーザー名   |
    |Description（任意）   |  プロフィール欄に書くメッセージ  |
    |Email      |  メールアドレス  |
    |Password       | パスワード  |
    |Birthday   |  誕生日      |
    |code|認証コード|

    - 使用例

    ```bash
    curl -X POST \
    -H "Content-Type: application/json" \
    -d '{
    "name": "onion gratin",
    "email": "user@example.com",
    "password": "password",
    "birthday": 20241015
    }' \
    http://localhost:8080/verifyAddUser?code=ABCDEF
    ```

- /login

    - 入力できる値
    - メアドとパスワードが一致するときにtokenを返す

    | 変数      |     説明     |
    |-----------|-----------|
    |Email     |  メールアドレス  |
    |password|パスワード|

    - 使用例

    ```
    curl -X POST http://localhost:8080/login?email=user@example&password=pas
    ```

- /restricted(token必要)
    - 入力できる値(JSONでの受け取り)
    - このエンドポイント以下はJWTトークンが必要になる
    - /restrictedこれを叩くと"ようこそ user@example.com さん！
"と返ってくる

    | 変数      |     説明     |
    |-----------|-----------|
    |JWTトークン     |  verifyで返ってきたやつ(24時間使用可能)  |

    - 使用例

    ```
    curl -H "Authorization: Bearer <JWTトークン>" http://localhost:8080/segon_pix_auth
    ```


