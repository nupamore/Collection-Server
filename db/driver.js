const mysql = require('mysql2/promise')
const CONFIG = require('../config')
const pool = mysql.createPool(CONFIG.db)

async function read(queryStr, params) {
    try {
        const [rows] = await pool.query(queryStr, params)
        if (rows.length)
            return {
                code: 200,
                data: rows[0],
            }
        else
            return {
                code: 404,
                message: '해당 아이디를 가진 데이터를 찾을 수 않음',
            }
    } catch (e) {
        return {
            code: e.errno,
            message: e.message,
        }
    }
}

async function insert(queryStr, params) {
    try {
        await pool.query(queryStr, params)
        return {
            code: 200,
            message: '생성 성공',
        }
    } catch (e) {
        return {
            code: e.errno,
            message: e.message,
        }
    }
}

async function update(queryStr, params) {
    try {
        const [{ changedRows }] = await pool.query(queryStr, params)
        if (changedRows) {
            return {
                code: 200,
                message: '수정 성공',
            }
        } else {
            return {
                code: 404,
                message: '해당 아이디를 가진 데이터를 찾을 수 않음',
            }
        }
    } catch (e) {
        return {
            code: e.errno,
            message: e.message,
        }
    }
}

async function remove(queryStr, params) {
    try {
        const [{ affectedRows }] = await pool.query(queryStr, params)
        if (affectedRows) {
            return {
                code: 200,
                message: '삭제 성공',
            }
        } else {
            return {
                code: 404,
                message: '해당 아이디를 가진 데이터를 찾을 수 않음',
            }
        }
    } catch (e) {
        return {
            code: e.errno,
            message: e.message,
        }
    }
}

module.exports = {
    pool,
    read,
    insert,
    update,
    remove,
}
