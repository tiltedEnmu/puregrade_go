create table users (
    id serial primary key,
    email varchar(255) not null unique,
    password varchar(31) not null,
    username varchar(255) not null,
    avatar varchar(255) not null unique,
    banned boolean not null default false,
    ban_reason varchar(255),
    status varchar(255),
    created_at timestamp
);

create table roles (
    id serial primary key,
    "name" varchar(255) not null unique
);

create table products (
    id serial primary key,
    title varchar(255) not null unique,
    body varchar(2047) not null,
    release_date date not null,
    created_at timestamp
);

create table genres (
    id serial primary key,
    name varchar(255) not null unique
);

create table platforms (
    id serial primary key,
    name varchar(255) not null unique
);

create table reviews (
    id serial primary key,
    body text not null unique,
    author_id int not null,
    product_id int not null,
    rate int not null default 50,
    created_at timestamp,
    updated_at timestamp
);

create table review_rating (
    id serial primary key,
    rate boolean not null, -- | 0 - dislike, 1 - like |
    review_id int not null,
    user_id int not null,
    foreign key ("user_id")
        references "users" ("id") on delete cascade,
    foreign key ("review_id")
        references "reviews" ("id") on delete cascade
);

create table review_comments (
    id serial primary key,
    body varchar(2047) not null,
    author_id int not null,
    review_id int not null,
    comment_id int,
    created_at timestamp,
    updated_at timestamp,
    foreign key ("author_id")
        references "users" ("id") on delete cascade,
    foreign key ("review_id")
        references "reviews" ("id") on delete cascade
);

create table users_follows (
    id serial primary key,
    follower_id int not null,
    publisher_id int not null,
    foreign key ("follower_id")
        references "users" ("id") on delete cascade,
    foreign key ("publisher_id")
        references "users" ("id") on delete cascade
);

create table users_roles (
    id serial primary key,
    user_id int not null,
    role_id int not null default 0,
    foreign key ("user_id")
        references "users" ("id") on delete cascade,
    foreign key ("role_id")
        references "roles" ("id") on delete set default
);

create table products_platforms (
    id serial primary key,
    product_id int not null,
    platform_id int not null,
    foreign key ("product_id")
        references "products" ("id") on delete cascade,
    foreign key ("platform_id")
        references "platforms" ("id") on delete cascade
);

create table products_genres (
    id serial primary key,
    product_id int not null,
    genre_id int not null default 0,
    foreign key ("product_id")
        references "products" ("id") on delete cascade,
    foreign key ("genre_id")
        references "genres" ("id") on delete cascade
);