# DELETE

- /delete/user(token必要)
    - 入力できる値(queryでの受け取り)
    - ユーザー情報を削除する。

    | 変数      |    説明     |
    |-----------|-----------|
    |userID       |  ユーザーID   |

    - 使用例

```bash
curl -X DELETE -H "Authorization: Bearer <JWTトークン>" \
"http://localhost:8080/segon_pix_auth/delete/user?userID=1234"
```

- /delete/image(token必要)
    - 入力できる値(queryでの受け取り)
    - 画像をDBとGCSから削除する。

    | 変数      |     説明     |
    |-----------|-----------|
    |imageID       |   画像ID  |
    |userID       |   ユーザーのID  |


    - 使用例

    ```
    curl -X DELETE -H "Authorization: Bearer <JWTトークン>" \"http://localhost:8080/segon_pix_auth/delete/image?userID=1234&imageID=1234"
    ```
- /delete/like(token必要)
    - 入力できる値(queryでの受け取り)
    - いいねを取り消す。ユーザー情報のいいね欄からも消す。

    | 変数      |     説明     |
    |-----------|-----------|
    |userID     |  ユーザーID  |
    |imageID    |  画像ID  |

    - 使用例

    ```
    curl -X DELETE -H "Authorization: Bearer <JWTトークン>" \"http://localhost:8080/segon_pix_auth/delete/like?userID=1234&imageID=5678"
    ```
- /delete/comment(token必要)
    - 入力できる値(queryでの受け取り)
    - コメントを消す

    | 変数      |     説明     |
    |-----------|-----------|
    |userID       |  userのID  |
    |commentID       |  コメントのID  |
    - 使用例

    ```
    curl -X DELETE -H "Authorization: Bearer <JWTトークン>" \"http://localhost:8080/segon_pix_auth/delete/comment?userID=1234&commentID=5678"
    ```

- /delete/follow(token必要)
    - 入力できる値(queryでの受け取り)
    - followingIDのユーザーのFollows(自分がフォローしてるユーザー)からfollowedIDのユーザーを削除する。

    | 変数      |     説明     |
    |-----------|-----------|
    |followingID     |  フォローしてるユーザーID  |
    |followedID    |  フォローされてるユーザーID  |

    - 使用例

    ```
    curl -X DELETE -H "Authorization: Bearer <JWTトークン>" \"http://localhost:8080/segon_pix_auth/delete/follow?followingID=1234&followedID=5678"
    ```
