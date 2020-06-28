const router = require('express').Router()
const db = require('../db/driver')

router.get('/members', async (req, res) => {
    try {
        const result = await db.models.members.findAll()
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

router.get('/members/:id', async (req, res) => {
    try {
        const result = await db.models.members.findOne({
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

router.post('/members/:id', async (req, res) => {
    const form = req.body
    form.id = req.params.id

    try {
        await db.models.members.create(form)
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

router.put('/members/:id', async (req, res) => {
    try {
        const result = await db.models.members.findOne({
            where: { id: req.params.id },
        })
        if (!result) {
            res.status(404).json({
                code: 404,
                message: '해당 아이디를 가진 데이터를 찾을 수 않음',
            })
        } else {
            await db.models.members.update(req.body, {
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

router.delete('/members/:delete', async (req, res) => {
    try {
        await db.models.members.destroy({
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
