table! {
    stat (id) {
        id -> Integer,
        succeed -> Unsigned<Integer>,
        failed -> Integer,
        created_at -> Timestamp,
        updated_at -> Timestamp,
    }
}

table! {
    token (id) {
        id -> Unsigned<Integer>,
        plain_text -> Varchar,
        hash_data -> Varchar,
        trusty -> Unsigned<Tinyint>,
        trusty_at -> Timestamp,
        created_at -> Timestamp,
        updated_at -> Timestamp,
        is_deleted -> Unsigned<Tinyint>,
    }
}

allow_tables_to_appear_in_same_query!(
    stat,
    token,
);
