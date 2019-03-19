Cara menjalankan

1. $ docker-compose up
2. Tunggu beberapa saat
3. $ docker inspect plat_adm_db | Grep "IPAddress"
4. Setelah dapat IP, itu IP untuk database bisa dicheck via adminer localhost:8080

incase gagal
1. $ docker-compose down
2. hapus folder vendor
3. $ docker-compose up

lihat database:
1. Buka localhost:8080
    - System : PostgreSQL
    - Server : database
    - Username : postgres
    - Password : asdasd
