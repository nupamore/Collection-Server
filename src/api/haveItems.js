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
    
    const existItems = []
    const notExistItems = []
 
    try {
        // 존재하는 아이템 체크
        const existPromises = items.map(item => {
            return db.models.haveItems.findOne({
                where: { 
                    memberId,
                    itemKey: item.itemKey
                },
            })
        })
        (await Promise.all(existPromises)).forEach((result, index) => {
            if (result) existItems.push(items[index])
            else notExistItems.push(items[index])
        })
    
        // 레코드 삽입
        const createPromise = db.models.haveItems.bulkCreate(notExistItems)
        
        // 레코드 업데이트
        const updatePromises = existItems.map(existItem => {
            return db.models.haveItems.update({
                where: {
                    memberId,
                    itemKey: existItem.itemKey
                },
            })   
        })
        
        await Promise.all([
            createPromise,
            ...updatePromises
        ])
        res.status(200).json({
            code: 200,
            data: '요청 성공',
        })
    } catch (e) {
        res.status(400).json({
            code: 400,
            message: e.toString(),
        })
    }
})

module.exports = router
