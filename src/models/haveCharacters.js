/* jshint indent: 2 */

module.exports = function(sequelize, DataTypes) {
  return sequelize.define('haveCharacters', {
    id: {
      type: DataTypes.INTEGER(11),
      allowNull: false,
      primaryKey: true,
      autoIncrement: true
    },
    memberId: {
      type: DataTypes.STRING(50),
      allowNull: false
    },
    characterID: {
      type: DataTypes.INTEGER(10).UNSIGNED.ZEROFILL,
      allowNull: false
    },
    characterKey: {
      type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
      allowNull: false
    },
    characterLevel: {
      type: DataTypes.INTEGER(11).UNSIGNED,
      allowNull: false,
      defaultValue: '1'
    },
    characterExp: {
      type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
      allowNull: false
    },
    joinedParty: {
      type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
      allowNull: false
    },
    joinedSlot: {
      type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
      allowNull: false
    },
    EvHealth: {
      type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
      allowNull: false
    },
    EvAttack: {
      type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
      allowNull: false
    },
    EvDefense: {
      type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
      allowNull: false
    },
    EvSpAttack: {
      type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
      allowNull: false
    },
    EvSpDefense: {
      type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
      allowNull: false
    },
    EvSpeed: {
      type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
      allowNull: false
    },
    skill1Level: {
      type: DataTypes.INTEGER(11).UNSIGNED,
      allowNull: true,
      defaultValue: '1'
    },
    skill2Level: {
      type: DataTypes.INTEGER(10).UNSIGNED,
      allowNull: true,
      defaultValue: '1'
    },
    skill3Level: {
      type: DataTypes.INTEGER(10).UNSIGNED,
      allowNull: true,
      defaultValue: '1'
    },
    skill4Level: {
      type: DataTypes.INTEGER(10).UNSIGNED,
      allowNull: true,
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
    tableName: 'haveCharacters'
  });
};
