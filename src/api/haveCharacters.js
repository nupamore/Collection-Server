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
router.get(
    '/members/:memberId/haveCharacters/:characterId',
    async (req, res) => {
        const { memberId, characterId } = req.params
        try {
            const result = await db.models.haveCharacters.findOne({
                where: { memberId, characterId },
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
    },
)

// 특정 유저에게 새로운 캐릭터들을 추가하는 api
router.post('/members/:memberId/haveCharacters', async (req, res) => {
    const { memberId } = req.params
    const { characters } = req.body
    characters.forEach(character => {
        character.memberId = memberId
    })

    try {
        await db.models.haveCharacters.bulkCreate(characters)
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

// 특정 유저가 가진 특정 캐릭터의 정보를 수정하는 api
router.put(
    '/members/:memberId/haveCharacters/:characterId',
    async (req, res) => {
        const { memberId, characterId } = req.params
        try {
            const result = await db.models.haveCharacters.findOne({
                where: { memberId, characterId },
            })
            if (!result) {
                res.status(404).json({
                    code: 404,
                    message: '해당 아이디를 가진 데이터를 찾을 수 않음',
                })
            } else {
                await db.models.haveCharacters.update(req.body, {
                    where: { memberId, characterId },
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
    },
)

// 특정 유저가 가진 캐릭터들의 정보를 수정하는 api
router.put('/members/:memberId/haveCharacters', async (req, res) => {
    const { memberId } = req.params
    const { characters } = req.body

    const promises = characters.map(character => {
        return db.models.haveCharacters.update(character, {
            where: {
                memberId,
                characterId: character.characterId,
            },
        })
    })
    try {
        await Promise.all(promises)
        res.status(200).json({
            code: 200,
            message: '수정 성공',
        })
    } catch (e) {
        res.status(400).json({
            code: 400,
            message: e.toString(),
        })
    }
})

// 특정 유저가 가진 특정 캐릭터의 정보를 삭제하는 api
router.delete(
    '/members/:memberId/haveCharacters/:characterId',
    async (req, res) => {
        const { memberId, characterId } = req.params
        try {
            await db.models.haveCharacters.destroy({
                where: { memberId, characterId },
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
    },
)

// 특정 유저가 가진 캐릭터들의 정보를 삭제하는 api
router.delete('/members/:memberId/haveCharacters', async (req, res) => {
    const { memberId } = req.params
    const { characterIds } = req.body

    try {
        await db.models.haveCharacters.destroy({
            where: {
                memberId,
                characterId: characterIds,
            },
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
