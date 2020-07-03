/* jshint indent: 2 */

module.exports = function(sequelize, DataTypes) {
  return sequelize.define('haveEquipments', {
    id: {
      type: DataTypes.INTEGER(11).UNSIGNED,
      allowNull: false,
      primaryKey: true,
      autoIncrement: true
    },
    memberId: {
      type: DataTypes.STRING(50),
      allowNull: false
    },
    equipmentId: {
      type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
      allowNull: false
    },
    equipmentKey: {
      type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
      allowNull: false
    },
    equipmentLevel: {
      type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
      allowNull: false
    },
    ownerCharacterId: {
      type: DataTypes.INTEGER(11),
      allowNull: false,
      defaultValue: '-1'
    },
    attack: {
      type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
      allowNull: false
    },
    defense: {
      type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
      allowNull: false
    },
    spAttack: {
      type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
      allowNull: false
    },
    spDefense: {
      type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
      allowNull: false
    },
    speed: {
      type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
      allowNull: false
    },
    attackRange: {
      type: DataTypes.INTEGER(11).UNSIGNED.ZEROFILL,
      allowNull: false
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
    tableName: 'haveEquipments'
  });
};
