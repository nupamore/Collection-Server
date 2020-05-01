module.exports = class Member {
    constructor(json) {
        this.id = json.id
        this.nickname = json.nickname
        this.level = json.level
        this.exp = json.exp
        this.money = json.money
        this.berry = json.berry
        this.stardust = json.stardust
        this.maxClearStage = json.maxClearStage
    }
}
