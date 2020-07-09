/* jshint indent: 2 */

module.exports = function(sequelize, DataTypes) {
  return sequelize.define('haveItems', {
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
    itemId: {
      type: DataTypes.INTEGER(10).UNSIGNED,
      allowNull: false
    },
    itemKey: {
      type: DataTypes.INTEGER(10).UNSIGNED,
      allowNull: false
    },
    stackNum: {
      type: DataTypes.INTEGER(11).UNSIGNED,
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
    tableName: 'haveItems'
  });
};
