create table if not exists user_account(
  id bigserial primary key not null,
  username text not null,
  email text not null,
  password text not null,
  is_email_confirmed boolean not null,
  created_at timestamp not null,
  updated_at timestamp,
  unique(username),
  unique(email)
);
