-- +goose Up
create table user (
    id bigserial primary key,
    name varchar(100) not null,
    username varchar(50) not null unique,
    email varchar(100) not null unique,
    profile_picture varchar(50),
    created_at timestamptz not null default now()
);

create table follower (
    follower_id bigserial,
    following_id bigserial,
    primary key (follower_id, following_id),
    foreign key (following_id) references user(id) on delete cascade,
    foreign key (follower_id) references user(id) on delete cascade,
    followed_on timestamptz not null default now()
);

create table travel_entry (
    id bigserial primary key,
    owner_id bigserial,
    country varchar(50),
    street_address varchar(50),
    city varchar(50),
    state varchar(50),
    postal_code varchar(50),
    media text[],
    description text,
    visibility boolean,
    visit_date date not null default now(),
    foreign key (owner_id) references user(id) on delete cascade
);

create table comment (
    entry_id bigserial,
    user_id bigserial,
    content text,
    commented_on date not null default now(),
    foreign key (entry_id) references travel_entry(id) on delete cascade,
    foreign key (user_id) references user(id) on delete cascade,
    primary key (user_id, entry_id, commented_on)
);


-- +goose Down
