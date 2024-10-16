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

- /verify

    - 入力できる値
    - 認証コードから判断してtokenを返す

    | 変数      |     説明     |
    |-----------|-----------|
    |Email     |  メールアドレス  |
    |password|パスワード|
    |code|認証コード|

    - 使用例

    ```
    curl -X POST http://localhost:8080/verify?email=user@example&password=pas&code=ABCDEF
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
    curl -H "Authorization: Bearer <JWTトークン>" http://localhost:8080/restricted
    ```


