const router = require('express').Router()
const mysql = require('mysql2/promise')
const CONFIG = require('./config')
const QUERY = require('./query')

const pool = mysql.createPool(CONFIG.db)

router.get('/struct/:id', async (req, res) => {
    try {
        const [rows] = await pool.query(QUERY.GET_STRUCT, [req.params.id])
        if (rows.length) res.json({
            code: 200,
            data: rows[0],
        })
        else res.status(400).json({
            code: 404,
            message: "해당 아이디를 가진 데이터를 찾을 수 않음"
        })
    } catch (e) {
        res.status(400).json({
            code: e.errno,
            message: e.message,
        })
    }
})

router.post('/struct/:id', async (req, res) => {
    try {
        const [rows] = await pool.query(QUERY.INSERT_STRUCT, [
            req.params.id,
            req.body.type,
            req.body.content,
        ])
        res.json({
            code: 200,
            message: "생성 성공",
        })
    } catch (e) {
        res.status(400).json({
            code: e.errno,
            message: e.message,
        })
    }
})

router.put('/struct/:id', async (req, res) => {
    try {
        const [{changedRows}] = await pool.query(QUERY.UPDATE_STRUCT, [
            req.body.type,
            req.body.content,
            req.params.id,
        ])
        if (changedRows) {
            res.json({
                code: 200,
                message: "수정 성공",
            })
        } else {
            res.status(400).json({
                code: 404,
                message: "해당 아이디를 가진 데이터를 찾을 수 않음"
            })
        }
    } catch (e) {
        res.status(400).json({
            code: e.errno,
            message: e.message,
        })
    }
})

router.delete('/struct/:id', async (req, res) => {
    try {
        const [{changedRows}] = await pool.query(QUERY.DELETE_STRUCT, [req.params.id,])
        if (changedRows) {
            res.json({
                code: 200,
                message: "수정 성공",
            })
        } else {
            res.status(400).json({
                code: 404,
                message: "해당 아이디를 가진 데이터를 찾을 수 않음"
            })
        }
    } catch (e) {
        res.status(400).json({
            code: e.errno,
            message: e.message,
        })
    }
})

module.exports = router