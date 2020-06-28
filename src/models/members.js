/* jshint indent: 2 */

module.exports = function(sequelize, DataTypes) {
    return sequelize.define(
        'members',
        {
            id: {
                type: DataTypes.STRING(50),
                allowNull: false,
                primaryKey: true,
            },
            nickname: {
                type: DataTypes.STRING(50),
                allowNull: true,
            },
            level: {
                type: DataTypes.INTEGER(10).UNSIGNED.ZEROFILL,
                allowNull: false,
            },
            exp: {
                type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
                allowNull: false,
            },
            money: {
                type: DataTypes.INTEGER(10).UNSIGNED.ZEROFILL,
                allowNull: false,
            },
            berry: {
                type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
                allowNull: false,
            },
            stardust: {
                type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
                allowNull: false,
            },
            maxClearBattleStage: {
                type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
                allowNull: false,
            },
            lastSelectedChapter: {
                type: DataTypes.INTEGER(10).UNSIGNED.ZEROFILL,
                allowNull: false,
            },
            maxOpenedChapter: {
                type: DataTypes.INTEGER(10).UNSIGNED.ZEROFILL,
                allowNull: false,
            },
            lastSelectedParty: {
                type: DataTypes.INTEGER(10).UNSIGNED.ZEROFILL,
                allowNull: false,
            },
            battleClearCount: {
                type: DataTypes.INTEGER(10).UNSIGNED.ZEROFILL,
                allowNull: false,
            },
            getCharacterCount: {
                type: DataTypes.INTEGER(10).UNSIGNED.ZEROFILL,
                allowNull: false,
            },
            maxHaveCharactersNum: {
                type: DataTypes.INTEGER(10).UNSIGNED.ZEROFILL,
                allowNull: false,
            },
            maxHaveEquipmentsNum: {
                type: DataTypes.INTEGER(10).UNSIGNED.ZEROFILL,
                allowNull: false,
            },
            index: {
                type: DataTypes.STRING(50),
                allowNull: true,
            },
            battleSpeed: {
                type: DataTypes.INTEGER(10).UNSIGNED,
                allowNull: false,
                defaultValue: '1',
            },
            createdAt: {
                type: DataTypes.DATE,
                allowNull: true,
            },
            updatedAt: {
                type: DataTypes.DATE,
                allowNull: true,
            },
        },
        {
            tableName: 'members',
        },
    )
}
