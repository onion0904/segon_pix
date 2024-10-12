# auth（作り途中）

- /login->signupに変更

    - 入力できる値
    - メアドとパスワードを入力して認証コードをメアド先に送信する

    | 変数      |     説明     |
    |-----------|-----------|
    |Email     |  メールアドレス  |
    |password|パスワード|

    - 使用例

    ```
    curl -X POST -F "email=user@example.com" -F "password=password123" http://localhost:8080/signup
    ```

- /verify

    - 入力できる値
    - 認証コードから判断してtokenを返す(JWTにpasswordを追加する)

    | 変数      |     説明     |
    |-----------|-----------|
    |Email     |  メールアドレス  |
    |password|パスワード|
    |code|認証コード|

    - 使用例

    ```
    curl -X POST -F "email=user@example.com" -F "password=password123" -F "code=ABCDEF" http://localhost:8080/verify
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
    curl -X POST -F "email=user@example.com" -F "password=password123" http://localhost:8080/login
    ```

- /restricted
    - 入力できる値(JSONでの受け取り)
    - このエンドポイント以下はJWTトークンが必要になる
    - /restrictedこれを叩くと"ようこそ user@example.com さん！
"と返ってくる->Userを返すように変更する。emailとpasswordからUserを判断する

    | 変数      |     説明     |
    |-----------|-----------|
    |JWTトークン     |  verifyで返ってきたやつ(24時間使用可能)  |

    - 使用例

    ```
    curl -H "Authorization: Bearer <JWTトークン>" http://localhost:8080/restricted
    ```