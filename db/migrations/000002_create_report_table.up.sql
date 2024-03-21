create table if not exists budget(
  id bigserial primary key not null,
  user_id bigint not null,
  amount bigint not null,
  time_code timestamp not null,
  created_at timestamp not null,
  updated_at timestamp not null,
  unique(user_id, time_code),
  constraint fk_user_id foreign key (user_id) references users(id)
);

create table if not exists report_entries(
  id bigserial primary key not null,
  user_id bigint not null,
  budget_id bigint not null,
  title varchar(25) not null,
  notes text,
  amount bigint not null,
  created_at timestamp not null,
  updated_at timestamp not null,
  deleted_at timestamp, 
  constraint fk_budget_id foreign key (budget_id) references budget(id),
  constraint fk_user_id foreign key (user_id) references users(id)
);
