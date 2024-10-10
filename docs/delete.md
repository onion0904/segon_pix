# DELETE

- /delete/user
    - 入力できる値(queryでの受け取り)
    - ユーザー情報を削除する。

    | 変数      |    説明     |
    |-----------|-----------|
    |ID       |  ユーザーID   |

    - 使用例

    ```
    curl -X DELETE "http://localhost:8080/segon_pix/delete/user?ID=1234"
    ```
- /delete/image
    - 入力できる値(queryでの受け取り)
    - 画像をDBとGCSから削除する。

    | 変数      |     説明     |
    |-----------|-----------|
    |ID       |   画像ID  |

    - 使用例

    ```
    curl -X DELETE "http://localhost:8080/segon_pix/delete/image?ID=1234"
    ```
- /delete/like
    - 入力できる値(queryでの受け取り)
    - いいねを取り消す。ユーザー情報のいいね欄からも消す。

    | 変数      |     説明     |
    |-----------|-----------|
    |userID     |  ユーザーID  |
    |imageID    |  画像ID  |

    - 使用例

    ```
    curl -X DELETE "http://localhost:8080/segon_pix/delete/like?userID=1234&imageID=5678"
    ```
- /delete/comment
    - 入力できる値(queryでの受け取り)
    - コメントを消す

    | 変数      |     説明     |
    |-----------|-----------|
    |commentID       |  コメントのID  |
    |imageID    |  画像のID  |
    - 使用例

    ```
    curl -X DELETE "http://localhost:8080/segon_pix/delete/comment?commentID=5678&imageID=1234"
    ```
