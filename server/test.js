function groupData(data, groupBy) {
    // 创建一个对象，用于保存按照日期分组后的结果
    const groupedData = {};

    // 遍历数据数组
    data.forEach(item => {
        // 根据分组方式获取日期
        let date;
        switch (groupBy) {
            case 'day':
                date = item.date;
                break;
            case 'week':
                date = getYearWeek(item.date); // 修改为获取全年第几周的函数
                break;
            case 'month':
                date = item.date.slice(0, 7); // 获取年份和月份部分
                break;
            default:
                console.error('Invalid groupBy value');
                return;
        }

        // 如果日期不存在于 groupedData 中，则创建一个新的分组
        if (!groupedData[date]) {
            groupedData[date] = {count: 0};
        }

        // 将当前数据的 count 累加到对应的分组中
        groupedData[date].count += item.count;
    });

    // 将结果转换为数组形式并返回
    return Object.keys(groupedData).map(date => {
        return {date: date, count: groupedData[date].count};
    });
}

// 获取日期所在全年的第几周
function getYearWeek(dateStr) {
    const date = new Date(dateStr);
    const year = date.getFullYear();
    const month = date.getMonth()
    const days = date.getDay()
    const week = getWeekNumber(date);
    return year + '-W' + week;
}

// 获取日期所在周数
function getWeekNumber(date) {
    const d = new Date(Date.UTC(date.getFullYear(), date.getMonth(), date.getDate()));
    const dayNum = d.getUTCDay() || 7;
    d.setUTCDate(d.getUTCDate() + 4 - dayNum);
    const yearStart = new Date(Date.UTC(d.getUTCFullYear(), 0, 1));
    return Math.ceil((((d - yearStart) / 86400000) + 1) / 7);
}

// 测试数据
const data = [
    {count: 20, date: "2024-04-02"},
    {count: 20, date: "2024-04-03"},
    {count: 20, date: "2024-04-04"},
    {count: 20, date: "2024-04-05"},
    {count: 20, date: "2024-04-06"},
    {count: 20, date: "2024-04-07"},
    {count: 20, date: "2024-04-08"},
    {count: 20, date: "2024-04-09"},

    // 其他数据...
];

// 测试按日分组
console.log(groupData(data, 'day'));

// 测试按周分组
console.log(groupData(data, 'week'));

// 测试按月分组
console.log(groupData(data, 'month'));
