const fs = require('fs')
const CONFIG = require('../config')
const { Sequelize, DataTypes } = require('sequelize')

const sequelize = new Sequelize(CONFIG.db)

// define models
const files = fs.readdirSync('models')
const models = {}
files.map(file => {
    const define = require(`../models/${file}`)
    models[file] = define(sequelize, DataTypes)
})
// connection test
;(async () => {
    try {
        await sequelize.authenticate()
        console.log('Connection has been established successfully.')
    } catch (error) {
        console.error('Unable to connect to the database:', error)
    }
})()

module.exports = sequelize
exports.models = models
