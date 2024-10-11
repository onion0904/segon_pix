# auth

- /login

    - 入力できる値(JSONでの受け取り)
    - メアドとパスワードを入力して認証コードをメアド先に送信する

    | 変数      |     説明     |
    |-----------|-----------|
    |Email     |  メールアドレス  |
    |password|パスワード|

    - 使用例

    ```
    curl -X POST -F "email=user@example.com" -F "password=password123" http://localhost:8080/login
    ```

- /verify

    - 入力できる値(JSONでの受け取り)
    - メアドとパスワードを入力して認証コードをメアド先に送信する

    | 変数      |     説明     |
    |-----------|-----------|
    |Email     |  メールアドレス  |
    |code|認証コード|

    - 使用例

    ```
    curl -X POST -F "email=user@example.com" -F "code=ABCDEF" http://localhost:8080/verify
    ```

- /restricted

    - 入力できる値(JSONでの受け取り)
    - このエンドポイント以下はJWTトークンが必要になる
    - /restrictedこれを叩くと"ようこそ user@example.com さん！
"と返ってくる

    | 変数      |     説明     |
    |-----------|-----------|
    |JWTトークン     |  verifyで返ってきたやつ(24時間使用可能)  |

    - 使用例

    ```
    curl -H "Authorization: Bearer <JWTトークン>" http://localhost:8080/restricted
    ```