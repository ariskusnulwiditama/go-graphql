make database
create database graphdb;

use graphdb;

CREATE TABLE `level_user` (
 `id_level` int(2) NOT NULL AUTO_INCREMENT,
 `nama` varchar(100) NOT NULL,
 PRIMARY KEY (`id_level`)
);

CREATE TABLE `users` (
 `email` varchar(100) NOT NULL,
 `nama` varchar(50) NOT NULL,
 `no_handphone` varchar(15) DEFAULT NULL,
 `alamat` text DEFAULT NULL,
 `ktp` varchar(255) DEFAULT NULL,
 `id_level` int(2) NOT NULL,
 PRIMARY KEY (`email`) USING BTREE
);

INSERT INTO `level_user` (`id_level`, `nama`) VALUES
	(1, 'Super Admin'),
	(2, 'Admin');

INSERT INTO `users` (`email`, `nama`, `no_handphone`, `alamat`, `ktp`, `id_level`) VALUES
	('fulan@gmail.com', 'Ahmad Zainuddin', '0853xxxxx', 'Jalan Kemana', '', 1),
	('siapa@gmail.com', 'Zubayr', '08xxxx', 'Jalan - ', 'ktp.jpg', 2),
	('saya@gmail.com', Zakaria, '08xxxx', 'Jalan - ', 'ktp.jpg', 2);

run application
go run server.go

localhost:1323/graphql

run query
query{
    users{
        email
        nama
        no_handphone
    }
    level(id_level:1){
        id_level
        nama
    }
}