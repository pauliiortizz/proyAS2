CREATE TABLE IF NOT EXISTS `users` (
                                       `user_id` int(11) NOT NULL,
    `email`  varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    `nombre` varchar(100) NOT NULL,
    `apellido` varchar(100) NOT NULL,
    `admin` boolean NOT NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

ALTER TABLE `users`
    ADD PRIMARY KEY (`user_id`),
    ADD UNIQUE KEY `email` (`email`);

ALTER TABLE `users`
    MODIFY `user_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

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
