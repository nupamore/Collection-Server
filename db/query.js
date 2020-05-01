module.exports = {
    GET_MEMBER: `
        SELECT 
            id, nickname, level, exp, money, berry, stardust, maxClearStage, createdAt, updatedAt
        FROM members
        WHERE id = ?
    `,
    INSERT_MEMBER: `
        INSERT INTO members (
            id, nickname, level, exp, money, berry, stardust, maxClearStage, createdAt, updatedAt
        )
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())
    `,
    UPDATE_MEMBER: `
        UPDATE members
        SET 
            nickname = ?,
            level = ?,
            exp = ?,
            money = ?,
            berry = ?,
            stardust = ?,
            maxClearStage = ?,
            updatedAt = NOW()
        WHERE id = ?
    `,
    DELETE_MEMBER: `
        DELETE FROM members
        WHERE id = ?
    `,
}
