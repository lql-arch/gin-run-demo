create database douyin;

use douyin;

create table comment
(
    user_token  varchar(70)  not null,
    id          int auto_increment
        primary key,
    content     varchar(200) not null,
    create_date bigint       not null,
    video_id    int          not null,
    type        int          not null
);

create table message
(
    my_id      int                      not null,
    message    varchar(200) default ' ' not null,
    to_user_id int                      not null,
    create_at  bigint                   not null,
    id         int auto_increment
        primary key
);


create table user
(
    id               int auto_increment
        primary key,
    name             longtext         not null,
    follow_count     bigint           not null,
    follower_count   bigint           not null,
    token            varchar(100)     not null,
    background_image varchar(100)     not null,
    avatar           varchar(100)     not null,
    signature        varchar(100)     not null,
    total_favorited  bigint default 0 not null,
    work_count       bigint default 0 not null,
    favorite_count   bigint default 0 not null
);

create table user_video
(
    token          varchar(100) not null,
    video_id       int          not null,
    favorite_state int          not null,
    public_state   int          not null
);


create table videos
(
    id             int auto_increment
        primary key,
    author_id      bigint      null,
    play_url       longtext    null,
    cover_url      longtext    null,
    favorite_count bigint      null,
    comment_count  bigint      null,
    title          varchar(50) not null,
    create_at      bigint      not null,
    update_at      bigint      not null,
    constraint id
        unique (id)
);


create table relation
(
    my_id         int not null comment '关注者',
    other_user_id int not null comment '被关注者',
    state         int not null,
    constraint id
        unique (my_id, other_user_id)
);
