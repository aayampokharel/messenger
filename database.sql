  -- drop database messenger_db;
CREATE DATABASE if not exists messenger_db;
use messenger_db;

create table if not exists loginCredential_table(
registration_id int unique not null auto_increment ,
email VARCHAR(100) primary key,
display_name VARCHAR(50) NOT NULL,
passwords varchar(40) not null
-- profile is to be inserted in future. 
);
INSERT INTO chat_connections (room_id, sender_id, receiver_id) VALUES
(1234, 1, 2),
(2345, 3, 4);

CREATE table if not exists chat_connections(
room_id bigint primary key,
 sender_id int not null,
receiver_id int not null,
-- latest_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
 foreign key (sender_id) references loginCredential_table(registration_id),
 foreign key (receiver_id) references loginCredential_table(registration_id)

    
    -- !! THIS HAS  TO BE CHANGDE TO TAKE 
    -- INPUT FROM FLUTTER WHEN DATES AND TIME IS PASSED .!!
);

-- truncate table private_chats_table ;
-- INSERT INTO private_chats_table (private_room_id, receiver_id, chat, latest_time) VALUES
-- (1234, 1, 'Hey, how are you?', '2024-07-05 10:00:00'),
-- (1234, 2, 'I\'m good, thanks! How about you?', '2024-07-05 10:01:00'),
-- (1234, 1, 'Doing well, just working on some projects.', '2024-07-05 10:02:00'),
-- (1234, 2, 'That sounds interesting. What kind of projects?', '2024-07-05 10:03:00'),
-- (1234, 1, 'Building a new app. It\'s a messenger clone.', '2024-07-05 10:04:00'),
-- (1234, 2, 'Nice! What technologies are you using?', '2024-07-05 10:05:00'),
-- (1234, 1, 'Flutter for the frontend and Go for the backend.', '2024-07-05 10:06:00'),
-- (1234, 2, 'Cool combination. Need any help with it?', '2024-07-05 10:07:00'),
-- (1234, 1, 'Maybe later. Thanks for offering!', '2024-07-05 10:08:00'),
-- (1234, 2, 'No problem. Just let me know.', '2024-07-05 10:09:00'),
-- (1234, 1, 'Will do. Talk to you later!', '2024-07-05 10:10:00'),
-- (1234, 2, 'Sure, catch you later.', '2024-07-05 10:11:00'),
-- (2345, 3, 'Hey, did you watch the game last night?', '2024-07-05 11:00:00'),
-- (2345, 4, 'Yes, it was amazing!', '2024-07-05 11:01:00'),
-- (2345, 3, 'That last goal was incredible!', '2024-07-05 11:02:00'),
-- (2345, 4, 'Absolutely, can\'t wait for the next match.', '2024-07-05 11:03:00');
-- select * from private_chats_table;

create table user_search_history(
current_user_id int not null, 
other_user_id int not null ,
foreign key(current_user_id) references loginCredential_table(registration_id),
foreign key(other_user_id) references loginCredential_table(registration_id)
);


select * from loginCredential_table; 
create table if not exists private_chats_table(
private_room_id bigint not null,
receiver_id int not null,
chat longtext not null ,
latest_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
foreign key (private_room_id) references chat_connections(room_id),
foreign key (receiver_id) references loginCredential_table(registration_id)

);

--   select * from loginCredential_table order by registration_id asc ;
   -- select * from chat_connections;
 -- select * from private_chats_table ;



