export function useDateTimeForamat(date, options = {}) {
  let localeLg = 'zh-CN'

  let baseFormat = 'Y/m/d '

  switch (options.format) {
    case 'date':
      baseFormat = 'Y/m/d '
      break
    case 'time':
      baseFormat = 'H:i:s '
      break
    case 'datetime':
      baseFormat = 'Y/m/d H:i:s '
      break
    default:
      baseFormat = 'Y/m/d '
  }

  switch (options.local) {
    case 'en':
      localeLg = 'en-US'
      break
    default:
      localeLg = 'zh-CN'
  }

  const subdate = new Date(date)
  if (options.sub) {
    // Subtract days
    subdate.setDate(subdate.getDate() - options.sub)
  }

  if (!date) {
    return new Date().toLocaleDateString(localeLg).replace(/\//g, '-')
  }
  return subdate.toLocaleDateString(localeLg).replace(/\//g, '-')
}
