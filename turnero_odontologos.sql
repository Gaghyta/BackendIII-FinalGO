CREATE DATABASE  IF NOT EXISTS `turnos-odontologia`;
USE `turnos-odontologia`;

DROP TABLE IF EXISTS `odontologos`;

CREATE TABLE `odontologos` (
  `odontologo_id` int NOT NULL AUTO_INCREMENT,
  `apellido_odontologo` varchar(45) NOT NULL,
  `nombre_odontologo` varchar(45) NOT NULL,
  `matricula` varchar(10) NOT NULL,
  PRIMARY KEY (`odontologo_id`),
  UNIQUE KEY `matricula_UNIQUE` (`matricula`)
); 

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


DROP TABLE IF EXISTS `pacientes`;

CREATE TABLE `pacientes` (
  `paciente_id` int NOT NULL AUTO_INCREMENT,
  `nombre_paciente` varchar(45) NOT NULL,
  `apellido_paciente` varchar(45) NOT NULL,
  `domicilio` varchar(45) NOT NULL,
  `dni` varchar(10) NOT NULL,
  `fecha_de_alta` varchar(10) NOT NULL,
  PRIMARY KEY (`paciente_id`),
  UNIQUE KEY `dni_UNIQUE` (`dni`)
) ;

INSERT INTO `pacientes` (`nombre_paciente`, `apellido_paciente`, `domicilio`, `dni`, `fecha_de_alta`) VALUES
  ('Juan', 'Pérez', 'Calle 123', '12345678', '2023-01-01'),
  ('María', 'González', 'Calle 456', '87654321', '2023-01-02'),
  ('Pedro', 'Fernández', 'Calle 789', '98765432', '2023-01-03'),
  ('Ana', 'García', 'Calle 101112', '45678912', '2023-01-04'),
  ('José', 'López', 'Calle 131415', '11223345', '2023-01-05'),
  ('María', 'Rodríguez', 'Calle 161718', '22334456', '2023-01-06'),
  ('Juan', 'Martínez', 'Calle 192021', '33445567', '2023-01-07'),
  ('Ana', 'Sánchez', 'Calle 222324', '44556678', '2023-01-08'),
  ('José', 'Romero', 'Calle 252627', '55667789', '2023-01-09'),
  ('María', 'Alvarez', 'Calle 282930', '66778890', '2023-01-10'),
  ('Juan', 'Castro', 'Calle 313233', '77889901', '2023-01-11'),
  ('Ana', 'Gutiérrez', 'Calle 343536', '88990012', '2023-01-12'),
  ('José', 'Ramírez', 'Calle 373839', '99001123', '2023-01-13'),
  ('María', 'Díaz', 'Calle 404142', '00112234', '2023-01-14'),
  ('Juan', 'Flores', 'Calle 434445', '0000000', '2023-01-15'),
  ('Ana', 'Herrera', 'Calle 464748', '1111111', '2023-01-16'),
  ('José', 'Vega', 'Calle 495051', '2222222', '2023-01-17'),
  ('María', 'Torres', 'Calle 525354', '3333333', '2023-01-18'),
  ('Juan', 'Flores', 'Calle 555657', '4444444', '2023-01-19');



DROP TABLE IF EXISTS `turnos`;

CREATE TABLE `turnos` (
  `turno_id` int NOT NULL AUTO_INCREMENT,
  `fecha_y_hora` varchar(45) NOT NULL,
  `descripcion` varchar(100) NOT NULL,
  `dentista_id_dentista` int NOT NULL,
  `paciente_id_paciente` int NOT NULL,
  PRIMARY KEY (`turno_id`)
); 

INSERT INTO `turnos`(`fecha_y_hora`,`descripcion`, `dentista_id_dentista`,`paciente_id_paciente`) VALUES 
('21/05/2022 11:22', 'control', 3, 4),
('17/03/2024 15:10', 'ortodoncia', 6, 8),
('15/12/2023 08:30', 'control de ortodoncia', 1, 4),
('07/07/2023 16:30', 'consulta', 2, 8);