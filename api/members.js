const router = require('express').Router()
const QUERY = require('../db/query')
const db = require('../db/driver')

router.get('/members/:id', async (req, res) => {
    const result = await db.read(QUERY.GET_MEMBER, [req.params.id])
    if (result.code === 200) res.status(200).json(result)
    else res.status(400).json(result)
})

router.post('/members/:id', async (req, res) => {
    const result = await db.insert(QUERY.INSERT_MEMBER, [
        req.params.id,
        req.body.nickname,
        req.body.level,
        req.body.exp,
        req.body.money,
        req.body.berry,
        req.body.stardust,
        req.body.maxClearStage,
    ])
    if (result.code === 200) res.status(200).json(result)
    else res.status(400).json(result)
})

router.put('/members/:id', async (req, res) => {
    const result = await db.update(QUERY.UPDATE_MEMBER, [
        req.body.nickname,
        req.body.level,
        req.body.exp,
        req.body.money,
        req.body.berry,
        req.body.stardust,
        req.body.maxClearStage,
        req.params.id,
    ])
    if (result.code === 200) res.status(200).json(result)
    else res.status(400).json(result)
})

router.delete('/members/:id', async (req, res) => {
    const result = await db.remove(QUERY.DELETE_MEMBER, [req.params.id])
    if (result.code === 200) res.status(200).json(result)
    else res.status(400).json(result)
})

module.exports = router
