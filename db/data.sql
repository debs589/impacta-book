insert into users(name, nickName, email, password)
values
    ("User 1", "user_1", "usuario1@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"),
    ("User 2", "user_2", "usuario2@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"),
    ("User 3", "user_3", "usuario3@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy");

insert into followers(user_id, follower_id)
values
    (1, 2),
    (3, 1),
    (1, 3);

insert into publications(title, content, author_id)
values
    ("User 1 publication", "Publication text for user 1", 1),
    ("User 1 publication", "Publication text for user 2", 2),
    ("User 2 publication", "Publication text for user 3", 3),
    ("User 4 publication", "Publication text for user 4", 4);