START TRANSACTION;
SET time_zone = "+00:00";

--
-- Creación de tablas
--

CREATE TABLE IF NOT EXISTS `users` (
    `user_id` int(11) NOT NULL,
    `email`  varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    `nombre` varchar(100) NOT NULL,
    `apellido` varchar(100) NOT NULL,
    `admin` boolean NOT NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `courses` (
    `course_id` int(11) NOT NULL,
    `nombre` varchar(100) NOT NULL,
    `profesor_id`  int(11) NOT NULL,
    `categoria` varchar(100) NOT NULL,
    `descripcion` VARCHAR(500) NOT NULL,
    `valoracion` decimal(2,1) NOT NULL,
    `duracion` INT NOT NULL,
    `requisitos` VARCHAR(500) NOT NULL,
    `url_image` VARCHAR(400) NOT NULL,
    `fecha_inicio` DATE NOT NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `inscripciones` (
    `inscripciones_id` int(11) NOT NULL,
    `user_id` int(11) NOT NULL,
    `course_id` int(11) NOT NULL,
    `fecha_inscripcion` DATE NOT NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `archivos`{
    `archivo_id` int(11) NOT NULL,
    `course_id` int (11) NOT NULL,
    `archivo` blob
    }

CREATE TABLE IF NOT EXISTS `comentarios`{
    `comentario_id` int(11) NOT NULL,
    `course_id`    int(11) NOT NULL,
    `user_id`       int(11) NOT NULL,
    `comentario`    VARCHAR(400),
    `valoracion`    VARCHAR(400)

    }
--
-- Indexado de tablas
--

ALTER TABLE `users`
    ADD PRIMARY KEY (`user_id`),
    ADD UNIQUE KEY `email` (`email`);

ALTER TABLE `courses`
    ADD PRIMARY KEY (`course_id`),
    ADD UNIQUE KEY `nombre` (`nombre`);
-- ADD KEY `profesor_id` (`user_id`);

ALTER TABLE `inscripciones`
    PRIMARY KEY (`inscripciones_id`),
    UNIQUE KEY `users_x_courses_id` (`user_id`, `course_id`),
    CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_course_id` FOREIGN KEY (`course_id`) REFERENCES `courses` (`course_id`) ON DELETE CASCADE ON UPDATE CASCADE
    }

ALTER TABLE `archivos`{
    ADD PRIMARY KEY (`archivo_id`),
    ADD KEY `course_id` (course_id)
    }

ALTER TABLE `comentarios`{
    ADD PRIMARY KEY (`comentario_id`),
    ADD KEY `course_id` (`course_id`),
    ADD KEY `user_id` (`user_id`)
    }
--
-- Autoincrementado de tablas
--

ALTER TABLE `users`
    MODIFY `user_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

ALTER TABLE `courses`
    MODIFY `course_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

ALTER TABLE `users_x_courses`
    MODIFY `users_x_courses_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Rellenado de tablas
--
INSERT INTO `users` (`user_id`,`email`, `password`, `nombre`, `apellido`, `admin`) VALUES
                                                                                       (1, 'sofiaolivetoo@gmail.com', 'ba77a5448b1208afe6effd5194c2a8b6', 'Sofia', 'Oliveto', false),
                                                                                       (2, 'juanlopez@gmail.com', 'f5737d25829e95b9c234b7fa06af8736', 'Juan', 'Lopez', true),
                                                                                       (3, 'constanzastrumia@gmail.com', 'febf04180a62e8710868cafd8741515f', 'Constanza', 'Strumia', false),
                                                                                       (4, 'margarita@gmail.com', '828fca74e9e1d7e55b76d46a304b5f55', 'Margarita', 'de Marcos', true),
                                                                                       (5, 'pedro@gmail.com','d3ce9efea6244baa7bf718f12dd0c331','Pedro','Juarez',false),
                                                                                       (6, 'josefinagonzalez@gmail.com','e577bc7b26b52afe6a33f02513b86b5c','Josefina', 'Gonzalez',false),
                                                                                       (7, 'ramiropaez@gmail.com','8e7a60d71791c1febdbc4998c963e87e','Ramiro', 'Paez', false),
                                                                                       (8, 'gustavojacobo@gmail.com','0805446e686aa72d45f9583f2d6cedef','Gustavo', 'Jacobo', true),
                                                                                       (9, 'matigarcia@gmail.com','0596f701227172915b2862b95b4c2e1a', 'Matias', 'Portillo', false),
                                                                                       (10, 'juliomansilla@gmail.com','16880e98af692b72ce3ba695654ee306', 'Julio', 'Mansilla', false),
                                                                                       (11, 'santiportillo@gmail.com','d5116c2a9607b0ea07d425506f610467', 'Santiago', 'Portillo', false),
                                                                                       (12, 'nicolasfigueroa@gmail.com','305735d035e7f7381d64d179126ff6d9', 'Nicolas', 'Figueroa', false),
                                                                                       (13, 'agostinacisneros@gmail.com','0a9827114b460ae8f2f96c5e8893c90d', 'Agostina', 'Cisneros', false),
                                                                                       (14, 'luciabernardi@gmail.com','b9ce57d6d6c2d6fda25d80da5a00a7d1', 'Lucia', 'Bernardi', false),
                                                                                       (15, 'pauladominguez@gmail.com','ca46a79286419c05172ca7b010a59d3c', 'Paula', 'Dominguez', false),
                                                                                       (16, 'luciovelarde@gmail.com','f4a1f4d408e436f3c294bf2ae346b3d7', 'Lucio', 'Velarde', false),
                                                                                       (17, 'rodolfoperez@gmail.com','4a2b3910b547e5212914378adaf76aac', 'Rodolfo', 'Perez', true),
                                                                                       (18, 'sebastiancolidio@gmail.com','5a7c2cf0d17f9d32c87de8efb8e689d6', 'Sebastian', 'Colidio', true),
                                                                                       (19, 'lucasbeltran@gmail.com','6d16ba70238c92a03ac04c7c86eb79e7', 'Lucas', 'Beltran', true),
                                                                                       (20, 'chilenodiaz@gmail.com','4494d10dc9752cba4083ce2cf8983d2c', 'Paulo', 'Diaz', false);

INSERT INTO `courses` (`course_id`,`nombre`, `profesor_id`, `categoria`, `descripcion`, `valoracion`, `duracion`, `requisitos`, `url_image`, `fecha_inicio`) VALUES
                                                                                                                                                                 (1, 'Golang para principiantes' , 8, 'Programacion', 'Curso basico de programacion de golang', 9000,'Bajo', 'https://www.freecodecamp.org/news/content/images/2021/10/golang.png', '2020-06-09'),
                                                                                                                                                                 (2, 'Bartender ricos tragos' , 2, 'Cocteleria', 'Aprende a hacer los mejores tragos', 45,'Bajo', 'https://s3-eu-west-1.amazonaws.com/verema/images/valoraciones/0011/0991/coctel.jpg?1353314930', '2024-06-03'),
                                                                                                                                                                 (3, 'Peluqueria intensiva', 4, 'Belleza', 'Aprende a hacer desde simples trenzas hasta peinados elaborados', 80, 'Bajo', 'https://joseppons.com/formacion/wp-content/uploads/2019/06/abrir-tu-salon-peluqueria.jpg', '2024-07-10'),
                                                                                                                                                                 (4, 'Golang avanzado', 8, 'Programacion', 'Curso avanzado del lenguaje golang', 600, 'Alto', 'https://www.freecodecamp.org/news/content/images/2021/10/golang.png', '2024-07-10'),
                                                                                                                                                                 (5, 'Pasteleria con Lucas', 19, 'Cocina', 'Curso basico y divertido de pasteleria', 150, 'Bajo', 'https://www.ceac.es/sites/default/files/2019-08/ceac-27-03-2019-reposteria.jpg.webp', '2024-08-12'),
                                                                                                                                                                 (6, 'Bossa Nova en guitarra', 18, 'Musica', 'Toca canciones del genero bossa nova en tu guitarra', 68, 'Intermedio', 'https://img.freepik.com/fotos-premium/foto-persona-irreconocible-tocando-guitarra_762026-81126.jpg', '2024-03-10'),
                                                                                                                                                                 (7, 'Maquillaje artistico, brillos y fantasia', 17, 'Belleza', 'Crea tu propio look con Rodolfo y sus esplendidos maquillajes artisticos', 72, 'Bajo', 'https://st2.depositphotos.com/6940744/12394/i/450/depositphotos_123947758-stock-photo-years-old-boy-with-blue.jpg', '2024-09-09'),
                                                                                                                                                                 (8, 'Coffee coctels', 19, 'Cocteleria', 'Aprende a preparar distintos coctels usando cafe', 70, 'Intermedio', 'https://apasionados-por-el-cafe.s3.amazonaws.com/wp-content/uploads/2020/12/shutterstock_1035562495-min_opt.jpg', '2024-11-01'),
                                                                                                                                                                 (9, 'C++ primeros pasos', 8, 'Programacion', 'Aprende a dar tus primeros pasos la programacion con el mejor profe', 120, 'Bajo', 'https://blog.facialix.com/wp-content/uploads/2023/11/curso-gratis-cpp.jpg', '2024-03-02');


COMMIT;