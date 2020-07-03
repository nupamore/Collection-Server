const router = require('express').Router()
const db = require('../db/driver')

router.get('/haveItems', async (req, res) => {
    try {
        const result = await db.models.haveItems.findAll()
        res.status(200).json({
            code: 200,
            data: result,
        })
    } catch (e) {
        res.status(400).json({
            code: 400,
            message: e.toString(),
        })
    }
})

module.exports = router
