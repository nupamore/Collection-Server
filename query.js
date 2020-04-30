module.exports = {
    GET_STRUCT: `
        SELECT id, type, content
        FROM struct
        WHERE id = ?
    `,
    INSERT_STRUCT: `
        INSERT INTO struct (id, type, content)
        VALUES (?, ?, ?)
    `,
    UPDATE_STRUCT: `
        UPDATE struct
        SET type = ?, content = ?
        WHERE id = ?
    `,
    DELETE_STRUCT: `
        DELETE FROM struct
        WHERE id = ?
    `,
}