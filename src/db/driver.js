const fs = require('fs')
const CONFIG = require('../config')
const { Sequelize, DataTypes } = require('sequelize')

const sequelize = new Sequelize({
    ...CONFIG.db,
    logging: false,
})

// define models
const files = fs.readdirSync('src/models')
const models = {}
files.map(file => {
    const define = require(`../models/${file}`)
    models[file] = define(sequelize, DataTypes)
})

module.exports = sequelize
exports.models = models
