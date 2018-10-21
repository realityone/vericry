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
