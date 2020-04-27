module.exports = {
    GET_STRUCT = `
        SELECT key, type, content
        FROM struct
        WHERE key = ?
    `,
    POST_STRUCT = `
        SELECT key, type, content
        FROM struct
        WHERE key = ?
    `,
    PUT_STRUCT = `
        SELECT key, type, content
        FROM struct
        WHERE key = ?
    `,
    DELETE_STRUCT = `
        SELECT key, type, content
        FROM struct
        WHERE key = ?
    `,
}