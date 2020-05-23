require('dotenv').config()

module.exports = {
    db: {
        database: 'collection',
        username: process.env.DB_USER,
        password: process.env.DB_PASS,
        host: process.env.DB_HOST,
        dialect: 'mysql',
    },
    auth: {
        key: process.env.AUTH_KEY,
    },
}
