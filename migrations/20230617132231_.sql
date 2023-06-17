-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS hotels
(
    id          BIGSERIAL PRIMARY KEY NOT NULL unique,
    name        text                  NOT NULL DEFAULT '',
    description text                  NOT NULL default '',
    facilities  text                  NOT NULL default '',
    address     text                  NOT NULL DEFAULT '',
    price       int                   NOT NULL,
    rating      float                 NOT NULL
);

CREATE TABLE IF NOT EXISTS bookings
(
    id             BIGSERIAL PRIMARY KEY                        NOT NULL unique,
    hotel_id       int references hotels (id) on delete cascade not null,
    arrival_date   timestamp                                    NOT NULL default now(),
    departure_date timestamp                                    NOT NULL default now()
);

CREATE TABLE IF NOT EXISTS reviews
(
    id       BIGSERIAL PRIMARY KEY                        NOT NULL unique,
    hotel_id int references hotels (id) on delete cascade not null,
    score    int                                          not null,
    feedback text                                         not null default ''
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table reviews, bookings, hotels;
-- +goose StatementEnd
