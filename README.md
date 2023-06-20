# K-Style Task

## Task Secara Garis Besar

1. CRUD Member
2. READ Product + Data Relasi
3. Like dan Unlike
4. Tech Stack Go (Gorm dan Echo)

## List Endpoint
    1. Member :
    - /v1/members/:id
    Gunakan Method Post untuk Menginputkan data member,Get untuk membaca data member + tambahkan parameter id yang valid untuk membaca data member by Id,PUT + tambahkan parameter id untuk mengedit data member,Delete + tambahkan parameter id untuk menghapus member.
    2. Product :
    - /v1/products/:id
     Gunakan Method Post untuk Menginputkan data product,Get untuk membaca data product + tambahkan parameter id untuk membaca data product by Id,PUT + tambahkan parameter id untuk mengedit data product,Delete + tambahkan parameter id untuk menghapus product.
    3. Review :
    - /v1/reviews
    Gunakan Method Post untuk Menginputkan data Review, id member yang tidak valid atau telah dihapus (deleted_at not null) tidak dapat melakukan create review, begitu juga dengan id product.
    4. Like
    - /v1/likes
    Endpoint ini bisa digunakan untuk meng-like (menyukai) suatu review sekaligus mengunlike (mengcancel like) suatu review, id member yang tidak valid atau telah dihapus (deleted_at not null) tidak dapat melakukan create review, begitu juga dengan id review.
