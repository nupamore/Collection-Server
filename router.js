const router = require('express').Router()
const mysql = require('mysql2/promise')
const CONFIG = require('./config')
const QUERY = require('./query')

const pool = mysql.createPool(CONFIG.db)

router.get('/struct/:id', async (req, res) => {
    const [rows] = await pool.query(QUERY.GET_STRUCT, [req.params.id])
    if (rows.length) res.json(rows[0])
    else res.status(404).json({
        message: "해당 아이디를 가진 데이터를 찾을 수 않음"
    })
})

router.post('/struct/:id', async (req, res) => {
    try {
        const [rows] = await pool.query(QUERY.INSERT_STRUCT, [
            req.params.id,
            req.body.type,
            req.body.content,
        ])
        res.json({
            message: "생성 성공"
        })
    } catch (e) {
        res.status(409).json({
            message: "해당 아이디를 가진 데이터가 이미 존재함"
        })
    }
})

router.put('/struct/:id', async (req, res) => {
    try {
        const [rows] = await pool.query(QUERY.UPDATE_STRUCT, [
            req.body.type,
            req.body.content,
            req.params.id,
        ])
        res.json({
            message: "수정 성공"
        })
    } catch (e) {
        res.status(404).json({
            message: "해당 아이디를 가진 데이터를 찾을 수 않음"
        })
    }
})

router.delete('/struct/:id', async (req, res) => {
    try {
        const [rows] = await pool.query(QUERY.DELETE_STRUCT, [req.params.id])
        res.json({
            message: "삭제 성공"
        })
    } catch (e) {
        res.status(404).json({
            message: "해당 아이디를 가진 데이터를 찾을 수 않음"
        })
    }
})

module.exports = router