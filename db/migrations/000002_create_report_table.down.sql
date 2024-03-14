create if not exists report_entries(
  id bigserial primary key not_null,
  user_id not null,
  title varchar(25) not null,
  amount bigint not null, 
  created_at timestamp not null,
  updated_at timestamp not null,
  deleted_at timestamp,
  constraint fk_user_id foreign key (user_id) references users(id)
)
