# GET

- /get/user
    - 入力できる値
    - 与えられたユーザーの情報を返す。

    | 変数      |     説明     |
    |-----------|-----------|
    |ID       |  ユーザーID   |

    - 使用例

    ```
    curl -X GET "http://localhost:8080/segon_pix/get/user?ID=1234"
    ```
- /get/list/image
    - 入力できる値
    - 指定されたハッシュタグの部分一致を返す。

    | 変数      |     説明     |
    |-----------|-----------|
    |Hashtag       |  検索したいワード   |

    - 使用例

    ```
    curl -X GET "http://localhost:8080/segon_pix/get/list/image?Hashtag=旅行"
    ```
