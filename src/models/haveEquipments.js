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
      type: DataTypes.INTEGER(11).UNSIGNED,
      allowNull: false
    },
    equipmentKey: {
      type: DataTypes.INTEGER(11).UNSIGNED,
      allowNull: false
    },
    equipmentLevel: {
      type: DataTypes.INTEGER(11).UNSIGNED,
      allowNull: false,
      defaultValue: '0'
    },
    ownerCharacterId: {
      type: DataTypes.INTEGER(11),
      allowNull: false,
      defaultValue: '-1'
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
