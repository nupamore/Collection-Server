const request = require('supertest')
const server = require('../server')
const CONFIG = require('../config')
require('mysql2/node_modules/iconv-lite').encodingExists('foo')

const defaultHeaders = {
    Authorization: CONFIG.auth.key,
}

describe('members', () => {
    it('get all members', async done => {
        await request(server)
            .get('/api/members')
            .set(defaultHeaders)
        done()
    })
})

describe('haveCharacters', () => {
    it('get all characters', async done => {
        await request(server)
            .get('/api/haveCharacters')
            .set(defaultHeaders)
        done()
    })
})

describe('haveEquipments', () => {
    it('get all equipments', async done => {
        await request(server)
            .get('/api/haveEquipments')
            .set(defaultHeaders)
        done()
    })
})

afterAll(async done => {
    await server.close()
    done()
})
