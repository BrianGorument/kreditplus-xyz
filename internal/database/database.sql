-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 19, 2024 at 10:52 AM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `xyz_multifinance`
--

-- --------------------------------------------------------

--
-- Table structure for table `customers`
--

CREATE TABLE `customers` (
  `nik` varchar(20) NOT NULL,
  `full_name` varchar(100) DEFAULT NULL,
  `legal_name` varchar(100) DEFAULT NULL,
  `birth_place` varchar(100) DEFAULT NULL,
  `birth_date` date DEFAULT NULL,
  `salary` int(11) DEFAULT NULL,
  `photo_ktp` text DEFAULT NULL,
  `photo_selfie` text DEFAULT NULL,
  `one_month_limit` int(11) DEFAULT NULL,
  `two_month_limit` int(11) DEFAULT NULL,
  `three_month_limit` int(11) DEFAULT NULL,
  `four_month_limit` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `customers`
--

INSERT INTO `customers` (`nik`, `full_name`, `legal_name`, `birth_place`, `birth_date`, `salary`, `photo_ktp`, `photo_selfie`, `one_month_limit`, `two_month_limit`, `three_month_limit`, `four_month_limit`) VALUES
('1234567890', 'Budi Santoso', 'Budi', 'Jakarta', '1980-01-01', 5000000, 'http://example.com/ktp_budi.jpg', 'http://example.com/selfie_budi.jpg', 100000, 200000, 500000, 700000),
('1234567891', 'Jason Santoso', 'Jason', 'Jakarta', '1980-01-01', 8000000, 'http://example.com/ktp_json.jpg', 'http://example.com/selfie_json.jpg', 100000, 200000, 500000, 700000);

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `contract_number` varchar(20) NOT NULL,
  `nik` varchar(20) DEFAULT NULL,
  `otr` int(11) DEFAULT NULL,
  `admin_fee` int(11) DEFAULT NULL,
  `installment` int(11) DEFAULT NULL,
  `interest` int(11) DEFAULT NULL,
  `asset_name` varchar(100) DEFAULT NULL,
  `loan_date` date DEFAULT NULL,
  `due_date` date DEFAULT NULL,
  `is_paid` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`contract_number`, `nik`, `otr`, `admin_fee`, `installment`, `interest`, `asset_name`, `loan_date`, `due_date`, `is_paid`) VALUES
('CONTRACT004', '1234567891', 1800000, 50000, 4, 50000, 'Smart watch', '2024-05-18', '2024-06-18', 0),
('TRX0001', '1234567890', 1995000, 50000, 2, 50000, 'HP Xiaomi', '2024-05-20', '2024-06-20', 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `customers`
--
ALTER TABLE `customers`
  ADD PRIMARY KEY (`nik`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`contract_number`),
  ADD KEY `nik` (`nik`);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`nik`) REFERENCES `customers` (`nik`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
