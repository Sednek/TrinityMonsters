#TrinityMonsters  
##_DB_:  
>***_Users_  
CREATE TABLE users  
(  
    id SERIAL PRIMARY key,  
    name VARCHAR(24) NOT NULL  
);  
***_Comments_  
CREATE TABLE "comments"   
(  
    id SERIAL PRIMARY key,  
    video_id INTEGER,  
    text VARCHAR(255) not null,  
    user_id INTEGER REFERENCES users (id),  
    date VARCHAR not null  
);  
***_Likes_  
CREATE TABLE likes  
(  
    id SERIAL PRIMARY key,  
    video_id INTEGER,  
    user_id INTEGER REFERENCES users (id),  
    date VARCHAR not null  
);  

