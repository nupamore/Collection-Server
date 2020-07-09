/* jshint indent: 2 */

module.exports = function(sequelize, DataTypes) {
  return sequelize.define('members', {
    id: {
      type: DataTypes.STRING(50),
      allowNull: false,
      primaryKey: true
    },
    nickname: {
      type: DataTypes.STRING(50),
      allowNull: true
    },
    level: {
      type: DataTypes.INTEGER(10).UNSIGNED,
      allowNull: false
    },
    exp: {
      type: DataTypes.INTEGER(11).UNSIGNED,
      allowNull: false
    },
    money: {
      type: DataTypes.INTEGER(10).UNSIGNED,
      allowNull: false
    },
    berry: {
      type: DataTypes.INTEGER(11).UNSIGNED,
      allowNull: false
    },
    stardust: {
      type: DataTypes.INTEGER(11).UNSIGNED,
      allowNull: false
    },
    maxClearBattleStage: {
      type: DataTypes.INTEGER(11).UNSIGNED,
      allowNull: false
    },
    lastSelectedChapter: {
      type: DataTypes.INTEGER(10).UNSIGNED,
      allowNull: false
    },
    maxOpenedChapter: {
      type: DataTypes.INTEGER(10).UNSIGNED,
      allowNull: false
    },
    lastSelectedParty: {
      type: DataTypes.INTEGER(10).UNSIGNED,
      allowNull: false
    },
    battleClearCount: {
      type: DataTypes.INTEGER(10).UNSIGNED,
      allowNull: false
    },
    getCharacterCount: {
      type: DataTypes.INTEGER(10).UNSIGNED,
      allowNull: false
    },
    maxHaveCharactersNum: {
      type: DataTypes.INTEGER(10).UNSIGNED,
      allowNull: false,
      defaultValue: '100'
    },
    maxHaveEquipmentsNum: {
      type: DataTypes.INTEGER(10).UNSIGNED,
      allowNull: false,
      defaultValue: '100'
    },
    index: {
      type: DataTypes.STRING(1000),
      allowNull: true
    },
    battleSpeed: {
      type: DataTypes.INTEGER(10).UNSIGNED,
      allowNull: false,
      defaultValue: '1'
    },
    createdAt: {
      type: DataTypes.DATE,
      allowNull: true
    },
    updatedAt: {
      type: DataTypes.DATE,
      allowNull: true
    }
  }, {
    tableName: 'members'
  });
};
