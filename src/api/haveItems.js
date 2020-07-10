const router = require('express').Router()
const db = require('../db/driver')

// 특정 유저가 가진 아이템들을 가져오는 api
router.get('/members/:memberId/haveItems', async (req, res) => {
    try {
        const result = await db.models.haveItems.findAll({
            where: { memberId: req.params.memberId },
        })
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

// 특정 유저에게 새로운 아이템들을 추가하는 api
router.post('/members/:memberId/haveItems', async (req, res) => {
    const { memberId } = req.params
    const { items } = req.body
    items.forEach(item => {
        item.memberId = memberId
    })

    try {
        await db.models.haveItems.bulkCreate(items)
        res.status(200).json({
            code: 200,
            data: '생성 성공',
        })
    } catch (e) {
        res.status(400).json({
            code: 400,
            message: e.toString(),
        })
    }
})

module.exports = router
