# PUT

- /update/comment
    - 入力できる値(queryでの受け取り)
    - コメントの内容を更新する。

    | 変数      |     説明     |
    |-----------|-----------|
    |commentID       |  コメントのID  |
    |imageID    |  画像のID  |
    |newContent       |  更新後のコメントの内容  |


    - 使用例

    ```
   curl -X PUT "http://localhost:8080/segon_pix/update/comment?commentID=5678&imageID=1234&newContent=更新されたコメント内容"
    ```
