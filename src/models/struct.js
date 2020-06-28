/* jshint indent: 2 */

module.exports = function(sequelize, DataTypes) {
    return sequelize.define(
        'struct',
        {
            id: {
                type: DataTypes.STRING(100),
                allowNull: false,
                primaryKey: true,
            },
            type: {
                type: DataTypes.STRING(100),
                allowNull: true,
            },
            content: {
                type: DataTypes.STRING(10000),
                allowNull: true,
            },
        },
        {
            tableName: 'struct',
        },
    )
}
