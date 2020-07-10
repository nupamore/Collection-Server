const router = require('express').Router()
const db = require('../db/driver')

// 특정 유저가 가진 장비들을 가져오는 api
router.get('/members/:memberId/haveEquipments', async (req, res) => {
    try {
        const result = await db.models.haveEquipments.findAll({
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

router.get('/haveEquipments/:id', async (req, res) => {
    try {
        const result = await db.models.haveEquipments.findOne({
            where: { id: req.params.id },
        })
        if (result)
            res.status(200).json({
                code: 200,
                data: result.dataValues,
            })
        else
            res.status(404).json({
                code: 404,
                message: '해당 아이디를 가진 데이터를 찾을 수 않음',
            })
    } catch (e) {
        res.status(400).json({
            code: 400,
            message: e.toString(),
        })
    }
})

// 특정 유저에게 새로운 장비들을 추가하는 api
router.post('/members/:memberId/haveEquipments', async (req, res) => {
    const { memberId } = req.params
    const { equipments } = req.body
    equipments.forEach(equipment => {
        equipment.memberId = memberId  
    })

    try {
        await db.models.haveEquipments.bulkCreate(equipments)
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

router.put('/haveEquipments/:id', async (req, res) => {
    try {
        const result = await db.models.haveEquipments.findOne({
            where: { id: req.params.id },
        })
        if (!result) {
            res.status(404).json({
                code: 404,
                message: '해당 아이디를 가진 데이터를 찾을 수 않음',
            })
        } else {
            await db.models.haveEquipments.update(req.body, {
                where: { id: req.params.id },
            })
            res.status(200).json({
                code: 200,
                message: '수정 성공',
            })
        }
    } catch (e) {
        res.status(400).json({
            code: 400,
            message: e.toString(),
        })
    }
})

router.delete('/haveEquipments/:delete', async (req, res) => {
    try {
        await db.models.haveEquipments.destroy({
            where: { id: req.params.id },
        })
        res.status(200).json({
            code: 200,
            message: '삭제 성공',
        })
    } catch (e) {
        res.status(400).json({
            code: 400,
            message: e.toString(),
        })
    }
})

module.exports = router
