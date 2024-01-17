INSERT INTO usuarios (nome, nick, email, senha)
values
('Usuario 1', 'usuario1', 'usuario1@email.com', '$2a$10$3sr1vPYAVxBgCq8Kk52QwudS6VkMBdGrPS70HoEx0HHUUvEO2fbx2'),
('Usuario 2', 'usuario2', 'usuario2@email.com', '$2a$10$3sr1vPYAVxBgCq8Kk52QwudS6VkMBdGrPS70HoEx0HHUUvEO2fbx2'),
('Usuario 3', 'usuario3', 'usuario3@email.com', '$2a$10$3sr1vPYAVxBgCq8Kk52QwudS6VkMBdGrPS70HoEx0HHUUvEO2fbx2'),
('Usuario 4', 'usuario4', 'usuario4@email.com', '$2a$10$3sr1vPYAVxBgCq8Kk52QwudS6VkMBdGrPS70HoEx0HHUUvEO2fbx2'),
('Usuario 5', 'usuario5', 'usuario5@email.com', '$2a$10$3sr1vPYAVxBgCq8Kk52QwudS6VkMBdGrPS70HoEx0HHUUvEO2fbx2');

INSERT INTO seguidores (usuario_id, seguidor_id)
VALUES (1, 2),(2, 3),(3, 4),(4, 5),(5, 1);

INSERT INTO publicacoes (titulo, conteudo, autor_id)
VALUES("Publicação do Autor 1", "Essa é a publicacao do usuario 1", 1),
("Publicação do Autor 2", "Essa é a publicacao do usuario 2", 2),
("Publicação do Autor 3", "Essa é a publicacao do usuario 3", 3);