const request = require('supertest')
const server = require('../server')
const CONFIG = require('../config')
require('mysql2/node_modules/iconv-lite').encodingExists('foo')

const defaultHeaders = {
    Authorization: CONFIG.auth.key,
}

describe('members', () => {
    it('get all members', async () => {
        const { statusCode } = await request(server)
            .get('/api/members')
            .set(defaultHeaders)
        expect(statusCode).toBe(200)
    })
})

describe('haveCharacters', () => {
    it('get test characters', async () => {
        const { statusCode } = await request(server)
            .get('/api/members/test/haveCharacters')
            .set(defaultHeaders)
        expect(statusCode).toBe(200)
    })
})

describe('haveEquipments', () => {
    it('get test equipments', async () => {
        const { statusCode } = await request(server)
            .get('/api/members/test/haveEquipments')
            .set(defaultHeaders)
        expect(statusCode).toBe(200)
    })
})

describe('haveItems', () => {
    it('get test items', async () => {
        const { statusCode } = await request(server)
            .get('/api/members/test/haveItems')
            .set(defaultHeaders)
        expect(statusCode).toBe(200)
    })
})

afterAll(async done => {
    await server.close()
    done()
})
