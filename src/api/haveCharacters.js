const router = require('express').Router()
const db = require('../db/driver')

// 특정 유저가 가진 캐릭터들을 가져오는 api
router.get('/members/:memberId/haveCharacters', async (req, res) => {
    try {
        const result = await db.models.haveCharacters.findAll({
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

// 특정 유저가 가진 특정 캐릭터를 가져오는 api
router.get('/members/:memberId/haveCharacters/:characterId', async (req, res) => {
    try {
        const result = await db.models.haveCharacters.findOne({
            where: { memberId: req.params.memberId, characterId: req.params.characterId },
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

// 특정 유저에게 새로운 캐릭터들을 추가한다.
router.post('/members/:memberId/haveCharacters', async (req, res) => {
    const form = req.body
    form.id = req.params.id

    try {
        await db.models.haveCharacters.create(form)
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

router.put('/haveCharacters/:id', async (req, res) => {
    try {
        const result = await db.models.haveCharacters.findOne({
            where: { id: req.params.id },
        })
        if (!result) {
            res.status(404).json({
                code: 404,
                message: '해당 아이디를 가진 데이터를 찾을 수 않음',
            })
        } else {
            await db.models.haveCharacters.update(req.body, {
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

router.delete('/haveCharacters/:delete', async (req, res) => {
    try {
        await db.models.haveCharacters.destroy({
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
