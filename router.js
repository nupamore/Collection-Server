const router = require('express').Router()
const mysql = require('mysql2/promise')
const CONFIG = require('./config')
const QUERY = require('./query')

const pool = mysql.createPool(CONFIG.db)

router.get('/struct/:key', async (req, res) => {
    const rows = await pool.query(QUERY.GET_STRUCT, req.params.key)
    if (rows.length) res.json(rows[0])
    else res.status(404).json({
        message: "해당 key를 찾을 수 없음"
    })
})

module.exports = router