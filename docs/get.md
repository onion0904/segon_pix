# GET

- /get/user
    - 入力できる値(queryでの受け取り)
    - 与えられたユーザーの情報を返す。(passwordやEmail以外)

    | 変数      |     説明     |
    |-----------|-----------|
    |ID       |  ユーザーID   |

    - 使用例

    ```
    curl -X GET "http://localhost:8080/segon_pix/get/user?userID=1234"
    ```

- /get/user(token必要)
    - 入力できる値(queryでの受け取り)
    - このエンドポイント以下はJWTトークンが必要になる
    - Userを返す。emailとpasswordからUserを判断する

    | 変数      |     説明     |
    |-----------|-----------|
    |JWTトークン     |  verifyで返ってきたやつ(24時間使用可能)  |
    |ID       |  ユーザーID   |
    |  email   |  メールアドレス  |
    |  password   |  パスワード  |


    - 使用例

```bash
curl -X GET -H "Authorization: Bearer <JWTトークン>" \
"http://localhost:8080/segon_pix_auth/get/user?userID=1234&email=john@example.com&password=password"
``` 

- /get/image_detail
    - 入力できる値(queryでの受け取り)
    - 指定されたimageIDの画像の情報を返す。
    - PostedImageの全部を返す。

    | 変数      |     説明     |
    |-----------|-----------|
    |imageID       |  画像のID   |

    - 使用例

    ```
    curl -X GET "http://localhost:8080/segon_pix/get/image_detail?imageID=1234"


- /get/list
    - 入力できる値(queryでの受け取り)
    - Searchでは指定されたハッシュタグの部分一致を返す。
    - URLとimageIDの40個を返す。
    - 以下は全部自動で新しい順になります。
    
**Search, Recent, Like の組み合わせルール**
Search | Like | 取得順序 | 初期検索 | 次ページ取得条件
-- | -- | -- | -- | --
false | false | 新しい順（ID降順） | like_numとcurrentを省略 | current に前回ID、like_num = -1
true | true | 検索結果をいいね数順 | like_numとcurrentを省略 | current に前回ID、like_num に前回のいいね数
false | true | いいね数順 | like_numとcurrentを省略 | current に前回ID、like_num に前回のいいね数
true | false | 検索結果を新しい順（ID降順） | like_numとcurrentを省略 | current に前回ID、like_num = -1

**クエリの説明**
変数 | 説明     
-- | --
Hashtag | 検索したいワード   
Search | true or false   
Like | true or false    
like_num | 前回のいいね数   
current | 前回のid   
     
**使用例**

```
# ハッシュタグ検索 + いいね順 + 次ページ
curl "http://localhost:8080/segon_pix/get/list?Search=true&Hashtag=旅行&Like=true&like_num=7&current=123"

# 全体の新着順
curl "http://localhost:8080/segon_pix/get/list"

# 検索のみ
curl "http://localhost:8080/segon_pix/get/list?Search=true&Hashtag=猫"

# いいね順のみ
curl "http://localhost:8080/segon_pix/get/list?Like=true"
```
