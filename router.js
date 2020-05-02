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

module.exports = router
