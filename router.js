const router = require('express').Router()
const CONFIG = require('./config')

router.use((req, res, next) => {
    const key = req.get('Authorization')
    if (key === CONFIG.auth.key) next()
    else
        res.status(403).json({
            code: 403,
            message: '인증 필요',
        })
})
router.use(require('./api/members'))
router.use(require('./api/haveCharacters'))
router.use(require('./api/haveEquipments'))

module.exports = router
