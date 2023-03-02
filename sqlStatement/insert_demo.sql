insert into user(name, follow_count, follower_count, is_follow, token)
    value ('TestUser',0,0,false,1);

insert into comment(user_token, content, create_date, video_id)
    value ('TestUserTestUser','Test comment',now(),1);

insert into videos(author_id, play_url, cover_url, favorite_count, comment_count, is_favorite)
    value (1,'https://www.w3schools.com/html/movie.mp4','https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg',0,0,false);

