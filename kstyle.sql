-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 20 Jun 2023 pada 18.58
-- Versi server: 10.4.27-MariaDB
-- Versi PHP: 8.2.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `kstyle`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `likes`
--

CREATE TABLE `likes` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `id_review` bigint(20) UNSIGNED DEFAULT NULL,
  `id_member` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `likes`
--

INSERT INTO `likes` (`id`, `created_at`, `updated_at`, `deleted_at`, `id_review`, `id_member`) VALUES
(2, '2023-06-20 16:54:24.191', '2023-06-20 16:54:24.191', NULL, 1, 1),
(3, '2023-06-20 16:54:29.843', '2023-06-20 16:54:29.843', NULL, 1, 2),
(4, '2023-06-20 16:54:36.787', '2023-06-20 16:54:36.787', NULL, 1, 3),
(5, '2023-06-20 16:55:18.812', '2023-06-20 16:55:18.812', NULL, 2, 2),
(6, '2023-06-20 16:55:24.033', '2023-06-20 16:55:24.033', NULL, 3, 2),
(7, '2023-06-20 16:55:28.815', '2023-06-20 16:55:28.815', NULL, 3, 3),
(8, '2023-06-20 16:55:33.276', '2023-06-20 16:55:33.276', NULL, 3, 4);

-- --------------------------------------------------------

--
-- Struktur dari tabel `members`
--

CREATE TABLE `members` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` longtext DEFAULT NULL,
  `gender` varchar(255) DEFAULT NULL,
  `skintype` varchar(255) DEFAULT NULL,
  `skincolor` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `members`
--

INSERT INTO `members` (`id`, `created_at`, `updated_at`, `deleted_at`, `username`, `gender`, `skintype`, `skincolor`) VALUES
(1, '2023-06-20 16:42:08.006', '2023-06-20 16:42:08.006', NULL, 'aryadevaraj', 'laki-laki', 'normal to dry', 'white'),
(2, '2023-06-20 16:42:48.090', '2023-06-20 16:42:48.090', NULL, 'vinkaannisa', 'perempuan', 'oily combination', 'black'),
(3, '2023-06-20 16:43:06.516', '2023-06-20 16:43:06.516', NULL, 'vinkasaja', 'perempuan', 'oily combination', 'pale'),
(4, '2023-06-20 16:43:21.777', '2023-06-20 16:43:21.777', NULL, 'vinkatok', 'perempuan', 'oily combination', 'medium'),
(5, '2023-06-20 16:43:34.241', '2023-06-20 16:46:01.837', '2023-06-20 16:46:40.761', 'testong', 'perempuan', 'oily combination', 'medium');

-- --------------------------------------------------------

--
-- Struktur dari tabel `products`
--

CREATE TABLE `products` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name_product` varchar(255) DEFAULT NULL,
  `price` double DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `products`
--

INSERT INTO `products` (`id`, `created_at`, `updated_at`, `deleted_at`, `name_product`, `price`) VALUES
(1, '2023-06-20 16:47:35.377', '2023-06-20 16:47:35.377', NULL, 'cosrx', 60000),
(2, '2023-06-20 16:47:58.492', '2023-06-20 16:47:58.492', NULL, 'unnis pick', 75000),
(3, '2023-06-20 16:48:11.674', '2023-06-20 16:48:11.674', NULL, 'hadalabo', 50000),
(4, '2023-06-20 16:48:27.484', '2023-06-20 16:48:27.484', NULL, 'skintific', 45000),
(5, '2023-06-20 16:48:37.610', '2023-06-20 16:48:37.610', NULL, 'avoskin', 30000),
(6, '2023-06-20 16:48:46.899', '2023-06-20 16:49:49.613', '2023-06-20 16:50:06.122', 'biore', 25000);

-- --------------------------------------------------------

--
-- Struktur dari tabel `reviews`
--

CREATE TABLE `reviews` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `id_product` bigint(20) UNSIGNED DEFAULT NULL,
  `id_member` bigint(20) UNSIGNED DEFAULT NULL,
  `desc_review` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `reviews`
--

INSERT INTO `reviews` (`id`, `created_at`, `updated_at`, `deleted_at`, `id_product`, `id_member`, `desc_review`) VALUES
(1, '2023-06-20 16:50:33.943', '2023-06-20 16:50:33.943', NULL, 1, 1, 'sangat cocok untuk kulit saya'),
(2, '2023-06-20 16:51:09.054', '2023-06-20 16:51:09.054', NULL, 1, 2, 'agak kental'),
(3, '2023-06-20 16:51:22.665', '2023-06-20 16:51:22.665', NULL, 2, 2, 'bagus sekali'),
(4, '2023-06-20 16:51:46.200', '2023-06-20 16:51:46.200', NULL, 3, 3, 'kurang cocok');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `likes`
--
ALTER TABLE `likes`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_member_review` (`id_review`,`id_member`),
  ADD KEY `idx_likes_deleted_at` (`deleted_at`),
  ADD KEY `fk_members_likes` (`id_member`);

--
-- Indeks untuk tabel `members`
--
ALTER TABLE `members`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_members_deleted_at` (`deleted_at`);

--
-- Indeks untuk tabel `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_products_deleted_at` (`deleted_at`);

--
-- Indeks untuk tabel `reviews`
--
ALTER TABLE `reviews`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_reviews_deleted_at` (`deleted_at`),
  ADD KEY `fk_members_reviews` (`id_member`),
  ADD KEY `fk_products_reviews` (`id_product`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `likes`
--
ALTER TABLE `likes`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT untuk tabel `members`
--
ALTER TABLE `members`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT untuk tabel `products`
--
ALTER TABLE `products`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT untuk tabel `reviews`
--
ALTER TABLE `reviews`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `likes`
--
ALTER TABLE `likes`
  ADD CONSTRAINT `fk_members_likes` FOREIGN KEY (`id_member`) REFERENCES `members` (`id`),
  ADD CONSTRAINT `fk_reviews_likes` FOREIGN KEY (`id_review`) REFERENCES `reviews` (`id`);

--
-- Ketidakleluasaan untuk tabel `reviews`
--
ALTER TABLE `reviews`
  ADD CONSTRAINT `fk_members_reviews` FOREIGN KEY (`id_member`) REFERENCES `members` (`id`),
  ADD CONSTRAINT `fk_products_reviews` FOREIGN KEY (`id_product`) REFERENCES `products` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
