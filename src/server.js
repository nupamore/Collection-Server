const express = require('express')
const router = require('./router')

const app = express()
app.use(express.json())
app.use(express.urlencoded({ extended: false }))
app.use('/api', router)

const port = 9001
module.exports = app.listen(port, () => {
    console.log(`Server start ${port}`)
})
