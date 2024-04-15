export function iterateDays(startDateStr, endDateStr) {
  // 开始日期
  const startDate = new Date(startDateStr)

  // 结束日期
  const endDate = new Date(endDateStr)

  const rs = []
  // 按天迭代日期
  // eslint-disable-next-line no-unmodified-loop-condition
  for (let d = startDate; d <= endDate; d.setDate(d.getDate() + 1)) {
    rs.push(d.toISOString().split('T')[0])
  }
  return rs
}

/**
 *
 * @param {Array} data
 * @param {"day"|"week"|"month"} groupBy
 * @returns Array
 */
export function groupData(data, groupBy, date_key = 'date', count_key = 'count') {
  // 创建一个对象，用于保存按照日期分组后的结果
  const groupedData = {}

  // 遍历数据数组
  data.forEach((item) => {
    // 根据分组方式获取日期
    let date
    switch (groupBy) {
      case 'day':
        date = item[date_key]
        break
      case 'week':
        date = getYearWeek(item[date_key]) // 修改为获取全年第几周的函数
        break
      case 'month':
        date = item[date_key].slice(0, 7) // 获取年份和月份部分
        break
      default:
        console.error('Invalid groupBy value')
        return
    }

    // 如果日期不存在于 groupedData 中，则创建一个新的分组
    if (!groupedData[date]) {
      groupedData[date] = { [count_key]: 0 }
    }

    // 将当前数据的 count 累加到对应的分组中
    groupedData[date][count_key] += item[count_key]
  })

  // 将结果转换为数组形式并返回
  return Object.keys(groupedData).map((date) => {
    return { [date_key]: date, [count_key]: groupedData[date][count_key] }
  })
}

// 获取日期所在全年的第几周
function getYearWeek(dateStr) {
  const date = new Date(dateStr)
  const year = date.getFullYear()
  const week = getWeekNumberEx(dateStr)
  return year + '-W' + week
}

function getWeekNumberEx(dateString) {
  // 将日期字符串转换为 Date 对象
  var date = new Date(dateString);
  // 获取日期所在年份的第一天
  var yearStart = new Date(date.getFullYear(), 0, 1);
  // 计算日期与年初的毫秒差值
  var diffMilliseconds = date - yearStart;
  // 计算当前周数（从 0 开始）
  var weekNumber = Math.floor(diffMilliseconds / (7 * 24 * 60 * 60 * 1000));
  // 如果当前周数小于 0，则将其设置为 0
  if (weekNumber < 0) {
      weekNumber = 0;
  }
  // 加上初始偏移量，确保周数从 1 开始
  weekNumber += 1;
  return weekNumber;
}


// 获取日期所在周数
function getWeekNumber(date) {
  const d = new Date(
    Date.UTC(date.getFullYear(), date.getMonth(), date.getDate())
  )
  const dayNum = d.getUTCDay() || 7
  d.setUTCDate(d.getUTCDate() + 4 - dayNum)
  const yearStart = new Date(Date.UTC(d.getUTCFullYear(), 0, 1))
  return Math.ceil(((d - yearStart) / 86400000 + 1) / 7)
}

/**
 *
 * @param {Number} num
 * @returns String
 */
export function formatNumber(num) {
  return num.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
}