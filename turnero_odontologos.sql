-- Delete user if already exists
DROP USER IF EXISTS 'user1'@'localhost';

-- Create user with all privileges
CREATE USER 'user1'@'localhost' IDENTIFIED BY 'secret_password';
GRANT ALL PRIVILEGES ON *.* TO 'user1'@'localhost';
-- 


CREATE DATABASE  IF NOT EXISTS `turnos-odontologia` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `turnos-odontologia`;
-- MySQL dump 10.13  Distrib 8.0.36, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: turnero_odontologos
-- ------------------------------------------------------
-- Server version	8.0.36

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `odontologos`
--

DROP TABLE IF EXISTS `odontologos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `odontologos` (
  `odontologo_id` int NOT NULL AUTO_INCREMENT,
  `apellido_odontologo` varchar(45) NOT NULL,
  `nombre_odontologo` varchar(45) NOT NULL,
  `matricula` varchar(10) NOT NULL,
  PRIMARY KEY (`odontologo_id`),
  UNIQUE KEY `matricula_UNIQUE` (`matricula`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `odontologos`
--

LOCK TABLES `odontologos` WRITE;
/*!40000 ALTER TABLE `odontologos` DISABLE KEYS */;
INSERT INTO `odontologos` (`apellido_odontologo`, `nombre_odontologo`, `matricula`) VALUES
  ('Pérez', 'Juan', '123456'),
  ('González', 'María', '654321'),
  ('Fernández', 'Pedro', '987654'),
  ('García', 'Ana', '456789'),
  ('López', 'José', '112233'),
  ('Rodríguez', 'María', '223344'),
  ('Martínez', 'Juan', '334455'),
  ('Sánchez', 'Ana', '445566'),
  ('Romero', 'José', '556677'),
  ('Alvarez', 'María', '667788');
/*!40000 ALTER TABLE `odontologos` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pacientes`
--

DROP TABLE IF EXISTS `pacientes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pacientes` (
  `paciente_id` int NOT NULL AUTO_INCREMENT,
  `nombre_paciente` varchar(45) NOT NULL,
  `apellido_paciente` varchar(45) NOT NULL,
  `domicilio` varchar(45) NOT NULL,
  `dni` varchar(10) NOT NULL,
  `fecha_de_alta` varchar(10) NOT NULL,
  PRIMARY KEY (`paciente_id`),
  UNIQUE KEY `dni_UNIQUE` (`dni`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pacientes`
--

LOCK TABLES `pacientes` WRITE;
/*!40000 ALTER TABLE `pacientes` DISABLE KEYS */;
INSERT INTO `pacientes` (`nombre_paciente`, `apellido_paciente`, `domicilio`, `dni`, `fecha_de_alta`) VALUES
  ('Juan', 'Pérez', 'Calle 123', '12345778', '2023-01-01'),
  ('María', 'González', 'Calle 456', '87614321', '2023-01-02'),
  ('Pedro', 'Fernández', 'Calle 789', '98715432', '2023-01-03'),
  ('Ana', 'García', 'Calle 101112', '45671912', '2023-01-04'),
  ('José', 'López', 'Calle 131415', '11223345', '2023-01-05'),
  ('María', 'Rodríguez', 'Calle 161718', '22234456', '2023-01-06'),
  ('Juan', 'Martínez', 'Calle 192021', '33405567', '2023-01-07'),
  ('Ana', 'Sánchez', 'Calle 222324', '44554678', '2023-01-08'),
  ('José', 'Romero', 'Calle 252627', '55667780', '2023-01-09'),
  ('María', 'Alvarez', 'Calle 282930', '66778890', '2023-01-10'),
  ('Juan', 'Castro', 'Calle 313233', '77889901', '2023-01-11'),
  ('Ana', 'Gutiérrez', 'Calle 343536', '88990012', '2023-01-12'),
  ('José', 'Ramírez', 'Calle 373839', '99001123', '2023-01-13'),
  ('María', 'Díaz', 'Calle 404142', '00112234', '2023-01-14'),
  ('Juan', 'Flores', 'Calle 434445', '11220345', '2023-01-15'),
  ('Ana', 'Herrera', 'Calle 464748', '22334456', '2023-01-16'),
  ('José', 'Vega', 'Calle 495051', '33445567', '2023-01-17'),
  ('María', 'Torres', 'Calle 525354', '44556678', '2023-01-18'),
  ('Juan', 'Flores', 'Calle 555657', '55667789', '2023-01-19');
/*!40000 ALTER TABLE `pacientes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `turnos`
--

DROP TABLE IF EXISTS `turnos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `turnos` (
  `turno_id` int NOT NULL AUTO_INCREMENT,
  `fecha_y_hora` varchar(45) NOT NULL,
  `descripcion` varchar(100) NOT NULL,
  `dentista_id_dentista` int NOT NULL,
  `paciente_id_paciente` int NOT NULL,
  PRIMARY KEY (`turno_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `turnos`
--

LOCK TABLES `turnos` WRITE;
/*!40000 ALTER TABLE `turno` DISABLE KEYS */;
/*!40000 ALTER TABLE `turno` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-03-21 10:23:50
